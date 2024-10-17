package main

import (
	"fmt"
	"math"
)

func main() {
	var r float64
	const phi = 3.141592653589793

	fmt.Print("Masukkan r lingkaran: ")
	fmt.Scan(&r)

	luas := phi * math.Pow(r, 2)

	keliling := 2 * phi * r

	fmt.Printf("Luas lingkaran dengan r %.2f adalah: %.2f\n", r, luas)
	fmt.Printf("Keliling lingkaran dengan r %.2f adalah: %.2f\n", r, keliling)
}
