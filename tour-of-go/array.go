package main

import "fmt"

func main() {

	// Array
	var a [2]string
	a[0] = "Hello"
	a[1] = "world"
	fmt.Println(a)

	// Slices - 長さのない配列リテラル
	primes := [6]int{2, 3, 5, 7, 11, 13}
	var s []int = primes[1:4]
	fmt.Println(s)

	a2 := [3]bool{true, true, false} // 配列
	s2 := []bool{true, true, false}  // slice
}
