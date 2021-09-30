// PROGRAM pilkart
package main

import "fmt"

type kotak [500]int

func main() {
	var masuk, i, total, sah, ketua, wakil, max,max2, jum int
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

	ketua = 0
	wakil = 0
	max = 0
	max2 = 0
	jum = 0

	fmt.Println("Suara masuk: ", total)
	fmt.Println("Suara sah: ", sah)

	for i := 0; i < sah; i++ {
		if suara[i] > max {
			max2 = max
			max = suara[i]
		}
	}

	for i := 0; i < sah; i++ {
		if suara[i] == max {
			jum++
		}
	}

	

	if jum == 1 {
		for i := 0; i < sah; i++ {
			if suara[i] == max {
				ketua = i
			}
		}
		for i := 0; i < sah; i++ {
			if suara[i] == max2 {
				wakil = i
			}
		}
	} 

	fmt.Println("Ketua RT: ", ketua)
	fmt.Println("Wakil ketua: ", wakil)
}