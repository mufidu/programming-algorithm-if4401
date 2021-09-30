// PROGRAM PILKADA

package main

import "fmt"

func main() {
	var n1, n2, n3, pil int

	// diset ke 0 dulu
	n1 = 0
	n2 = 0
	n3 = 0
	pil = 0

	for pil != -1 {
			fmt.Scan(&pil)
			// masukkan hasil vote ke kotak peserta
			if pil == 1 {
				n1 = n1 + 1
			} else if pil == 2 {
				n2 = n2 + 1
			} else if pil == 3 {
				n3 = n3 + 1
			}
		}

		// print hasil sesuai soal
		if n1 != n3 && n1 != n2 && n2 != n3 {
		fmt.Println("Calon no. 1 dapat", n1, "suara, no.2 dapat", n2, "suara, dan no. 3 dapat", n3, "suara.")
		} else if n1 == n2 && n2 == n2 && n1 == n3 {
			fmt.Println("Calon no. 1, 2, dan 3 dapat", n1, "suara.")
		} else if n1 == n2 {
			fmt.Println("Calon no. 1 dan 2 dapat",n1, "suara, sementara no.3 dapat", n3, "suara.")
		} else if n1 == n3 {
			fmt.Println("Calon no. 1 dan 3 dapat",n1, "suara, sementara no.2 dapat", n2, "suara.")
		} else if n2 == n3 {
			fmt.Println("Calon no. 2 dan 3 dapat",n2, "suara, sementara no.1 dapat", n1, "suara.")
		}
}


