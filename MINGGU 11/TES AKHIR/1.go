// PROGRAM MINMAX
package main

import "fmt"

func main() {
	var jumlah, dicari int
	var nums[100]int

	// scan baris pertama
	fmt.Scanln(&jumlah, &dicari)

	// scan baris kedua
	for i := 0; i < jumlah; i++ {
		fmt.Scan(&nums[i])
	}

	// cek ganjil genap, print hasilnya
	if dicari % 2 == 0 {
		fmt.Print(genap(nums,jumlah,dicari))
	} else {
		fmt.Println(ganjil(nums, jumlah, dicari))
	}
}

// mencari indeks dengan angka terkecil
func min(nums [100]int, jumlah int) int {
	var temp, indeks int
	temp = 100000000

	for i := 0; i < jumlah; i++ {
		if nums[i] < temp {
			temp = nums[i]
			indeks = i
		}
	}

	return indeks
}

// mencari indeks max
func max(nums [100]int, jumlah int) int {
	var temp, indeks int
	temp = -1

	for i := 0; i < jumlah; i++ {
		if nums[i] > temp {
			temp = nums[i]
			indeks = i
		}
	}

	return indeks
}

// fungsi ini dijalankan kalau bilangan yang dicari genap
func genap(nums [100]int, jumlah int, dicari int) string {
	for i := min(nums, jumlah); i <= max(nums, jumlah); i++ {
		if nums[i] == dicari {
			return "ditemukan"
		}
	}
	return "tidak ditemukan"
}

// fungsi ini dijalankan kalau bilangan yang dicari ganjil
func ganjil(nums [100]int, jumlah int, dicari int) string {
	for i := 0; i <= min(nums, jumlah); i++ {
		if nums[i] == dicari {
				return "ditemukan"
		}
	}

	for i := max(nums, jumlah); i < jumlah; i++ {
		if nums[i] == dicari {
				return "ditemukan"
		}
	}

	return "tidak ditemukan"
}

// maaf kak kalo pas dirun langsung keprint "ditemukan", dirun lagi aja. :)