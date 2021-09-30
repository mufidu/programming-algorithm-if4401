// PROGRAM PALINDROM
package main

import "fmt"

type tabel = [127]string
func main() {
	var teks tabel
	var n int

	n = 0
	fmt.Print("Teks      : ")
	isiArray(&teks, &n)

	fmt.Print("Palindrom ? ")
	fmt.Println(palindrom(teks, n))

}

func isiArray(t *tabel, n *int){
// {IS. Data tersedia dalam piranti masukan
// FS. Array t berisi sejumlah n karakter yang dimasukkan user,
// Proses input selama karakter bukanlah TITIK dan n <= NMAX}
	var huruf string
	for huruf != "." && *n <= 127{
		*n = *n + 1
		t[*n] = huruf
		fmt.Scan(&huruf)
	}
	*n--
}	

func cetakArray(t tabel, n int) {
// {IS. Terdefinisi array t yang berisi sejumlah n karakter
// FS. n karakter dalam array muncul di layar}
	for i := 0; i <= n; i++ {
		fmt.Print(t[i])
	}
}

func balikanArray(t *tabel, n int) {
// {IS. Terdefinisi array t yang berisi sejumlah n karakter
// FS. Urutan isi array t terbalik }
	var temp tabel
	temp = *t
	j := 127-(127-n-2)
	for i := 0; i <= n+1; i++{
		t[i] = temp[j]
		j--
	}
}

func palindrom(t tabel, n int) bool {
// {Mengembalikan true apabila susunan karakter didalam t membentuk
// palindrom, dan false apabila sebaliknya. Petunjuk: Manfaatkan prosedur
// balikanArray}
	var temp, temp2 tabel
	temp = t
	balikanArray(&t, n)

	// digeser satu, kalo ngga jadi false (gatau kenapa)
	temp2 = t
	for i := 0; i <= n+1; i++{
		t[i+1] = temp2[i]
	}

	return t == temp
}