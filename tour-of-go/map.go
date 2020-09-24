package main

import "fmt"

func main() {
	m := make(map[string]int)

	m["answer"] = 42
	fmt.Println("The value:", m["answer"])

	delete(m, "answer") // deleteするとゼロ値になる
	fmt.Println("The value:", m["answer"])
}
