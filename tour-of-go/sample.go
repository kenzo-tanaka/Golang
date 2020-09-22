package main

import (
	"fmt"
	"math"
)

var (
	TobBe bool = false
	x     int  = 4
)

// 定数は := で定義できない
const Pi = 3.1415

func add(x int, y int) int {
	return x + y
}

func split(sum int) (x, y int) {
	x = sum * 5
	y = sum - x
	return // named return value
}

func main() {
	fmt.Println(add(42, 13))
	fmt.Println(split(17))

	// Output value and type
	fmt.Printf("Value: %v, Type: %T\n", TobBe, TobBe)
	fmt.Printf("Value: %v, Type: %T\n", x, x)

	fmt.Println("Happy", Pi, "Day")
}

func forSample() {
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Printf(sum)

	// Goではwhileを使わずforだけを使う
	sum2 := 1
	for sum2 < 1000 {
		sum2 += sum
	}
	fmt.Println(sum2)
}

func pow(x, n, lim float64) float64 {

    // if文は条件の前に評価するための簡単なステートメントをかける
	if v := math.Pow(x, n); v < lim {
		return v
	}
	return lim
}
