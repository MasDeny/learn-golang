package main

import "fmt"

func main() {
	// return multiple value
	// untuk ignore value yang di tampilkan bisa menggunakan _ (underscore)
	firstname, _, lastname := multipleValue()
	fmt.Println(firstname, lastname)
}

// return multiple values
func multipleValue() (string, int, string) {
	return "Deny", 3445, "Prayantoro"
}