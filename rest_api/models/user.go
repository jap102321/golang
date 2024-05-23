package models

import (
	"errors"

	"example.com/rest-api/db"
	"example.com/rest-api/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u User) SaveUser() error {
	query := "INSERT INTO users( email, password) values (?, ?)"

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	hashedPass, err := utils.HashPassword(u.Password)

	if err != nil {
		return err
	}

	res, err := stmt.Exec(u.Email, hashedPass)

	if err != nil {
		return err
	}

	uId, err := res.LastInsertId()

	u.ID = uId

	if err != nil {
		return err
	}

	return err
}

func (u *User) ValidateCredentials() error {
	query := "SELECT id, password FROM users WHERE email = ?"

	row := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string

	err := row.Scan(&u.ID, &retrievedPassword)

	if err != nil {
		return errors.New("Credentials is not valid")
	}

	passwordIsValid := utils.CheckPasswords(u.Password, retrievedPassword)

	if !passwordIsValid {
		return errors.New("Credentials is not valid")
	}

	return nil
}
