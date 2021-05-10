package main

import (
	"5_conditional/kondisi"
	"fmt"
)

func main(){
	var tipe int
	fmt.Println("Pilih Jenis Conditional : ")
	fmt.Println("1. IfElse")
	fmt.Println("2. SwitchCase")
	fmt.Println("Masukkan angka untuk memilih jenis conditional ")
	fmt.Scanf("%d", &tipe)

	if tipe==1 {
		kondisi.Ifelse()
	} else {
		kondisi.Swtcase()
	}
}