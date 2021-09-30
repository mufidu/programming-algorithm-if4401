package main

import (
	"fmt"
	"math"
)


func main() {
	var xr, yr, r float64
	var xr2, yr2, r2 float64
	var x, y float64

	// scan semuanya
	fmt.Scanln(&xr, &yr, &r)
	fmt.Scanln(&xr2, &yr2, &r2)
	fmt.Scanln(&x, &y)

	// jalankan function
	cariTeks(xr, yr, r, xr2, yr2, r2, x, y)

}

// function untuk menghitung jarak antar dua titik
func jarak(a, b, c, d float64) float64 {
	return math.Sqrt(((a - c) * (a - c)) + ((b - d) * (b - d)))
}

// function untuk mencari apakah di dalam lingkaran atau tidak
func lingkaran(xr, yr, r, x, y float64) bool {
	return jarak(x, y, xr, yr) < r
}

// function untuk mencetak posisi titik
func cariTeks(xr, yr, r, xr2, yr2, r2, x, y float64) {
	if lingkaran(xr, yr, r, x, y) && lingkaran(xr2, yr2, r2, x, y) {
		fmt.Println("di dalam lingkaran 1 dan 2")
	} else if lingkaran(xr, yr, r, x, y) {
		fmt.Println("di dalam lingkaran 1")
	} else if lingkaran(xr2, yr2, r2, x, y){
		fmt.Println("di dalam lingkaran 2")
	} else {
		fmt.Println("di luar lingkaran 1 dan 2")
	}
}

// Muhammad Mufid Utomo - 1301204441