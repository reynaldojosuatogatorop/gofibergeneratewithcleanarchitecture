package main

import (
	"bufio"
	"fmt"
	"gofibergenerator/dependencies"
	"os"
	"strings"
)

// Fungsi rekursif untuk membuat folder dan file berdasarkan struktur folder

type Config struct {
	Database Database `json:"database" form:"database"`
	Redis    Redis    `json:"redis" form:"redis"`
}

type Server struct {
	Port string `json:"port" form:"port"`
}

type Database struct {
	Server   string `json:"server" form:"server"`
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

type Redis struct {
	Server   string `json:"server" form:"server"`
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Masukkan nama folder: ")
	folderName, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error membaca input:", err)
		return
	}

	folderName = strings.TrimSpace(folderName) // Menghapus karakter newline atau spasi

	if err := dependencies.GenerateFolderStructure(folderName); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Struktur folder berhasil dibuat dan modul Go diinisialisasi!")
	}

	// Membuka folder di VSCode
	if err := dependencies.OpenInVSCode(folderName); err != nil {
		fmt.Println("Error membuka folder di VSCode:", err)
		return
	}
}
