package main

import "fmt"

func main() {
	var a1, a2, kunci int

	// scan kunci
	fmt.Scan(&kunci)

	// iterate selama inputan bukan -1 
	for a1 != -1 && a2 != -1 {
		fmt.Scanln(&a1, &a2)
		if a1 % kunci == 0 {
			fmt.Println(a2)
		} else if a2 % kunci == 0 {
			fmt.Println(a1)
		}
	}
}