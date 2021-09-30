package main

import "fmt"

func main() {
	// Pendeklrasian variabel
	var n, i int
	var mi, mo, ma int
	var jum_ma int
	var tahun int

	// Scan n, set jumlah mahasiswa dan tahun ke 0
	fmt.Scanln(&n)
	jum_ma = 0
	tahun = 0

	// Iterate n, scan mi mo ma di tiap iterasi
	for i = 1; i <= n; i++ {
		fmt.Scanln(&mi, &mo, &ma)
		if (jum_ma + mi - mo != ma) && tahun == 0 {
			tahun = i
		} 
		jum_ma = ma
	}

	// Output tahun
	fmt.Println(tahun)

}