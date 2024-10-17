package main

import "fmt"

func main() {
	var n int

	fmt.Print("Masukkan bilangan bulat non negatif: ")
	fmt.Scan(&n)

	factorial := 1
	for i := 1; i <= n; i++ {
		factorial *= i
	}

	fmt.Printf("Faktorial dari %d adalah: %d\n", n, factorial)
}
