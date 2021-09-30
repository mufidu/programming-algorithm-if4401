// program minimarket
package main

import "fmt"

const NMAX = 100

type barang struct {
	harga int
	jumlah int
}

type belanja [NMAX]barang

func main(){
	// Deklarasi Variabel
	var t belanja
	var n, ha int
	var hp float64
	
	// Lakukan proses masukan
	entri(&t, &n)
	
	// Hitung total belanja
	ha = total(t,n);

	// tentukan apakah mendapatkan promo atau tidak
	if ha > 100000 {
		// lakukan pengurutan
		urut(&t, n)
		// lakukan perhitungan promo
		promo(t,n,&hp,ha)
		// tampilkan keluaran yang diminta
		fmt.Println(ha,hp)
	
	} else {
		// tampilkan keluaran yang diminta
		fmt.Println(ha,ha)
	}
}

func entri(t *belanja, n *int){
/* IS. data belanja telah siap pada piranti masukan
FS. array t berisi sejumlah n harga barang yang dibeli */
	var temp barang
	var i = 0
	fmt.Scanln(&temp.harga, &temp.jumlah)

	for temp.harga != 0 && temp.jumlah != 0 {
		t[i].harga = temp.harga
		t[i].jumlah = temp.jumlah
		*n++
		i++
		fmt.Scanln(&temp.harga, &temp.jumlah)
	}
}

func urut(t *belanja, n int){
/* IS. terdefinisi array t yang berisi n harga barang yang dibeli
FS. array t terurut secara ascending berdasarkan harga barang dengan
selection/insertion sort */
	for i:=0; i<n; i++ {
		var j = i
		
		for j > 0 {
			if t[j-1].harga*t[j-1].jumlah > t[j].harga*t[j].jumlah {
				t[j-1].harga, t[j].harga = t[j].harga, t[j-1].harga
				t[j-1].jumlah, t[j].jumlah = t[j].jumlah, t[j-1].jumlah
			}
			j--
		}
	}
}

func total(t belanja, n int) int {
/* IS. terdefinisi array t yang berisi n harga barang yang dibeli
FS. mengembalikan total harga barang yang terdapat pada array t */	
	var total int
	for i:=0; i<n; i++ {
		total += t[i].harga * t[i].jumlah
	}

	return total
}

func promo(t belanja, n int, hp *float64, ha int){
/* IS. terdefinisi array t yang berisi n harga barang yang dibeli dan terurut
ascending
FS. hp berisi total harga setiap barang setelah dikurangi dengan diskonnya*/
	*hp = float64(ha)
	for i:=0; i<n; i++ {
		*hp -= float64((t[i].harga * t[i].jumlah) * (i+1) / 100)
	}
}