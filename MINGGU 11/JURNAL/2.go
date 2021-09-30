package main

import "fmt"

type tag struct {
	topik  string
	banyak int
}

type tabTopic [100]string

type tabTag [100]tag

func main() {
	var tags tabTag

	isiTabTag(&tags)
	fmt.Println(tags[cariTrending(tags)].topik)
}

// mengecek keberaddaan topik di array
func cekTopik(topik string, tags tabTag) bool {
	var i int
	i = 0

	for tags[i].topik != "" {
		if tags[i].topik == topik {
			return true
		} else {
			i++
		}
	}

	return false
}

// mTtabTagisi array tabtag
func isiTabTag(tags *tabTag) {
	var topik string
	var i int
	i = 0

	fmt.Scan(&topik)

	for topik != "." {
		if !(cekTopik(topik, *tags)) {
			*&tags[i].topik = topik
			*&tags[i].banyak = 1
		} else {
			*&tags[i].banyak++
		}
		i++
		fmt.Scan(&topik)
	}
}

// mencari topik yang muncul terbanyak
func cariTrending(tags tabTag) int {
	var temp int
	temp = 0
	for i:= 0; i < 100; i++ {
		if tags[i].banyak > temp {
			temp = i
		}
	}

	return temp
}