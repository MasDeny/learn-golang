package main

import "fmt"

func main() {
	// return multiple value
	// untuk ignore value yang di tampilkan bisa menggunakan _ (underscore)
	firstname, _, lastname := multipleValue()
	fmt.Println(firstname, lastname)

	//return named values
	yeah, one, piece := namedValues()
	fmt.Println(yeah, "|", one, "|", piece)

	//variadic function
	sum := variadicFunction // <== contoh penggunaan fuction as value
	data := sum(100,70,50,50,30) // cara penyebutan function as value
	fmt.Println("total uang yang dibagi untuk ampau", data)
	//variadic function dapat digunakan pada slice yang sudah ada sebelumnya
	slice := []int{50,60,70,80,90}
	fmt.Println("total yang lain adalah", sum(slice...)) // cara penyebutan function as value dengan contoh lain
}

// return multiple values
func multipleValue() (string, int, string) {
	return "Deny", 3445, "Prayantoro"
}

// return named values
func namedValues() (name, skills, ability string) {
	name = "Monkey D Luffy"
	skills = "Gomu No Gomu"
	ability = "Gear Forth"
	return
}

// variadic function
// variable argument yang dapat diinputkan slice / lebih dari satu value
func variadicFunction(numbers ...int) int {
	total:=0
	for _, value := range numbers {
		total+=value
	}
	
	return total
}