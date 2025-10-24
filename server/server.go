package server

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Mahasiswa struct {
	Name  string `json:"name"`
	Kelas string `json:"kelas"`
	NPM   int    `json:"npm"`
}

// POST Request
func acceptMahasiswa(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Only POST method is supported", http.StatusMethodNotAllowed)
		return
	}
	var mahasiswa Mahasiswa

	err := json.NewDecoder(r.Body).Decode(&mahasiswa)
	if err != nil {
		http.Error(w, "Invalid Input", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Nama: %s, Kelas: %s, NPM: %d", mahasiswa.Name, mahasiswa.Kelas, mahasiswa.NPM)
}

// GET Request
func sayHello(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Only GET method is supported", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintf(w, "Hello World")
}

func main() {
	http.HandleFunc("/", sayHello)
	http.HandleFunc("/name", acceptMahasiswa)
	fmt.Println("Server started at localhost:8080")
	http.ListenAndServe(":8080", nil)
}
