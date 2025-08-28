package userblog

import "weblog/db"

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

	result, err := stm.Exec(u.Name, u.UserName, u.Email, u.Password)

	if err != nil {
		panic(err)
	}

	id, err := result.LastInsertId()
	u.Id = int(id)

	return err
}
