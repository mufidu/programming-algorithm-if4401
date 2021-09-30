package main

type Buku struct {
	judul   string
	penulis string
	tahun   int
}

type TabBuku = [5]Buku

func main() {
	var kardus TabBuku
	var atas int
	var ambil Buku

	// {1. Buatlah sebuah kardus kosong, di mana atas bernilai -1}
	atas = -1

	// {2. Tambahkan 4 buku seperti contoh gambar, data tahun dan penulis bebas}
	kardus[0].judul = "book1"
	kardus[0].penulis = "auth1"
	kardus[0].tahun = 2020

	kardus[1].judul = "book2"
	kardus[1].penulis = "auth2"
	kardus[1].tahun = 2020

	kardus[2].judul = "book3"
	kardus[2].penulis = "auth3"
	kardus[2].tahun = 2020

	kardus[3].judul = "book4"
	kardus[3].penulis = "auth4"
	kardus[3].tahun = 2020

	// {3. Cari buku dengan judul “C”, harusnya yang tampil adalah K dan ZZ, KETEMU}

	// {4. Tambahkan buku sampai kardus penuh}
	tambahBuku()

	// {5. Cari buku dengan judul yang tidak terdapat pada kardus tersebut.}
}

func tambahBuku(kardus *TabBuku, atas *int) {

}

func ambilBuku(kardus *TabBuku, atas *int, ambil *Buku) {

}
