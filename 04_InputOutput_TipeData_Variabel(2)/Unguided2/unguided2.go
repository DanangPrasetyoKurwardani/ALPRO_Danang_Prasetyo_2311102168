package main

import "fmt"

func main() {
	var bmi, tinggi float64

	fmt.Print("Masukkan nilai BMI: ")
	fmt.Scan(&bmi)

	fmt.Print("Masukkan tinggi badan (dalam meter): ")
	fmt.Scan(&tinggi)

	beratBadan := bmi * (tinggi * tinggi)

	fmt.Printf("Berat badan Anda adalah: %.2f kg\n", beratBadan)
}
