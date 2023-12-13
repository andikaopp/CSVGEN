package csvgenerator

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

const (
	codeLength = 10
)

var (
	charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

func generateUniqueCode() string {
	rand.Seed(time.Now().UnixNano())

	code := make([]byte, codeLength)
	for i := range code {
		code[i] = charset[rand.Intn(len(charset))]
	}
	return string(code)
}

func GenerateTxt() {
	file, err := os.Create("unique_codes.txt")
	if err != nil {
		fmt.Println("Gagal membuat file:", err)
		return
	}
	defer file.Close()

	for i := 0; i < 100000; i++ {
		code := generateUniqueCode()
		_, err := file.WriteString(code + "\n")
		if err != nil {
			fmt.Println("Gagal menulis ke file:", err)
			return
		}
	}

	fmt.Println("100000 kode unik telah dibuat dan disimpan dalam file unique_codes.txt")
}
