// PROGRAM usaha_doa

package main

import "fmt"

func main() {
	// pendeklarasian variabel
	var total, usaha, nDoa int
	var doaOrtu bool
	var nilai string

	// scan semuanya
	BacaData(&usaha, &nDoa, &doaOrtu, &nilai)

	// hitung poin
	TabungUsahaDoa(usaha, nDoa, &total)
	total = TabungDoaOrtu(doaOrtu, total)
	HasilNilaiAlpro(nilai, &total)

	// print hasil akhir
	fmt.Println(HasilAkhir(total))
}

// function yang menscan semuanya
func BacaData(usaha,jumlahDoa *int, doaOrtu *bool, nilai *string) {
	fmt.Print("Banyaknya usaha? ")
	fmt.Scanln(usaha)
	fmt.Print("Banyaknya doa? ")
	fmt.Scanln(jumlahDoa)
	fmt.Print("Doa orang tua? ")
	fmt.Scanln(doaOrtu)
	fmt.Print("Nilai Algoritma? ")
	fmt.Scanln(nilai)
}

// function untuk menghitung total dari usaha dan doa
func TabungUsahaDoa(usaha, doa int, total *int) {
	*total = usaha + doa
}

// function yang menghitung doa ortu
func TabungDoaOrtu(doa bool, total int) int {
	if doa {
		return total * 2
	} else {
		return total
	}
}

// function untuk mengolah nilai alpro 
func HasilNilaiAlpro(nilai string, total *int) {
	if nilai == "A" {
		*total = *total - 150
	} else if nilai == "B" {
		*total = *total - 130
	} else if nilai == "C" {
		*total = *total - 100
	}
}

// function yang menghitung hasil akhir
func HasilAkhir(poin int) string {
	if poin >= 130 {
		return "Lulus langsung dapat kerja gaji 2 digit"
	} else if poin >= 50 && poin < 130 {
		return "Lulus langsung dapat kerja"
	} else {
		return "Jangan lelah berdoa dan berusaha, tidak ada yang sia-sia dari usaha dan doa"
	}
}