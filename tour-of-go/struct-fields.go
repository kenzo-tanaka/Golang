package main

import "fmt"

// struct名とフィールド名はそれぞれ大文字
type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}
	fmt.Println(v.X)

	v.X = 4
	fmt.Println(v.X)
}
