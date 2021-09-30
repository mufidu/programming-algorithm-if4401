// PROGRAM WISUDA (SELEKSI)
package main

import "fmt"

func main() {
	var n, m int
	var daftar, wisuda [107]string
	var nim string

	// lakukan proses input untuk baris pertama
	fmt.Scan(&n)
	for i := 0; i < n; i++ {
		fmt.Scan(&nim)
		daftar[i] = nim
	}
	// lakukan proses input untuk baris kedua
	fmt.Scan(&m)
	for i := 0; i < m; i++ {
		fmt.Scan(&nim)
		wisuda[i] = nim
	}
	// panggil prosedur update kelulusan di sini
	updateKelulusan(&daftar, wisuda, &n, m)
	// lakukan proses output array daftar di sini
	if n > 0 {
		fmt.Print(n)
		for i := 0; i < n; i++ {
			fmt.Print(" ",daftar[i])
		}
	} else {
		fmt.Println("0")
	}
}

func posisi(tab [107]string, n int, x string) int {
	/*mengembalikan indeks dari x pada array tab, apabila x ditemukan di dalam tab yang
	berisi n buah nilai, -1 apabila tidak ditemukan*/
	
	for i := 0; i < n; i++ {
		if tab[i] == x {
			return i
		}
	}

	return -1
}



func hapus(tab *[107]string, n *int, x string) {
	/*IS. terdefinisi sebuah array tab yang berisi n buah nim mahasiswa, dan x adalah nim
	mahasiswa yang wisuda
	FS. apabila x ditemukan pada tab, maka nim dihapus, seluruh elemen setelahnya
	bergeser, dan n berkurang 1. tab dan n tidak berubah apabila x tidak di temukan*/

	// for i := 0; i < *n; i++ {
	// 	if tab[i] == x {
	// 		tab[i] = ""
	// 		for j := i + 1; j < *n; j++ {
	// 			tab[j-1] = tab[j]
	// 		}  
	// 		*n--
	// 	}
	// } 
	// harus pake posisi()

	var index int
	index = posisi(*tab, *n, x)
	tab[index] = ""
	for i := index + 1; i < *n; i++ {
		tab[i-1] = tab[i]
	}
	*n--

}

func updateKelulusan(mhs *[107]string, wisuda [107]string, n *int, m int) {
	/* IS. terdefinisi array mhs dan wisuda yang berisi sejumlah n dan m nim mahasiswa
	FS. seluruh nim mahasiswa wisuda dihapus dari array mhs, nilai n terupdate */
	for i := 0; i < m; i++ {
		hapus(mhs, n, wisuda[i])
	}
}