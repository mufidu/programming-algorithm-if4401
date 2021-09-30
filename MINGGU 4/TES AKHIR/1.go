package main

import "fmt"

func main() {
	var a, faktor int

	// scan angka
	fmt.Scanln(&a)
	
	faktor = 0
	// iterate dari 1 sampai angka, print jika angka yang diiterate adalah faktor
	for i := 1; i <= a; i++ {
		if a % i == 0 {
			fmt.Print(i, " ")
			faktor++
		}
	}

	// tentukan apakah oleole atau bukan
	fmt.Println("")
	if faktor > 2 {
		fmt.Print("Bukan Oleole")
	} else {
		fmt.Print("Oleole")
	}
}