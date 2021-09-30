package main

import (
	"fmt"
	"math"
)


func main() {
	// pendeklarasian variabel
	var cx, cy, r, x, y float64

	// scan semua variabel
	fmt.Scanln(&cx, &cy, &r)
	fmt.Scanln(&x, &y)

	// print teks sesuai kondisi
	if posisi(cx, cy, r, x, y) {
		fmt.Println("Anda berada di dalam taman")
	} else {
		fmt.Println("Anda berada di luar taman")
	}
}

func jarak(a, b, c, d float64) float64 {
	// return berdasarkan rumus jarak
	return math.Sqrt(((a - c) * (a - c)) + ((b - d) * (b - d)))
}

func posisi(cx, cy, r, x, y float64) bool {
	// return jarak < jari-jari
	return jarak(x, y, cx, cy) < r
}