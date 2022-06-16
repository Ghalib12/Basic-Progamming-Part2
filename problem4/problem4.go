package main

import "fmt"

func Palindrome(input string) bool {
	// your code here
	panjang_input := len(input)
	for i := 0; i < panjang_input/2; i++ {
		if input[i] != input[panjang_input-i-1] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(Palindrome("civic"))       // true
	fmt.Println(Palindrome("katak"))       // true
	fmt.Println(Palindrome("kasur rusak")) // true
	fmt.Println(Palindrome("kupu-kupu"))   // false
	fmt.Println(Palindrome("lion"))        // false
}
