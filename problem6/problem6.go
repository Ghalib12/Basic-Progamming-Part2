package main

import (
	"fmt"
)

func FullPrima(n int) bool {
	// write your codej
	angka := 23
	angkastr := fmt.Sprintf("%d", angka)
	for _, v := range angkastr {
		fmt.Println(string(v))
	}
	var arr = [5]int{1, 2, 3, 4, 5}
	fmt.Println(arr)
	var salary_a = map[string]int{"umam": 1000, "iswanul": 2000}
	fmt.Println(salary_a)

	var a = 10
	fmt.Println(&a)
	return
}

func main() {
	fmt.Println(FullPrima(2))  // true
	fmt.Println(FullPrima(3))  // true
	fmt.Println(FullPrima(11)) // false
	fmt.Println(FullPrima(13)) // false
	fmt.Println(FullPrima(23)) // true
	fmt.Println(FullPrima(29)) // false
	fmt.Println(FullPrima(37)) // true
	fmt.Println(FullPrima(41)) // false
	fmt.Println(FullPrima(43)) // false
	fmt.Println(FullPrima(53)) // true
}
