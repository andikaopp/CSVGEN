package csvgenerator

import (
	"bufio"
	"encoding/csv"
	"fmt"
	"os"

	"github.com/google/uuid"
)

func GenerateCSV() {
	// Buka file txt yang berisi kode unik
	txtFile, err := os.Open("kode_unik.txt")
	if err != nil {
		fmt.Println("Error membuka file txt:", err)
		return
	}
	defer txtFile.Close()

	// Baca isi file txt
	var kodeUnikSlice []string
	scanner := bufio.NewScanner(txtFile)
	for scanner.Scan() {
		kodeUnikSlice = append(kodeUnikSlice, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error membaca file txt:", err)
		return
	}

	// Buat file CSV untuk menulis data
	csvFile, err := os.Create("output.csv")
	if err != nil {
		fmt.Println("Error membuat file CSV:", err)
		return
	}
	defer csvFile.Close()

	// Inisialisasi writer CSV
	writer := csv.NewWriter(csvFile)
	defer writer.Flush()

	// Tulis header ke file CSV
	header := []string{"field_1", "field_2", "field_3", "field_4"}
	err = writer.Write(header)
	if err != nil {
		fmt.Println("Error menulis header ke file CSV:", err)
		return
	}

	// Generate UUID dan tulis ke file CSV bersama dengan kode unik
	for i := 0; i < len(kodeUnikSlice) && i < 1000000; i++ {
		newUUID := uuid.New()
		row := []string{newUUID.String(), kodeUnikSlice[i], "0", "false"}
		err := writer.Write(row)
		if err != nil {
			fmt.Println("Error menulis baris ke file CSV:", err)
			return
		}
	}

	fmt.Println("File CSV berhasil dibuat dengan 1 juta baris data.")
}
