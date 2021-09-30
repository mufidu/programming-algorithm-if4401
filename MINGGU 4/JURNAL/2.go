package main

import "fmt"

func main() {
	var jumlah, nilai, besar, mhs int

	// scan jumlah nilai
	fmt.Scanln(&jumlah)

	besar = 0
	mhs = 0

	// iterate sebanyak jumlah nilai
	for i := 0; i < jumlah; i++ {
		fmt.Scan(&nilai)
		if nilai > besar {
			besar = nilai
			mhs = 1
		}else if nilai == besar {
			mhs++
		}
	}

	fmt.Println(`Nilai terbesar adalah`, besar, `yang diperoleh`, mhs, `mahasiswa`)
}