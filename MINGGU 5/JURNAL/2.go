package main

import "fmt"

func main() {
	// Pendeklarasian variabel
	var i, n, bobot, sks int
	var nilai string
	var jum_sks, jum int

	// Scan jumlah matkul
	fmt.Scanln(&n)

	// Set value awal ke o
	jum = 0
	jum_sks = 0

	// Iterate sebanyak jumlah mata kuliah
	for i = 1; i <= n; i++ {
		fmt.Scanln(&nilai, &sks)
		for !(sks > 0 && nilai == "A" || nilai == "B" || nilai == "C" || nilai == "D" || nilai == "E") {
			fmt.Scanln(&nilai, &sks)
			if nilai == "A" {
				bobot = 4
			} else if nilai == "B" {
				bobot = 3
			} else if nilai == "C" {
				bobot = 2
			} else if nilai == "D" {
				bobot = 1
			} else {
				bobot = 0
			}
			jum = jum + bobot * sks
			jum_sks = jum_sks + sks
		}
	}

	// Output IPS
	print(jum/jum_sks)
}