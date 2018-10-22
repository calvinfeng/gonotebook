package main

import (
	"database/sql"
	"fmt"
	"os"

	randomdata "github.com/Pallinder/go-randomdata"
	"github.com/gchaincl/dotsql"
	_ "github.com/lib/pq"
)

const (
	user     = "cfeng"
	password = "cfeng"
	dbname   = "rawsql"
)

var dbURL = fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", user, password, dbname)

func runMigration(db *sql.DB) error {
	dot, err := dotsql.LoadFromFile("./migrations/dog.sql")
	if err != nil {
		return err
	}

	_, err = dot.Exec(db, "create_dogs_table")
	if err != nil {
		return err
	}

	fmt.Println("Migration is successful")

	return nil
}

func createDogs(db *sql.DB) error {
	dot, err := dotsql.LoadFromFile("./queries/dog.sql")
	if err != nil {
		return err
	}

	for i := 0; i < 10; i++ {
		_, err := dot.Exec(db, "insert_dog", 5, randomdata.SillyName(), true)
		if err != nil {
			return err
		}
	}

	fmt.Println("Inserted 10 dogs!")

	return nil
}

func main() {
	fmt.Println("Connecting to", dbURL)
	db, err := sql.Open("postgres", dbURL)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	err = runMigration(db)
	if err != nil {
		fmt.Println(err)
	}

	err = createDogs(db)
	if err != nil {
		fmt.Println(err)
	}
}
