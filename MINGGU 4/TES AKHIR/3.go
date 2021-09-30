package main

import (
	"fmt"
	"math/rand"
)


func main() {
	var seed int64
	var tebak, dadang, dadu int
	
	// scan semua
	fmt.Print("Angka rahasia: ")
	fmt.Scanln(&seed)
	fmt.Print("Anda: ")
	fmt.Scanln(&tebak)
	
	dadang = rand.Intn(6) + 1
	fmt.Println("Dadang: ", dadang)

	dadu = rand.Intn(6) + 1

	// cari yang menang
	if dadu == tebak {
		fmt.Print("Nilai dadu ", dadu, ", Pemenang adalah anda")
	} else if dadu == dadang {
		fmt.Print("Nilai dadu ", dadu, ", Pemenang adalah Dadang")
	} else if dadu == dadang && dadu == tebak {
		fmt.Print("Nilai dadu ", dadu, ", Seri")
	} else {
		fmt.Print("Nilai dadu ", dadu, ", tidak ada pemenang.")
	}
}