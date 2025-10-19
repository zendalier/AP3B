package main

import "fmt"

func main() {
	nilaiMahasiswa := make(map[string]int)
	nilaiMahasiswa["Budi"] = 85
	nilaiMahasiswa["Ani"] = 90
	nilaiMahasiswa["Citra"] = 88

	fmt.Println("Map nilai mahasiswa:", nilaiMahasiswa)

	kodePos := map[string]string{
		"Jakarta":  "10110",
		"Bandung":  "40111",
		"Surabaya": "60111",
		"Depok":    "16424",
	}

	fmt.Println("\nKode pos Depok:", kodePos["Depok"])

	kodePos["Yogyakarta"] = "55000"
	fmt.Println("Kode pos Yogyakarta:", kodePos["Yogyakarta"])

	nilai, ada := nilaiMahasiswa["Budi"]
	if ada {
		fmt.Printf("\nNilai Budi: %d\n", nilai)
	}

	fmt.Println("\nDaftar nilai mahasiswa:")
	for nama, nilai := range nilaiMahasiswa {
		fmt.Printf("%s: %d\n", nama, nilai)
	}

	delete(nilaiMahasiswa, "Ani")
	fmt.Println("\nSetelah Ani dihapus:", nilaiMahasiswa)
	fmt.Println("Jumlah data:", len(nilaiMahasiswa))
}
