package main

import "fmt"

func main() {
	var n int
	var jumlah int

	fmt.Print("Masukkan bilangan bulat positif n: ")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		jumlah += i
	}

	fmt.Println("Jumlah dari 1 hingga", n, "adalah:", jumlah)
}
