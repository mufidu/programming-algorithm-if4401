package main

import "fmt"
import "math/rand"

func main() {
	var angka, rand, tebakan int

	fmt.Print('Angka rahasia: ')
	fmt.Scanln(&angka)
	fmt.Print('Anda: ')
	fmt.Scanln(&tebakan)
	rand = rand.Intln(6) + 1
}