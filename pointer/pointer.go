package main

import "fmt"

func main() {
	angka := 10
	fmt.Println("Nilai awal:", angka)

	ptr := &angka
	fmt.Println("Alamat memori:", ptr)
	fmt.Println("Nilai melalui pointer:", *ptr)

	*ptr = 20
	fmt.Println("Nilai setelah diubah:", angka)
}
