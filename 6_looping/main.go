package main

import (
	"6_looping/loop"
	"fmt"
)

func main() {
	var numb int
	fmt.Println("==== Perulangan ==== \n")
	fmt.Println("Silahkan Pilih ingin menggunakan perungan yang mana")
	fmt.Println("1. For Loop")
	fmt.Println("2. ForWhile Loop")
	fmt.Println("3. Range Loop\n")
	fmt.Scanf("%d", &numb)

	switch numb {
		case 1:
		loop.Forloop()
		case 2:
		loop.Whileloop()
		case 3:
		loop.Rangeloop()
	}
	
}