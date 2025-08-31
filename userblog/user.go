package userblog

import (
	"errors"
	"weblog/authentication"
	"weblog/db"
)

type User struct {
	Id       int    `json:"id"`
	Name     string `json:"name"`
	UserName string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (u *User) RegisterUser() error {
	query := `INSERT INTO users(name,username,email,password) VALUES(?,?,?,?)`
	stm, err := db.DB.Prepare(query)

	if err != nil {
		panic(err)
	}
	defer stm.Close()

	hashedpassword, err := authentication.HashPassword(u.Password)
	if err != nil {
		return err
	}

	result, err := stm.Exec(u.Name, u.UserName, u.Email, hashedpassword)

	if err != nil {
		panic(err)
	}

	id, err := result.LastInsertId()
	u.Id = int(id)

	return err
}

func GetUserById(id int64) (*User, error) {
	query := "SELECT * FROM users WHERE id = ?"
	row := db.DB.QueryRow(query, id)
	var user User
	err := row.Scan(&user.Id, &user.Name, &user.UserName, &user.Email, &user.Password)

	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (u *User) ValidateCredentials() error {
	query := "SELECT id, password FROM users WHERE email = ?"

	row := db.DB.QueryRow(query, u.Email)

	var retrievedPassword string

	err := row.Scan(&u.Id, &retrievedPassword)

	if err != nil {
		return err
	}
	passwordIsValid := authentication.CheckPasswordHash(u.Password, retrievedPassword)

	if !passwordIsValid {
		return errors.New("credentials invalid")
	}

	return nil
}

func (u *User) Delete() error {
	query := `DELETE FROM users WHERE id = ? `

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(u.Id)

	return err
}
