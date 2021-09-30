// PROGRAM konversi_desimal_ke_biner

package main

import "fmt"

func main() {
	// pendeklarasian varibel
	var biner string
	var desimal int

	// scan desimalnya
	fmt.Scanln(&desimal)

	// masukkan ke func
	biner = Des2Bin(desimal)

	// print hasil biner
	fmt.Println(biner)
}

// function yang menmberi nilai ke result dan remainder
func Division(a, b int, result, remainder *int) {
	*result = a / b
	*remainder = a % b
}

// function yang mengubah int menjadi string (0 dan 1)
func Num2Str(x int) string {
	if x == 0 {
		return "0"
	} else if x == 1 {
		return "1"
	} else {
		return ""
	}
}

// function yang mengubah desimal menjadi biner berbentuk string
func Des2Bin(desimal int) string {
	// pendeklarasian var lokal
	var hasil string
	var rDiv, rMod int

	hasil = ""

	for desimal > 0 {
		Division(desimal, 2, &rDiv, &rMod)
		desimal = rDiv
		hasil = Num2Str(rMod) + hasil
	}

	return hasil
}