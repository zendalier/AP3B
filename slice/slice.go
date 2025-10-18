package main

import "fmt"

func main() {
	var arrInt [5]int
	arrInt[0] = 10
	arrInt[1] = 20
	arrInt[2] = 30
	arrInt[3] = 40
	arrInt[4] = 50

	fmt.Println("Data dalam Array: ", arrInt)

	subArrInt := arrInt[1:3]
	fmt.Println("Sub array: ", subArrInt)

	buah := []string{"Apel", "Jeruk", "Mangga"}
	fmt.Println("\nSlice buah:", buah)

	buah = append(buah, "Pisang")
	buah = append(buah, "Durian", "Semangka")
	fmt.Println("Setelah append:", buah)

	fmt.Println("\nIterasi slice:")
	for i, nama := range buah {
		fmt.Printf("%d. %s\n", i+1, nama)
	}

	buah = append(buah[:2], buah[3:]...)
	fmt.Println("\nSetelah hapus index 2:", buah)
}
