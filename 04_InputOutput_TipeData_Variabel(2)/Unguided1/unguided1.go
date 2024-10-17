package main

import "fmt"

func main() {
	var totalBelanja, diskon int

	fmt.Print("Masukkan total belanja: ")
	fmt.Scan(&totalBelanja)

	fmt.Print("Masukkan diskon: ")
	fmt.Scan(&diskon)

	potongan := totalBelanja * diskon / 100
	totalSetelahDiskon := totalBelanja - potongan

	fmt.Println("Total belanja setelah diskon adalah:", totalSetelahDiskon)
}
