// PROGRAM PERTEMUAN FAKTOR

package main

import "fmt"

func main() {
	var n, p, q, hari int
	fmt.Scanln(&n, &p, &q)
	hitung(n,p,q,hari)
}

func hitung(n, p, q, hari  int) {
	var ntemp int
	if (n % p == 0) && (n % q == 0) {
		ntemp = n + 7
		hitung(ntemp,p,q,hari)
	} else if (n % p == 0) || (n % q == 0) {
		fmt.Println("Kita bertemu hari ini")
	} else if (n % p) < (n % q) {
		hari = (n % q) - (n % p)
		fmt.Println("Kita akan bertemu", hari, "hari lagi")
	} else if (n % p) > (n % q) {
		hari = (n % p) - (n % q)
		fmt.Println("Kita akan bertemu", hari, "hari lagi")
	} else if p == q {
		fmt.Println("Kita tidak pernah bertemu 🙏")
	}
}



