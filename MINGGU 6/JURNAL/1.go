package main

import "fmt"

func main() {
	// pendeklarasian variable
	var n, r, p, c int

	// scan inputan
	fmt.Scanln(&n, &r)

	// masukkan nilai permutasi dan kombinasi
	p = permutation(n, r)
	c = combination(n, r)

	// print hasil
	fmt.Println(p, c)
}

func findFactorial(n int) int {
	// mendeklarasikan variabel lokal
	var result int
	result = 1

	// iterate dari i sampai n, tambahkan setiap i ke result
	for i:=1; i <= n; i++ {
		result *= i;
	}

	// return nilai faktorial
	return result
}

func permutation(n, r int) int {
	// menghitung permutasi
	return findFactorial(n) / findFactorial(n - r)
}

func combination(n, r int) int {
	// menghitung kombinasi
	return findFactorial(n) / (findFactorial(r) * findFactorial(n-r))
}