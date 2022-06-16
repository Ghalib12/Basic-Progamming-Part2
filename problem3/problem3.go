package main

import (
	"fmt"
)

func PrimeNumber(number int) bool {
	// your code here
	for bil := 1; bil < 10; bil++ {
		i := 0
		for bag := 1; bag < 10; bag++ {
			if bil%bag == 0 {
				i++
			}
		}
		if (i == 2) && (bil != 1) {
			fmt.Println(bil)
		}
	}
}

func main() {
	fmt.Println(PrimeNumber(3))
	fmt.Println(PrimeNumber(7))
	fmt.Println(PrimeNumber(10))
}
