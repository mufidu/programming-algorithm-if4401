package main

import (
	"fmt"
)


func main() {
	var n, bil, digits, angka int
	// scan n
	fmt.Scan(&n)

	// loop sebanyak n kali
	digits = 1
	for digits < n + 1 {
		fmt.Scanln(&bil)
		if bil >= 0 && bil < 10 {
			angka += bil*10^digits
			digits++
			fmt.Print(digits)
		}
	}

	// print angkanya
	fmt.Print(angka)
}