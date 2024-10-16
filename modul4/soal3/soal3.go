package main

import (
	"fmt"
	"math"
)

func main() {
	var ax, ay, bx, by, cx, cy float64

	fmt.Printf("masukan koordinat titik A (x, y): ")
	fmt.Scan(&ax, &ay)

	fmt.Printf("masukan koordinat titik B (x, y): ")
	fmt.Scan(&bx, &by)

	fmt.Printf("masukan koordinat titik C (x, y): ")
	fmt.Scan(&cx, &cy)

	ab := math.Sqrt(math.Pow(bx-ax, 2) + math.Pow(by-ay, 2))
	bc := math.Sqrt(math.Pow(cx-bx, 2) + math.Pow(cy-by, 2))
	ac := math.Sqrt(math.Pow(cx-ax, 2) + math.Pow(cy-ay, 2))

	var longestSide float64
	if ab >= bc && ab >= ac {
		longestSide = ab
	} else if bc >= ab && bc >= ac {
		longestSide = bc
	} else {
		longestSide = ac
	}

	fmt.Printf("Panjang sisi terpanjang adalah: %.2f\n", longestSide)
}