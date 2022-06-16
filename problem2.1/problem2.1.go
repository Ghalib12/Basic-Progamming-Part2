package main

import "fmt"

func FaktorBilangan(n int) string {
	var result string
	for a := 1; a <= n; a++ {
		if n%a == 0 {
			result += fmt.Sprintf("%d\n", a)
		}
	}

	fmt.Printf("tipe: %T %s\n", n, result)
	return result
}

func main() {
	var number int
	fmt.Scanf("%d", &number)
	fmt.Println(FaktorBilangan(number))
}
