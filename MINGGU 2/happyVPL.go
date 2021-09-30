package main

import "fmt"

func main() {
	var awal, akhir int

	fmt.Scanln(&awal, &akhir)

	for i := awal; i <= akhir; i++ {
		fmt.Printf("Simbol ASCII dari %d adalah %c\n", i, i)
	}
}

