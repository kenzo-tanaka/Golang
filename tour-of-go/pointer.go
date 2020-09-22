package main

import "fmt"

func main() {
	i := 43
	p := &i
	fmt.Println(p)  // メモリのアドレス
	fmt.Println(*p) // メモリの値(43)

}
