// PROGRAM URUT
package main

import "fmt"

func main(){
	var bil [100]int
	var n, i1, i2 int

	isiArray(&bil, &n, &i1, &i2)
	sorting(&bil, i1, i2, n)
	tampilArray(bil, n)
}

func isiArray(t *[100]int, n, i1, i2 *int){
/* IS. terdefinisi bilangan bulat n, dan n buah bilangan bulat telah siap pada
piranti masukan
FS. array t berisi n buah bilangan bulat yang berasal dari user */
	fmt.Scanln(n, i1, i2)
	for i:=0; i<*n; i++ {
		fmt.Scan(&t[i])
	}
}

func tampilArray(t [100]int, n int){
/* IS. terdefinisi sebuah array t yang berisi n buah bilang bulat
FS. menampilkan array t ke layar secara horizontal */
	for i:=0; i<n; i++ {
		fmt.Print(t[i], " ")
	}
}

func sorting(t *[100]int, u,d,n int){
/* IS. terdefinisi sebuah array t yang berisi n bilangan bulat, dan indeks bilangan
bulat u dan d
FS. array t terurut descending dari indeks ke-u hingga ke-d dengan menggunakan
algoritma insertion sort */
	for i := u; i <= d; i++ {
        var j = i

		for j > u {
            if t[j-1] < t[j] {
                t[j-1], t[j] = t[j], t[j-1]
            }
            j--
        }
    }
}	