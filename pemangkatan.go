package main

import "fmt"

func main() {
	var dasar, pangkat, result int
	fmt.Print("Masukkan bilangan dasar : ")
	fmt.Scan(&dasar)
	fmt.Print("Masukkan bilangan pangkat : ")
	fmt.Scan(&pangkat)

	// Inisialisasi hasil dengan 1 (karena bilangan apapun dipangkatkan 0 hasilnya 1)
	result = 1

	// Menghitung base^exponent menggunakan perkalian berulang
	for i := 1; i <= pangkat; i++ {
		result *= dasar
	}

	// Mencetak hasil pemangkatan
	fmt.Printf("Hasil dari %d pangkat %d adalah: %d\n", dasar, pangkat, result)
}
