package main

import "fmt"

func main() {
	// Pendeklarasian variabel
	var winner, player, temp string
	var i, ronde int
	var nilai, answer int

	// Set pemain awal
	winner = "A"
	player = "B"
	ronde = 1

	// Input pertama dari A
	fmt.Println("Ronde", ronde, ":")
	fmt.Print(winner, "- masukkan angka rahasia: ")
	fmt.Scanln(&nilai)

	// Main terus selama nilai != -101
	for nilai != -101 {
		i = 1
		fmt.Print(player, "- masukkan angka tebakan ke-", i, ":")
		fmt.Scanln(&answer)
		for i < 3 && answer != nilai {
			i = i + 1
			fmt.Print(player, "- masukkan angka tebakan ke-", i, ":")
			fmt.Scanln(&answer)
		}

		// Tukar pemain kalau menang
		if nilai == answer {
			temp = winner
			winner = player
			player = temp
		}

		fmt.Println(winner, "adalah pemenangnya")
		ronde = ronde + 1
		fmt.Println("Ronde", ronde, ":")
		fmt.Print(winner, "- masukkan angka rahasia:")
		fmt.Scanln(&nilai)
	}
	
	// game selesai
	fmt.Println("Permainan selesai")
}