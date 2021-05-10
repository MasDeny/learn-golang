package calc

// - nama function harus Capitalize agar bisa di akses public oleh function lain
func Add(number int, numberTwo int ) int {
	return add(number, numberTwo)
}

// - nama function lowercase hanya dapat diakses oleh function sama
func add(number int, numberTwo int) int {
	return number + numberTwo
}