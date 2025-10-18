package main 
import "fmt"

func main() {
    // Deklarasi array dengan ukuran tetap
    var angka [5]int
    angka[0] = 10
    angka[1] = 20
    angka[2] = 30
    angka[3] = 40
    angka[4] = 50
    
    fmt.Println("Array angka:", angka)
    fmt.Println("Panjang array:", len(angka))
    
    // Deklarasi dan inisialisasi langsung
    buah := [4]string{"Apel", "Jeruk", "Mangga", "Pisang"}
    fmt.Println("\nArray buah:", buah)
   
}
