package main

import (
	"fmt"
)

func main() {
	// Pendeklarasian variabel
	var e1, e2 float64
	var netral, pos int
	var aman bool

	// Set nilai awal
	e1 = 1
	e2 = 1
	aman = true
	netral = 0
	pos = 0

	// Scan energi
	for aman {
		fmt.Scanln(&e1, &e2)
		aman = e1 + e2 >= 0
		// Cek apakah netral atau positif
		if aman {
			if e1 + e2 == 0 {
			netral = netral + 1
			} else if e1 + e2 > 0 {
			pos = pos + 1
			}
		}
	}

	// Output total netral dan positif
	fmt.Println("Netral: ", netral)
	fmt.Println("Positive: ", pos)
}