package main

import "fmt"

func main() {
	var dasar, pangkat, result int
	fmt.Print("Masukkan bilangan dasar : ")
	fmt.Scan(&dasar)
	fmt.Print("Masukkan bilangan pangkat : ")
	fmt.Scan(&pangkat)

	result = 1
	for i := 1; i <= pangkat; i++ {
		result *= dasar
	}
	fmt.Printf("Hasil dari %d pangkat %d adalah: %d\n", dasar, pangkat, result)
}
