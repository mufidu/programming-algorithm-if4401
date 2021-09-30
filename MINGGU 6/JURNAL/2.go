package main

import "fmt"

func main() {
	// menggunakan float karena bisa mengandung koma
	var x float64

	// scan x
	fmt.Scanln(&x)

	// print hasil function (menggunakan print formatting)
	fmt.Printf("f(%.2f) = %.2f\n", x, f(x))
	fmt.Printf("g(%.2f) = %.2f\n", x, g(x))
	fmt.Printf("h(%.2f) = %.2f\n", x, h(x))
	fmt.Printf("(fogoh)(%.2f) = %.2f\n", x, komposisi(x))
}

func f(x float64) float64 {
	// return sesuai fungsi yang diperlukan
	return x*x
}

func g(x float64) float64 {
	// return sesuai fungsi yang diperlukan
	return x - 2
}

func h(x float64) float64 {
	// return sesuai fungsi yang diperlukan
	return x + 1
}

func komposisi(x float64) float64 {
	// return komposisi
	return f(g(h(x)))
}