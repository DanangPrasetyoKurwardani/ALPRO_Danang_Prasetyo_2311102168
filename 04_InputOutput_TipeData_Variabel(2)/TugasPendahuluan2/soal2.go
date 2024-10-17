package main

import (
	"fmt"
)

func main() {
	var jamKerjaPerMinggu, upahPerJam float64

	fmt.Print("Masukkan jumlah jam kerja per minggu: ")
	fmt.Scan(&jamKerjaPerMinggu)

	fmt.Print("Masukkan upah per jam: ")
	fmt.Scan(&upahPerJam)

	jamNormal := 40.0
	jamLembur := 0.0

	if jamKerjaPerMinggu > jamNormal {
		jamLembur = jamKerjaPerMinggu - jamNormal
	} else {
		jamNormal = jamKerjaPerMinggu
	}

	gajiMingguan := (jamNormal * upahPerJam) + (jamLembur * 1.5 * upahPerJam)
	gajiBulanan := gajiMingguan * 4

	fmt.Printf("Total gaji bulanan Anda adalah: %.2f\n", gajiBulanan)
}
