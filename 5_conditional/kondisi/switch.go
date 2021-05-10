package kondisi

import "fmt"

func Swtcase() {
	var makanApa string
	fmt.Println("\n==== Switch Case Conditional ====\n")
	fmt.Println("Jenis Pemakan Hewan ada 3 (Karnivora|Herbivora|Omnivora)")
	fmt.Println("Ingin mengetahui jenis apa ?")
	fmt.Scanf("%s", &makanApa)

	switch makanApa {
		case "Karnivora":
		fmt.Println("Karnivora ini adalah jenis hewan pemakan daging.")
		fmt.Println("Contoh Hewan Karnivora adalah : Singa, Harimau, Buaya, dsb")

		case "Herbivora":
		fmt.Println("Herbivora ini adalah jenis hewan pemakan tumbuhan.")
		fmt.Println("Contoh Hewan Karnivora adalah : Sapi, Kijang, Kambing, dsb")

		case "Omnivora":
		fmt.Println("Karnivora ini adalah jenis hewan pemakan segala.")
		fmt.Println("Contoh Hewan Karnivora adalah : OrangUtan, Tikus, Ayam, dsb")

		default: 
		fmt.Println("Belum Terdaftar")
		Swtcase()
	}
}