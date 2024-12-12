package models

import (
	"errors"

	"restapi.com/dagem/db"
	"restapi.com/dagem/utils"
)

type User struct {
	ID       int64
	Email    string `binding:"required,email"`
	Password string `binding:"required,min=8"`
}

func (u User) Save() error {
	query := "INSERT INTO users(email,password) VALUES(?,?)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	hashedPassword, err := utils.HashPasswordd(u.Password)
	if err != nil {
		return nil
	}
	result, err := stmt.Exec(u.Email, hashedPassword)
	if err != nil {
		return nil
	}
	id, err := result.LastInsertId()
	u.ID = id
	return err
}

func (u *User) ValidateCredentials() error {
	query := "SELECT id, password from users WHERE email=?"
	row := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string
	err := row.Scan(&u.ID, &retrievedPassword)
	if err != nil {
		return errors.New("credentials is invalid")
	}
	passwordIsValid := utils.CheckPassword(u.Password, retrievedPassword)

	if !passwordIsValid {
		return errors.New("credentials is invalid")
	}
	return nil
}
