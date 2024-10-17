package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	target := rand.Intn(100) + 1
	var guess int

	fmt.Println("Selamat datang di permainan tebak angka!")
	fmt.Println("Saya telah memilih angka antara 1 hingga 100.")
	fmt.Println("Anda memiliki 5 kesempatan untuk menebak angka tersebut.")

	for attempts := 1; attempts <= 5; attempts++ {
		fmt.Print("Tebakan Anda (kesempatan ke-", attempts, "): ")
		fmt.Scan(&guess)

		if guess < target {
			fmt.Println("Tebakan Anda terlalu kecil.")
		} else if guess > target {
			fmt.Println("Tebakan Anda terlalu besar.")
		} else {
			fmt.Println("Selamat! Anda telah menebak angka yang benar:", target)
			return
		}
	}

	fmt.Println("Maaf, Anda telah kehabisan kesempatan. Angka yang benar adalah:", target)
}
