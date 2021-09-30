package main

import "fmt"

func main() {
	// Pendeklarasian variabel
	var goal, i, tendang int

	// Set nilai i, goal, dan tendang awal
	i = 1
	tendang = 1
	goal = 0

	for tendang > 0 {
		fmt.Print("Tendangan ke-", i, ": ")
		fmt.Scanln(&tendang)
		
		// Cek apakah gol
		if tendang % 4 == 0 {
			goal = goal + 1
		} 

		// Increment i
		i++
	}

	// Tentukan pemenang
	if goal > i/2 {
		fmt.Println("Penyerang menang")
	} else if goal == i/2 {
		fmt.Println("Draw")
	} else if goal < i/2 {
		fmt.Println("Kiper menang")
	}

	// Output skor akhir
	fmt.Println("Skor gol: ", goal-1, "dari ", i-2, "tendangan")
}