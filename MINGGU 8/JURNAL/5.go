package main

import "fmt"

func main() {
	var a, b, c, d, e, f, g, h, soal, skor, soaltemp, skortemp int
	var p, nama string

	nama = "X"	
	p = "A"
	skor = -1
	soal = -1

	for p != "Selesai" {
		fmt.Scanln(&p, &a, &b, &c, &d, &e, &f, &g, &h)
		skortemp = a + b + c + d + e + f + g + h
		
		if h == 0 && g == 0 && f == 0 && e == 0 && d == 0 && c == 0 && b == 0 && a == 0 {
			soaltemp = 0	
		} else if g == 0 && f == 0 && e == 0 && d == 0 && c == 0 && b == 0 && h == 0 {
			soaltemp = 1
		} else if f == 0 && e == 0 && d == 0 && c == 0 && g == 0 && h == 0 {
			soaltemp = 2
		} else if e == 0 && d == 0 && f == 0 && g == 0 && h == 0 {
			soaltemp = 3
		} else if e == 0 && f == 0 && g == 0 && h == 0 {
			soaltemp = 4
		} else if f == 0 && g == 0 && h == 0 {
			soaltemp = 5
		} else if g == 0 && h == 0 {
			soaltemp = 6
		} else if h == 0 {
			soaltemp = 7
		}

		if soaltemp > soal {
			soal = soaltemp
			skor = skortemp
			nama = p
		} else if soal == soaltemp {
			if skortemp > skor {
			soal = soaltemp
			skor = skortemp
			nama = p				
			}
		}
	}
}
