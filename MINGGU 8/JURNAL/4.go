package main

import "fmt"

func main() {
	var plat string
	var bdg, jkt, dll int

	plat = "D"
	// bdg dikurang 1 karena nilai awalnya D, biar netral 0 dulu
	bdg = -1
	jkt = 0
	dll = 0

	// selama input bukan ".", minta terus
	for plat != "." {

		// kelompokin plat
		hitungNopol(plat, &bdg, &jkt, &dll)

		fmt.Scan(&plat)
	}

	fmt.Println(bdg, jkt, dll)
}

func hitungNopol(noPol string, nBandung,nJakarta,nLainnya *int) {

	// {I.S. terdefinisi noPol suatu kendaraan
	// F.S. nilai nBandung, nJakarta, atau nLainnya bertambah sesuai noPol yang dibaca}
	if noPol == "D" || noPol == "Z" {
			*nBandung = *nBandung + 1
		} else if noPol == "B" || noPol == "F" {
			*nJakarta = *nJakarta + 1
		} else {
			*nLainnya = *nLainnya + 1
		}
}
