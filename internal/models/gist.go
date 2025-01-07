package models

import (
	"database/sql"
	"errors"
	"time"
)

type Gist struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}

type GistModel struct {
	DB *sql.DB
}

func (g *GistModel) Insert(title, content string, expires int) error {
	sqlStatement := `INSERT INTO gist (title, content, created, expires) 
                     VALUES ($1, $2, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP + INTERVAL $3 DAY)`

	_, err := g.DB.Exec(sqlStatement, title, content, expires)
	if err != nil {
		return err 
	}

	return nil 
}

func (g *GistModel) Get(id int) (*Gist, error) {
	sqlStatement := `SELECT id, title, content, created, expires FROM gist 
					WHERE expired > UTC_TIMESTAMP() AND id=$1`

	row := g.DB.QueryRow(sqlStatement, id)

	s := &Gist{}

	err := row.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoRecords //better encapsulation
		} else {
			return nil, err 
		}
	}
}

func (g *GistModel) Latest() ([]*GistModel, error) {
	return nil, nil
}
