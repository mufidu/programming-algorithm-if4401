package main

import "fmt"

func main() {
	// pendeklarasian variabel
	var x int
	var a, b, c string

	// scan semua var
	fmt.Scanln(&x, &a, &b, &c)

	// memanggil function sesuai kondisi
	if a == "f" && b == "g" && c == "h" {
		fmt.Println(h(g(f(x))))
	} else if a == "g" && b == "f" && c == "h" {
		fmt.Println(h(f(g(x))))
	} else if a == "g" && b == "h" && c == "f" {
		fmt.Println(f(h(g(x))))
	} else if a == "f" && b == "h" && c == "g" {
		fmt.Println(g(h(f(x))))
	} else if a == "h" && b == "f" && c == "g" {
		fmt.Println(g(f(h(x))))
	} else if a == "h" && b == "g" && c == "f" {
		fmt.Println(f(g(h(x))))
	}
}

// Pendeklarasian function yang mereturn sesuai rumus 
func f(x int) int {
	return 2 * x + 5
}

func g(x int) int {
	return x * x + 2 * x
}

func h(x int) int {
	return x- 3
}