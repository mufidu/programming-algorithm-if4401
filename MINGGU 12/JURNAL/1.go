package main

import "fmt"

type pemain struct {
	nama        string
	gol, assist int
}

type list [100]pemain

func main() {
	var n int
	var list list

	// input total pemain ada berapa
	fmt.Scanln(&n)

	// iterate buat masukin ke data pemain
	for i:=0; i<n; i++ {
		fmt.Scanln(&list[i].nama, &list[i].gol, &list[i].assist)
	}

	// print berdasarkan urutan
	for i:=0; i<n; i++ {
		var temp, maxgol, maxass int
		maxgol = 0;

		// iterate cari gol maks
		for i:=0; i<n; i++ {
			if list[i].gol > maxgol {
				temp = i
				maxgol = list[i].gol
				maxass = list[i].assist
			} else if list[i].gol == maxgol {
				if list[i].assist > maxass {
					temp = i
					maxgol = list[i].gol
					maxass = list[i].assist
				}
			}
		}

		fmt.Println(list[temp].nama, list[temp].gol, list[temp].assist)
		// hapus nilai gol maks buat iterate selanjutnya
		list[temp].gol = 0;
	}
}