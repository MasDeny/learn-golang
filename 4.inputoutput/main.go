package main

import "fmt"

func main() {
	var nama, matkul string
	var nilai, nim int

	// mengambil input dari terminal
	fmt.Println("Masukkan NIM anda : ")
	fmt.Scanf("%d", &nim)
	fmt.Println("Masukkan Nama anda : ")
	fmt.Scanf("%s", &nama)
	fmt.Println("Masukkan Mata Kuliah anda : ")
	fmt.Scanf("%s", &matkul)
	fmt.Println("Masukkan Nilai anda : ")
	fmt.Scanln(&nilai)

	//mengembalikan nilai dari input
	fmt.Printf("Nama(NIM) => %s (%d) \n", nama, nim)
	fmt.Printf("Matkul(Nilai) => %s (%d) \n", matkul, nilai)

}

/*
	%v, untuk semua tipe data dengan nilai kembalian string.
	%t, untuk tipe data boolean, TRUE atau FALSE.
	%b, untuk menampilkan data biner yaitu 0 dan 1. Misalnya input nilai integer maka akan di cetak nilai biner dari integer tersebut.
	%c, ,untuk menampilkan data unicode.
	%d, untuk menampilkan data integer dengan panjang maksimal 10.
	%f, untuk menampilkan data desimal dengan input tipe data float.
	%s, untuk menampilkan data string.
	%t, untuk menampilkan jenis variable yang digunakan.
	%%, untuk menampilkan karakter %.
	dan lainnya.
*/