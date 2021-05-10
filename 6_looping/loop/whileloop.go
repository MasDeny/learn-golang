package loop

import "fmt"

func Whileloop() {
	benar:= true
	i := 0

	for benar {
		if i==7 {
			// benar = false
			// bisa juga menggunakan break
			break
		}
		fmt.Println("Perulangan Angka ke-", i)
		i++
	}

	// or 

	index := 0

	for index<6 {
		fmt.Println("Perulangan Metode ke Dua, Angka ke-", index)
		index++
	}
}