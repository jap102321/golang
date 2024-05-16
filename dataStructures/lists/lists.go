package lists

import "fmt"

type Product struct {
	title string
	id    string
	price float64
}

func main() {
	prices := []float64{25, 30, 35, 45, 30, 90, 110, 240, 340, 50}

	fmt.Println(prices[1:])
	fmt.Println(cap(prices))

}
