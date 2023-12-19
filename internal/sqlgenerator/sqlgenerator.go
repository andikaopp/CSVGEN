package sqlgenerator

import (
	"bufio"
	"fmt"
	"os"

	"github.com/google/uuid"
)

func GenerateSql() {
	// Buka file txt dan scan file yang ada di dalam folder input
	txtFile, err := os.Open("input/file.txt")
	if err != nil {
		fmt.Println("Error membuka file txt:", err)
		return
	}
	defer txtFile.Close()

	// Baca isi file txt yang ada di dalam folder input
	var kodeUnikSlice []string
	scanner := bufio.NewScanner(txtFile)
	for scanner.Scan() {
		kodeUnikSlice = append(kodeUnikSlice, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Println("Error membaca file txt:", err)
		return
	}

	//Buat file sql
	sqlFile, err := os.Create("output/output.sql")
	if err != nil {
		fmt.Println("Error membuat file SQL:", err)
		return
	}
	defer sqlFile.Close()

	//Inisialisi writer sql
	sqlWriter := bufio.NewWriter(sqlFile)
	defer sqlWriter.Flush()

	// Generate UUID dan tulis ke file CSV bersama dengan kode unik
	for i := 0; i < len(kodeUnikSlice) && i < 1000000; i++ {
		newUUID := uuid.New()
		sqlStatement := fmt.Sprintf("INSERT INTO tbl_cif (OID, CIF, COUNTER, STATUS_PDN_PEKERJAAN, CREATED_AT, UPDATED_AT) VALUES ('%s', '%s', 0, false, now(), now());\n",
			newUUID.String(), kodeUnikSlice[i])

		_, err := sqlWriter.WriteString(sqlStatement)
		if err != nil {
			fmt.Println("Error menulis ke file SQL:", err)
			return
		}
	}

	fmt.Println("File SQL berhasil dibuat dengan perintah INSERT yang sesuai.")
}
