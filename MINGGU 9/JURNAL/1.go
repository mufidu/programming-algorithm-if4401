//PROGRAM PERTANDINGAN BOLA {RATA2 DARI ARRAY}
package main

import "fmt"

type tabGol [500]int

func main() {
	// setiap tim diberi satu array
	var tim1 tabGol
	var tim2 tabGol
	var tim3 tabGol
	var n1, n2, n3, r1, r2, r3 int

	n1, n2, n3 = 0, 0 ,0
	inputData(&tim1, &n1)
	inputData(&tim2, &n2)
	inputData(&tim3, &n3)

	r1 = rataan(tim1, n1)
	r2 = rataan(tim2, n2)
	r3 = rataan(tim3, n3)

	fmt.Println(r1)
	fmt.Println(r2)
	fmt.Println(r3)
}

func inputData(t *tabGol, n *int){

	// "menang" sebagai var temp
	var menang int
	menang = 1
	for menang > -1 && *n <= 500{
		*n = *n + 1
		t[*n] = menang
		fmt.Scan(&menang)
	}
	*n--
}	

func rataan(t tabGol, n int) int {

	// "total" untuk menampung jumlah gol
	var total int
	total = 0
	for i := 0; i < len(t); i++ {
		total = total + t[i]
	}

	return total / n
}







