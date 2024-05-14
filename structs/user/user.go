package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	FirstName string
	LastName  string
	BirthDate string
	createdAt time.Time
}

type Admin struct {
	Email    string
	Password string
	UserData User
}

func NewUser(firstName, lastName, birthDate string) (*User, error) {
	if firstName == "" || lastName == "" || birthDate == "" {
		return nil, errors.New("can't be empty")
	}
	return &User{
		FirstName: firstName,
		LastName:  lastName,
		BirthDate: birthDate,
		createdAt: time.Now(),
	}, nil
}

func NewAdmin(email, password string) *Admin {
	return &Admin{
		Email:    email,
		Password: password,
		UserData: User{
			FirstName: "ADMIN",
			LastName:  "ADMIN",
			BirthDate: "----",
			createdAt: time.Now(),
		},
	}

}

func (u User) OutputUserDetails() {
	fmt.Println(u.FirstName, u.LastName, u.BirthDate)
}

func (a *Admin) OutputAdminData() {
	fmt.Println(a.Email, a.Password)
}

func (u *User) ClearUserName() {
	u.FirstName = ""
	u.LastName = ""

}
