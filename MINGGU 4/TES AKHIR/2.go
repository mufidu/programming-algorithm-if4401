package main

import "fmt"

func main() {
	var kunci, a1, a2 int

	// scan kunci
	fmt.Scanln(&kunci)

	// iterate sebanyak kunci, print sesuai kunci
	for i := 0; i < kunci; i++ {
		fmt.Scanln(&a1, &a2)
		if kunci % 2 == 0 {
			fmt.Println(a1)
		} else {
			fmt.Println(a2)
		}
	}
}