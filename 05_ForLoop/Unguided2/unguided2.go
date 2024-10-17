package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Print("Masukkan jumlah kerucut: ")
	fmt.Scan(&n)

	for i := 0; i < n; i++ {
		var r, h float64
		fmt.Print("Masukkan jari-jari alas dan tinggi kerucut (r h): ")
		fmt.Scan(&r, &h)

		volume := (1.0 / 3.0) * math.Pi * math.Pow(r, 2) * h

		fmt.Println(volume)
	}
}
