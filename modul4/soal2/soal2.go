package main

import "fmt"

func main(){
    var BMI, tinggiBadan, beratBadan float64

    fmt.Printf("masukan nilai BMI: ")
    fmt.Scan(&BMI)

    fmt.Printf("masukan tinggi badan (dalam meter): ")
    fmt.Scan(&tinggiBadan)

    beratBadan = BMI * tinggiBadan * tinggiBadan

    fmt.Printf("Berat badan seseorang adalah: %.2f kg\n", beratBadan)
}