package main

import "fmt"

func Pangkat(base, pangkat int) int {
	// your code here
	var x int = base
	var n int = pangkat
	var res int = 1
	for i := 0; i < n; i++ {
		res *= x
	}

	return res
}

func main() {
	fmt.Println(Pangkat(2, 3))  // 8
	fmt.Println(Pangkat(5, 3))  // 125
	fmt.Println(Pangkat(10, 2)) // 100
	fmt.Println(Pangkat(2, 5))  // 32
	fmt.Println(Pangkat(7, 3))  // 343
}
