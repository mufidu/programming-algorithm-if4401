// PROGRAM KPK

package main

import "fmt" 

func main() {
	var m, n, kpk, temp int

	fmt.Scanln(&m, &n)
	temp = 1

	// set kpk sebagai pengali
	if m > n {
		kpk = m
	} else if m < n {
		kpk = m
	}

	// iterate kpk sampai ketemu
	for (kpk % n) != 0 {
		kpk = m * temp
		temp = temp + 1
	}

	fmt.Println(kpk)

}