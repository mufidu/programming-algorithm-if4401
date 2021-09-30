package main

import "fmt"

func main() {
	var angka int
	var a, b, c, d, e, f, g, h, i int

	fmt.Scanln(&angka)

	a = angka / 1000 
	b = angka % 1000 /100
	c = angka % 100 / 10
	d = angka % 10
	e = angka / 100
	f = angka % 1000 / 10
	g = angka % 100
	h = angka / 10
	i = angka % 1000

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)
	fmt.Println(e)
	fmt.Println(f)
	fmt.Println(g)
	fmt.Println(h)
	fmt.Println(i)
	fmt.Println(angka)

}