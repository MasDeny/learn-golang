package main

import "fmt"

func main() {
	birthdayList := []int{5,7,3,20,2} // <- contoh penggunaan slice yang lebih dinamis
	var name = [...]string{
		"Andi",
		"Habsi",
		"Tika",
		"Hendro",
		"Dedy",
		"Siska",
		"Tiwi",
	}


	//menghapus array
		// name[len(name)-1] = ""
		// name[len(name)-3] = ""

	// menambah array
	birthdayList = append(birthdayList, 13, 22)
	// name = append(name, )

	fmt.Println("\nBerikut adalah daftar anak yang lahir bulan Juli :\n")
	for i, valueName := range name {
		if valueName != "" {
			fmt.Println("Anak yang bernama", valueName, "berulang tahun pada tanggal", birthdayList[i])
		}else {
			fmt.Println("Tidak ada anak yang berulang tahun pada tanggal", birthdayList[i])
		}
	}

	fmt.Println("\nBerikut Contoh Penggunaan pada array multidimensi :\n")
	var angka = [5][2]int{{0,3}, {5,5}, {5,9}, {4,2}, {2,7}}
	for i := 0; i < 5; i++ {
		for j := 0; j < 2; j++ {
			fmt.Printf("Data ke [%d][%d] = %d\n", i, j, angka[i][j])
		}
		
	}

	// cara 2 penggunaan slice yaitu memotong beberapa data saja yang hanya ditampilkan
	fmt.Println("\nBerikut Contoh Penggunaan slice  :\n")
	OnePiece := []string{
		"Luffy",
		"Zorro",
		"Nami",
		"Sanji",
		"Ussop",
		"Chopper",
		"Robin",
		"Franky",
		"Brook",
		"Law",
	}
	// memotong data yang ditampilkan
	var slices = OnePiece[3:5]
	fmt.Println(slices)
	// penggunaan fungsi pada slice
	fmt.Println(len(slices))
	fmt.Println(cap(slices))
	nameCopies := make([]string, len(OnePiece))
	copy(nameCopies, OnePiece)
	fmt.Println(nameCopies)
	// contoh penggunaan append
	var slice2 = OnePiece[7:]
	fmt.Println(slice2)

	var slice3 = append(slice2, "Vivi")
	slice3[2] = "Jimbei"
	fmt.Println(slice3)
	fmt.Println(OnePiece)



}

/*
	Pada go-lang untuk penyebutan array harus menentukan dengan jumlah yang akan 
	digunakan, bisa sebuah array tidak harus menentukan jumlah namun harus dengan
	syarat memasukkan nilai awalnya
*/

/*
	Slice hampir mirip seperti array namun data pada slice lebih dinamis karena
	dapat ditambahkan dengan cara append, jika disimpukan array lebih cocok untuk
	data yang lebih statis dan slice lebih cocok digunakan untuk data yang dinamis
	Fungsi pada slice lebih beragam, seperti
	 - Append => untuk menambahkan data pada slice yang dibuat sebelumnya
	 - Len => untuk menghitung panjang slice yang dibuat
	 - copy => untuk menyalin sebuah slice dari yang sudah ada sebelumnya
	 - cap => untuk mendapatkan kapasitas slice
	 - make([]dataType, length, capacity)
*/