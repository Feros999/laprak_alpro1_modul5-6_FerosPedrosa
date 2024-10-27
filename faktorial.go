package main

import "fmt"

func main() {
	var n, result int
	fmt.Print("Masukkan bilangan bulat non-negatif: ")
	fmt.Scan(&n)

	result = 1
	for i := 1; i <= n; i++ {
		result *= i
	}
	fmt.Printf("Faktorial dari %d adalah: %d\n", n, result)
}
