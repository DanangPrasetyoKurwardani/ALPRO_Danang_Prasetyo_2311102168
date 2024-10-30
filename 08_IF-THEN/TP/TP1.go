package main

import "fmt"

func main() {
	var nilai int

	fmt.Print("Nilai Ujian: ")
	fmt.Scan(&nilai)

	if nilai >= 70 {
		fmt.Println("Output: Lulus")
	} else {
		fmt.Println("Output: Tidak Lulus")
	}
}
