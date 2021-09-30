package main

import "fmt"

func main() {
	// pendeklarasian variabel
	var min1, min2, min, max1, max2, max, a, b, c, d int

	// scan 4 bilangan
	fmt.Scanln(&a, &b, &c, &d)

	// mulai perbandingan
	minmax(a, b, &min1, &max1)
	minmax(c, d, &min2, &max2)
	minmax(min1, min2, &min, &max1)
	minmax(max1, max2, &min1, &max)

	// print max dan min
	fmt.Println(max, min)
}

func minmax(f1, f2 int, min, max *int) {
	if f1 < f2 {
		*min = f1 
		*max = f2
	} else {
		*max = f1
		*min = f2
	}
}