package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
)

var (
	db   *sql.DB
	mock bool
)

var normalQuery = func(name string) (age int, err error) {
	err = db.QueryRow("SELECT age FROM users WHERE name=?", name).Scan(&age)
	if err != nil {
		err = fmt.Errorf("no such user: %v; %w ", name, err)
	}
	return age, err
}

var mockQuery = func(name string) (age int, err error) {
	return 0, fmt.Errorf("no such user: %v; %w ", name, sql.ErrNoRows)
}

func QueryAgeByName(name string) (age int, err error) {
	if mock {
		return mockQuery(name)
	} else {
		return normalQuery(name)
	}
}

func main() {
	mock = true
	name := "lumoping"
	age, err := QueryAgeByName(name)
	if errors.Is(err, sql.ErrNoRows) {
		log.Printf("Cannot get user %v'age, cause:  %v", name, err)
	} else {
		log.Printf("user %v's age : %v", name, age)
	}
}
