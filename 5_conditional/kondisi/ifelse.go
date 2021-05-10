package kondisi

import "fmt"

func Ifelse() {
	var nilai int
	fmt.Println("\n==== IF ELSE CONDITIONAL ====")
	fmt.Println("\nMasukkan nilai anda : ")
	fmt.Scanf("%d", &nilai)
	if nilai > 7 {
		fmt.Println("Selamat Anda Lulus dengan nilai", nilai)
	}else {
		fmt.Println("Maaf Anda Belum Lulus karena mendapat nilai", nilai, ". Tetap Semangat !!")
	}
}