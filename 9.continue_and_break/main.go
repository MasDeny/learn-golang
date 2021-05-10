package main

import "fmt"

func main() {
	fmt.Println("Penggunaan Break: ")
	// contoh penggunaan break
	for i := 0; i < 10; i++ {
		if i==4 {
			break
		}
		fmt.Println("perulangan ke-", i)
	}

	fmt.Println("\nPenggunaan Continue: ")
	// contoh penggunaan continue
	for i := 10; i < 20; i++ {
		if i % 2 == 0 {
			continue
		}
		fmt.Println("perulangan ke-", i)
	}
}