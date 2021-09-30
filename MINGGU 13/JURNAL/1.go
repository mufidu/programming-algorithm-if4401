// PROGRAM URUT SELECTION SORT
package main

import "fmt"
const NMAX = 100
type mahasiswa struct{
	nama string
	tinggi int
}
type dataMhs [NMAX]mahasiswa

func main(){
	var jumlah int
	var data dataMhs

	bacaData(&data, &jumlah)
	urutData(&data, jumlah)
	cetakData(data, jumlah)
}

func bacaData(t *dataMhs, n *int){
/* IS. n data mahasiswa telah siap pada piranti masukan
FS. menerima input n dan array t berisi n data tinggi mahasiswa */
	fmt.Scanln(n)
	for i:=1; i<=*n; i++ {
		fmt.Scanln(&t[i].nama, &t[i].tinggi)
	}
}

func cetakData(t dataMhs, n int){
/* IS. terdefinisi sebuah array t yang berisi n data tinggi mahasiswa
FS. menampilkan array t ke layar monitor */
	for i:=1; i<=n; i++ {
		fmt.Println(t[i].nama, t[i].tinggi)
	}
}

func urutData(t *dataMhs, n int){
/* IS. terdefinisi sebuah array t yang berisi n data tinggi mahasiswa
FS. array t terurut ascending berdasarkan tinggi dengan algoritma selection sort
*/
	for i := 0; i < n; i++ {
        var min int
		var temp int
		var temp2 string
		min = i
        for j := i; j < n; j++ {
            if t[j].tinggi < t[min].tinggi {
                min = j
            }
        }

        temp = t[i].tinggi
		t[i].tinggi = t[min].tinggi
		t[min].tinggi = temp
        
		temp2 = t[i].nama
		t[i].nama = t[min].nama
		t[min].nama = temp2
    }
}