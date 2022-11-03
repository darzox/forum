package storage

import (
	"database/sql"
)

type User struct {
	Id       uint32
	Email    string
	Username string
	Passwrod string
}

func NewUser(email, username, password string) *User {
	return &User{
		Email:    email,
		Username: username,
		Passwrod: password,
	}
}

func createUser(db *sql.DB, user *User) error {
	records := `INSERT INTO user(email, username, password)
				VALUES (?,?,?)`
	query, err := db.Prepare(records)
	if err != nil {
		return err
	}
	_, err = query.Exec(user.Email, user.Username, user.Passwrod)
	if err != nil {
		return err
	}
	return nil
}
