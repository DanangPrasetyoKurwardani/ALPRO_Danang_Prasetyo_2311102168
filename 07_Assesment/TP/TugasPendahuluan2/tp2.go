package main

import "fmt"

func main() {
	var jumlahBarang, totalPoin int
	fmt.Print("Masukkan jumlah barang yang dibeli: ")
	fmt.Scan(&jumlahBarang)

	if jumlahBarang <= 5 {
		totalPoin = jumlahBarang * 10
	} else {
		totalPoin = (5 * 10) + ((jumlahBarang - 5) * 15)
	}

	fmt.Printf("Total poin yang didapat : %d poin\n", totalPoin)
}
