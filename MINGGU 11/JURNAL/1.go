// PROGRAM PLAYLIST
package main

import "fmt"

const NMAX = 1000

type lagu struct {
	judul, penyanyi string
	durasi          waktu
}
type waktu struct {
	menit, detik int
}
type playlist [NMAX]lagu

	

func main() {
	var list playlist
	var n int

	buatPlaylist(&list, &n)
	cetakPlaylist(list,n)
}

// mereturn lagu berdasarkan judul dan penyanyi
func cariLagu(T playlist, n int, judul, penyanyi string) bool {
	var i int
	i = 0

	// iterate sampai ketemu
	for i < n {
		if T[i].judul == judul && T[i].penyanyi == penyanyi {
			return true		
		} else {
			i++
		}
	}
	// kalo ga ketemu return false
	return false
}

// mengisi array dengan semua lagu yang diinput
func buatPlaylist(T *playlist, n *int) {
	// buat var temp buat menyimpan lagu sebelum dimasukkan ke array
	var temp lagu

	*n = 0

	// scan pertama
	fmt.Scanln(temp.judul, temp.penyanyi, temp.durasi.menit, temp.durasi.detik)
	
	// iterate terus sampai ketemu #
	for (temp.judul != "#") && (temp.penyanyi != "#") {
		if !(cariLagu(*T, *n, temp.judul, temp.penyanyi)) {
				*&T[*n].judul = temp.judul
				*&T[*n].penyanyi = temp.penyanyi
				*&T[*n].durasi.menit = temp.durasi.menit
				*&T[*n].durasi.detik = temp.durasi.detik
				*n++	
		}
		
		fmt.Scanln(temp.judul, temp.penyanyi, temp.durasi.menit, temp.durasi.detik)
	} 
}

// mencari lagu dengan durasi terpanjang
func terlama(T playlist, n int) int {
	var max, temp, index int
	max = 0

	for i := 0; i < n; i++ {
		temp = T[i].durasi.detik + T[i].durasi.menit * 60
		if temp > max {
			max = temp
			index = i
		}
	}

	return index
}

// mencetak playlist dengan ketentuan di soal 
func cetakPlaylist(T playlist, n int) {
	var longest int
	longest = terlama(T,n)

	for i := 0; i < n; i++ {
		if i == longest {
			fmt.Println("*",T[i].judul,T[i].durasi.menit,"menit",T[i].durasi.detik,"detik")
		} else {
			fmt.Println(T[i].judul)
		}
	}
}