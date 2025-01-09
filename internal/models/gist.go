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
					WHERE expired > NOW() AND id=$1`

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
	return s, nil
}

func (g *GistModel) Latest() ([]*Gist, error) {
	sqlStatement := `SELECT id, title, content, created, expires FROM gist
					WHERE expired > NOW() ORDER BY DESC LIMIT 10`

	rows, err := g.DB.Query(sqlStatement)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	gists := []*Gist{}

	for rows.Next() {
		s := &Gist{}

		err = rows.Scan(&s.ID, &s.Title, &s.Content, &s.Created, &s.Expires)
		if err != nil {
			return nil, err
		}
		gists = append(gists, s)
	}

	if err = rows.Err(); err != nil {
		return nil, err
	}
	return gists, nil
}
