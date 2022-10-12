package main

import (
	"fmt"

	"github.com/cahyoariawan21/go-math-module"
)

func main() {
	a := 10.0
	b := 20.0

	Addition := math.Addition(a, b)
	Subtraction := math.Subtraction(a, b)
	Distribution := math.Distribution(a, b)
	Multiplication := math.Multiplication(a, b)

	fmt.Println(Addition * 2)
	fmt.Println(Subtraction)
	fmt.Println(Distribution)
	fmt.Println(Multiplication)

}
