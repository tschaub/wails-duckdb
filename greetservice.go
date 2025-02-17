package main

import (
	"database/sql"
	"errors"
	"fmt"

	_ "github.com/marcboeker/go-duckdb"
)

type GreetService struct{}

func (g *GreetService) Greet(name string) (string, error) {
	db, err := sql.Open("duckdb", "")
	if err != nil {
		return "", err
	}
	defer db.Close()

	_, err = db.Exec(`CREATE TABLE people (id INTEGER, name VARCHAR)`)
	if err != nil {
		return "", err
	}
	_, err = db.Exec(`INSERT INTO people VALUES (42, ?)`, name)
	if err != nil {
		return "", err
	}

	row := db.QueryRow(`SELECT id, name FROM people`)

	var id int
	err = row.Scan(&id, &name)
	if errors.Is(err, sql.ErrNoRows) {
		return "no rows", nil
	} else if err != nil {
		return "", err
	}

	return fmt.Sprintf("Hello, %s (%d)", name, id), nil
}
