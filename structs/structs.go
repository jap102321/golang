package main

import (
	"fmt"

	"example.com/structs/user"
)

func main() {
	firstName := getUserData("Please enter your first name: ")
	lastName := getUserData("Please enter your last name: ")
	birthdate := getUserData("Please enter your birthdate (MM/DD/YYYY): ")

	// ... do something awesome with that gathered data!
	var appUser *user.User
	appUser, err := user.NewUser(firstName, lastName, birthdate)

	if err != nil {
		fmt.Println(err)
		return
	}

	appUser.OutputUserDetails()

	admin := user.NewAdmin("test@example.com", "1234")

	admin.UserData.OutputUserDetails()
	admin.OutputAdminData()

}

func getUserData(promptText string) string {
	fmt.Print(promptText)
	var value string
	fmt.Scanln(&value)
	return value
}
