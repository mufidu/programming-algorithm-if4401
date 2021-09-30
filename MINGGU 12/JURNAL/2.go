package main

import "fmt"

const NMAX = 1000000
// tipe bentukan partai
type partai struct {
	nama int
	suara int
}

// tipe array of partai dengan kapasitas NMAX
type tabPartai [NMAX]partai

func main(){
// deklarasi variabel
	var i int
	var parties tabPartai

	i = 0;
// lakukan proses input suara secara berulang di sini, simpan ke dalam
// array p, sehingga terdapat array p yang berisi hasil peroleh suara n
// partai.
	for i != -1 {
		fmt.Scan(&i)

		// cek sudah ada atau belum
		if i == -1 {
			continue
		} else if parties[i].nama == i {
			parties[i].suara++
		} else {
			parties[i].nama = i
			parties[i].suara = 1
		}
	}

// lakukan proses pengurutan dengan insertion sort descending berdasarkan
// jumlah suara yang diperoleh
	

	for i:=0; i < NMAX; i++ { 
		var temp, max int
		max = 0

		for i:=0; i < NMAX; i++ {
			if parties[i].suara > max {
				max = parties[i].suara
				temp = i
			}
		}
		
		if parties[temp].suara != 0 {
			fmt.Printf(`%d(%d) `, parties[temp].nama, parties[temp].suara)
			parties[i].suara = 0
		}
	}
}