// PROGRAM pilkart_1
package main

import "fmt"

type kotak [500]int

func main() {
	var masuk, i, total, sah int
	var suara kotak
	fmt.Scan(&masuk)
	
	i = 0
	total = 0
	sah = 0
	
	for masuk != 0 && i <= 500 {
		total = total + 1
		if masuk > 0 && masuk < 21 {
			suara[masuk] = suara[masuk] + 1
			sah = sah + 1
		}
		fmt.Scan(&masuk)
	}

	fmt.Println("Suara masuk: ", total)
	fmt.Println("Suara sah: ", sah)
	for i := 0; i < sah; i++ {
		if suara[i] != 0 {
			fmt.Println(i, ": ", suara[i])
		}
	}
}