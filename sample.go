// 問題: https://github.com/gopherdojo/dojo1/tree/master/kadai1

package main

import (
	"fmt"
	"log"
	"io/ioutil"
)


// @see https://news.mynavi.jp/article/gogogo-5/
func main() {
	files, err := ioutil.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		fmt.Println(file.Name())
	}
}