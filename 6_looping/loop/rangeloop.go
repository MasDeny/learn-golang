package loop

import "fmt"

func Rangeloop() {
	data := []string{"andi", "sulaeman"}
	data = append(data, "totok", "nugrahas")
	for index, v := range data {
		fmt.Println("Perulangan Index Ke-", index, v)
	}
}