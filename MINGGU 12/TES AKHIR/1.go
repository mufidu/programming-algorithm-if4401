// program dataPertumbuhan
package main

import "fmt"


const NMAX = 100

type provinsi struct {
	nama string
	populasi int
	tumbuh float32
}

type dataprovinsi struct {
	tabel [NMAX]provinsi
	nProvinsi int
}

func main() {
	var provinces dataprovinsi
	var prov1 provinsi
	var nama, string1, string2 string
	var pop int
	var tumb float32
	var i int

	fmt.Println("Data pertumbuhan provinsi: ")

	// input pertama
	i = 0
	fmt.Scanln(&nama, &pop, &tumb)

	// input terus sampai namanya NONE
	for nama != "NONE" {
		provinces.tabel[i].nama = nama
		provinces.tabel[i].populasi = pop
		provinces.tabel[i].tumbuh = tumb
		provinces.nProvinsi++
		i++
		
		fmt.Scanln(&nama, &pop, &tumb)
	}

	// cari provinsi
	fmt.Print("Nama provinsi? ")
	fmt.Scanln(&string1)
	prov1 = cariProvinsi(provinces, string1)
	// sebutkan atribut provinsi
	fmt.Println(prov1.nama, "   ", prov1.populasi, "   ", prov1.tumbuh)

	// cari provinsi 2
	fmt.Print("Prediksi populasi tahun depan provinsi? ")
	fmt.Scanln(&string2)
	// kasih prediksi
	fmt.Println("Populasi Provinsi", string2, "tahun depan:", prediksi(provinces, string2))

	// print sortingan
	fmt.Println("Urutan data pertumbuhan provinsi berdasarkan populasi terurut menurun: ")
	urutData(&provinces)
}

func cariProvinsi(data dataprovinsi, nama string) provinsi {
	// iterate sampai ketemu
	for i:=0; i < data.nProvinsi; i++ {
		if data.tabel[i].nama == nama {
			return data.tabel[i]
		}
	}

	// kasih returnan biar ga error
	var biargaerror provinsi
	return biargaerror
}

func prediksi(data dataprovinsi, nama string) int {
	// iterate sampai ketemu
	for i:=0; i < data.nProvinsi; i++ {
		if data.tabel[i].nama == nama {
			return data.tabel[i].populasi + data.tabel[i].populasi * int(data.tabel[i].tumbuh)
		}
	}

	// returnan kalo ga ketemu
	return -1
}

// print data yang terurut
func urutData(data *dataprovinsi) {
	// iterate sampai keprint semua
	for i:=1; i<*&data.nProvinsi; i++ {
		var max, temp int
		max = 0

		// iterate sampai ketemu yang terbesar
		for i:=1; i<*&data.nProvinsi; i++ {
			if *&data.tabel[i].populasi > max {
				temp = i
				max = *&data.tabel[i].populasi
			}
		}

		fmt.Println(*&data.tabel[temp].nama, "   ", *&data.tabel[temp].populasi, "   ", *&data.tabel[temp].tumbuh)
		*&data.tabel[temp].populasi = 0
	}
}

// kadang output urutData kurang satu provinsi