package main

import "fmt"

func main(){
	 var totalBelanjaAwal, diskonBelanja, totalBelanjaAkhir int

	 fmt.Printf("masukan total belanja awal: ")
	fmt.Scan(&totalBelanjaAwal)

	fmt.Printf("masukan diskon: ")
	fmt.Scan(&diskonBelanja)

	totalBelanjaAkhir = totalBelanjaAwal - (totalBelanjaAwal * diskonBelanja / 100)
	
	fmt.Println(totalBelanjaAkhir)
}
