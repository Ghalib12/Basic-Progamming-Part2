package main

import (
	"fmt"
)

func PrimeNumber(number int) bool {
	// your code here
	var prima int8
	var bil bool
	for i := 1; i <= number; i++ {
		if number%i == 0 {
			prima++
		}
	}
	if prima == 2 {
		bil = true
	} else {
		bil = false
	}
	return bil
}

func main() {
	fmt.Println(PrimeNumber(11))
	fmt.Println(PrimeNumber(13))
	fmt.Println(PrimeNumber(17))
	fmt.Println(PrimeNumber(20))
	fmt.Println(PrimeNumber(35))
}
