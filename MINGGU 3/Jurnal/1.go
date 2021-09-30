package main

import "fmt"

func main() {
	var a1, a2, a3, a4, a5 int
	var word string
	fmt.Scanln(&a1, &a2, &a3, &a4, &a5)
	fmt.Scanln(&word)
	fmt.Printf("%c", a1)
	fmt.Printf("%c", a2)
	fmt.Printf("%c", a3)
	fmt.Printf("%c", a4)
	fmt.Printf("%c", a5)
	fmt.Println()
	fmt.Printf("%c", word[0]+1)
	fmt.Printf("%c", word[1]+1)
	fmt.Printf("%c", word[2]+1)
}