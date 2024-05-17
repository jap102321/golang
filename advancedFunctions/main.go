package main

import (
	"fmt"

	"example.com/advancedfunctions/recursion"
)

func main() {

	double := transformerFn(4)

	fmt.Println(double(5))

	fact := recursion.Recursion(5)

	fmt.Println(fact)
}
func transformerFn(factor int) func(int) int {
	return func(number int) int {
		return number * factor
	}
}
