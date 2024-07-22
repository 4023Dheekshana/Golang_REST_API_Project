package models

import (
	"errors"
	"log"

	"dheek.com/restapi/db"
	"dheek.com/restapi/utils"
)

type User struct {
	Id       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (u *User) Save() error {
	query := "INSERT INTO users(email, password)VALUES(?, ?)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		log.Fatalf("Error preparing query: %v", err)
		return err
	}
	defer stmt.Close()
	hashedPW, err := utils.HashPassword(u.Password)
	if err != nil {
		log.Fatalf("Error hashing password %v", err)
		return err
	}
	result, err := stmt.Exec(u.Email, hashedPW)
	if err != nil {
		log.Fatalf("Error executing statement %v", err)
		return err
	}

	userId, err := result.LastInsertId()
	u.Id = userId
	return err
}

func (u *User) ValidateCredentials() error {
	query := "SELECT id, password FROM users WHERE email = ?"
	row := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string

	err := row.Scan(&u.Id, &retrievedPassword)

	if err != nil {
		log.Fatalf("Error getting retrieved password %v", err)
	}
	//log.Printf("Retreived email: %v,Retrieved Password: %v", retrievedEmail, retrievedPassword)
	passwordIsValid := utils.CheckPasswordHash(u.Password, retrievedPassword)
	if !passwordIsValid {
		return errors.New("credentials invalid")
	}
	return nil
}
