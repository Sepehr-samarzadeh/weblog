package post

import (
	"time"
	"weblog/db"
)

type BlogPost struct {
	Id    int       `json:"id"`
	Title string    `json:"title"`
	Text  string    `json:"text"`
	Time  time.Time `json:"time"`
}

func (p *BlogPost) Save() error {
	query := `INSERT INTO posts (title,text,time) VALUES(?,?,?)`

	stm, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stm.Close()

	result, err := stm.Exec(p.Title, p.Text, p.Time)

	if err != nil {
		return err
	}

	id, err := result.LastInsertId()

	p.Id = int(id)

	return err
}
