package main

import "fmt"

/*
	スライスは length と capacity を持っている
	length: スライスの要素数
	capacity: 元となる配列の要素数

	下の例だとcapの出力は常に「6」となる

*/

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// Zero length
	s = s[:0]
	printSlice(s)

	// Extends its length
	s = s[:4]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}
