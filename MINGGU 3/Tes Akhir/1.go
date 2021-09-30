package main

import "fmt"

func main() {
	var a, b, c, d, e byte
	var geser int
	var a2, b2, c2, d2, e2 int

	fmt.Scanf("%c%c%c%c%c\n", &a, &b, &c, &d, &e)
	fmt.Scanf("%d", &geser)

	a2 = int(a) + geser
	b2 = int(b) + geser
	c2 = int(c) + geser
	d2 = int(d) + geser
	e2 = int(e) + geser

	fmt.Printf("%c%c%c%c%c", a2, b2, c2, d2, e2)
}
