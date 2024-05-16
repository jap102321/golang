package practice

import "fmt"

type Product struct {
	title string
	id    string
	price float64
}

func practice() {
	//1
	hobbiesArray := [3]string{"videogames", "gym", "code"}

	//2

	newHobbieArray := hobbiesArray[1:3]
	fmt.Println(hobbiesArray[0])
	fmt.Println(newHobbieArray)

	//3
	sliceOfHobbies := hobbiesArray[:2]
	// sliceOfHobbiesV2 := hobbiesArray[0:2]
	fmt.Println(sliceOfHobbies)

	//4
	sliceOfHobbies = hobbiesArray[1:]
	fmt.Println(sliceOfHobbies)

	//5
	courseGoals := []string{"Learn Go in Depth,", "Find new opportunities in my current job with Go"}
	courseGoals[1] = "Land a Job in Japan with Golang,"
	courseGoals = append(courseGoals, "Learn backend with go,")
	fmt.Println(courseGoals)

	//Bonus
	sliceOfProducts := []Product{
		{title: "Cake", price: 12.99, id: "cakeId"},
		{title: "Coffee", price: 15.99, id: "coffeId"},
	}

	fmt.Println(sliceOfProducts)

	newProduct := Product{
		"Tea", "teaId", 12.55,
	}

	sliceOfProducts = append(sliceOfProducts, newProduct)

	fmt.Println(sliceOfProducts)

}

// Time to practice what you learned!

// 1) Create a new array (!) that contains three hobbies you have
// 		Output (print) that array in the command line.
// 2) Also output more data about that array:
//		- The first element (standalone)
//		- The second and third element combined as a new list
// 3) Create a slice based on the first element that contains
//		the first and second elements.
//		Create that slice in two different ways (i.e. create two slices in the end)
// 4) Re-slice the slice from (3) and change it to contain the second
//		and last element of the original array.
// 5) Create a "dynamic array" that contains your course goals (at least 2 goals)
// 6) Set the second goal to a different one AND then add a third goal to that existing dynamic array
// 7) Bonus: Create a "Product" struct with title, id, price and create a
//		dynamic list of products (at least 2 products).
//		Then add a third product to the existing list of products.
