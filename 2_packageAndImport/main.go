package main

import (
	"2_packageAndImport/calc"
	"2_packageAndImport/quiz"
	"fmt"
)

func main() {
	// 3. belajar penggunaan package lebih lanjut
	result := calc.Add(6,6)
	fmt.Println(result)

	// 4. penggunaan import
	// 		- penggunaan standart library
	// 		- mengakses package berbeda
	// 		- yang dibikin orang lain dan di publish di github

	// kuis 
	resultMultiply := quiz.Multiply(5,9)
	fmt.Println(resultMultiply)
} 