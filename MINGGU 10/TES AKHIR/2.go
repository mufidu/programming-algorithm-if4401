package main

import "fmt"

const NMAX = 1000000

type data [NMAX]int

func main() {
	/* buatlah kode utama yang membaca baris pertama (n dan k). kemudian data diisi
oleh prosedur isiArray(n), dan pencarian oleh fungsi posisi(n,k), dan setelah itu
output dicetak.*/

	var n int
	var k int
	var angka data
	fmt.Scan(&n)
	fmt.Scanln(&k)

	isiArray(n, &angka)
	fmt.Print(posisi(n, k, angka))

}

func isiArray(n int, angka *data) {
	/* IS. Data n sudah siap pada piranti masukan.
	   FS. Array data berisi n (<=NMAX) bilangan */
	for i:= 0; i < n; i++ {
		fmt.Scan(angka[i])
	}
}

func posisi(n, k int, angka data) int {
	/* mengembalikan posisi k dalam array data dengan n elemen. Posisi dimulai dari
	   posisi 0. Jika tidak ada kembalikan -1 */
}