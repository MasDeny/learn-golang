package main

import "fmt"

func main() {
	motorCycles := map[string]string{
		"Brand" 	: "Honda",
		"Engine" 	: "150cc",
		"Type"		: "Tracker",
		"Colour"	: "Red/Black/Gray",
	}

	fmt.Println(motorCycles)

	// menambahkan data pada map
	motorCycles["Name"] = "Honda CBR-150"
	// mengubah data pada map
	motorCycles["Type"] = "Touring"
	// menghapus data pada map
	delete(motorCycles, "Colour")

	fmt.Println(motorCycles)
	fmt.Println(motorCycles["Brand"])
	fmt.Println(motorCycles["Engine"])
}

/*
Penggunaan map hampir mirip dengan slice dan array namun yang membedakan map
dengan yang lain adalah pada map key yang digunakan untuk penyebutan value lebih
dinamis dan dapat diubah semau kita
*/