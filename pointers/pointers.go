package main

func main() {

	age := 32

	getAdultYears(&age)
}

func getAdultYears(age *int) {
	*age = *age - 18
}
