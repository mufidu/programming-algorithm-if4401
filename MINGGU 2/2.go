package main

import "fmt"

func main() {
	var r, luasLingkaran float64

    fmt.Scanln(&r)
    luasLingkaran = r *22/7
	
    fmt.Println("Luas lingkaran dengan jari-jari =", r, "adalah", luasLingkaran)
}