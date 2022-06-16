package main

import "fmt"

func FaktorBilanganDesc(n int) string {
	// your code here
	var result string
	for i := 10; i >= 1; i-- {
		if n%i == 0 {
			result += fmt.Sprintf("%d\n", i)
		}
	}

	fmt.Printf("tipe: %T %s\n", n, result)
	return result
}

func main() {
	var number int
	fmt.Scanf("%d", &number)
	fmt.Println(FaktorBilanganDesc(number))
}
