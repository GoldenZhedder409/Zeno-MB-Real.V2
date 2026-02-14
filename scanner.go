package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"sync"
	"time"
	"bufio"
)

// Tambahkan di scanner.go
func randomUserAgent() string {
	// Database user-agent berdasarkan region
	agents := []string{
		"Mozilla/5.0 (Windows NT 10.0; Win64; x64)", // US
		"Mozilla/5.0 (X11; Linux x86_64)",           // EU
		"Mozilla/5.0 (Android 10; Mobile)",          // Asia
	}
	return agents[time.Now().Unix()%int64(len(agents))]
}

// Modifikasi di scanner.go
func adaptiveScan(ip string) {
	// Sesuaikan kecepatan scan berdasarkan respon jaringan
	initialTimeout := 1 * time.Second
	if isNetworkSlow(ip) {
		initialTimeout = 3 * time.Second
	}
	// Lanjutkan scan dengan timeout yang adaptif
}

var COMMON_PORTS = []int{22, 23, 80, 443, 8080, 2323}

func isOpen(ip string, port int) bool {
	address := ip + ":" + strconv.Itoa(port)
	conn, err := net.DialTimeout("tcp", address, time.Second)
	if err != nil {
		return false
	}
	conn.Close()
	return true
}

func scanIP(ip string, wg *sync.WaitGroup, mu *sync.Mutex, targets *[]string) {
	defer wg.Done()
	openPorts := []int{}

	for _, port := range COMMON_PORTS {
		if isOpen(ip, port) {
			openPorts = append(openPorts, port)
		}
	}

	if len(openPorts) > 0 {
		for _, port := range openPorts {
			fmt.Printf("[OPEN] %s:%d\n", ip, port)
			mu.Lock()
			*targets = append(*targets, fmt.Sprintf("%s:%d", ip, port))
			mu.Unlock()
		}
	}
}

func RunScanner() {
	fmt.Println("[INFO] Memulai auto-scan IP dalam jaringan 192.168.1.1 - 254...")
	time.Sleep(1 * time.Second)

	var wg sync.WaitGroup
	var mu sync.Mutex
	targets := []string{}

	for i := 1; i <= 254; i++ {
		ip := "192.168.1." + strconv.Itoa(i)
		wg.Add(1)
		go scanIP(ip, &wg, &mu, &targets)
	}

	wg.Wait()

	if len(targets) > 0 {
		file, err := os.Create("targets.txt")
		if err != nil {
			fmt.Println("Gagal menulis file:", err)
			return
		}
		defer file.Close()

		writer := bufio.NewWriter(file)
		for _, target := range targets {
			fmt.Fprintln(writer, target)
		}
		writer.Flush()
		fmt.Println("\n[✓] Target rentan disimpan di targets.txt")
	} else {
		fmt.Println("\n[!] Tidak ditemukan target terbuka!")
	}

	fmt.Print("\n[ENTER] kembali ke menu...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')
}
