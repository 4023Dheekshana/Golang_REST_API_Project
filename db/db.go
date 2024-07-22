package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "modernc.org/sqlite"
)

var DB *sql.DB

func InitDB() {

	db, err := sql.Open("sqlite", "api.sql")
	if err != nil {
		log.Fatal("Database could not connect: ", err)
	}

	DB = db

	createTables()

	fmt.Println("Tables created successfully!")
}

func createTables() {
	createUsersTable := `
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			email TEXT NOT NULL UNIQUE,
			password TEXT NOT NULL
		)
	`

	_, err := DB.Exec(createUsersTable)

	if err != nil {
		panic("Could not create events table")
	}
	createEventsTable := `
        CREATE TABLE IF NOT EXISTS events (
            id INTEGER PRIMARY KEY AUTOINCREMENT,
            name TEXT NOT NULL,
            description TEXT NOT NULL,
            location TEXT NOT NULL,
            datetime DATETIME NOT NULL,
            user_id INTEGER,
			FOREIGN KEY(user_id) REFERENCES users(id)
        )
    `

	_, err = DB.Exec(createEventsTable)
	if err != nil {
		panic("Could not create event table")
	}

	createRegistrationTable := `
	CREATE TABLE IF NOT EXISTS registrations (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	event_id INTEGER,
	user_id INTEGER,
	FOREIGN KEY(event_id) REFERENCES events(id),
	FOREIGN KEY(user_id) REFERENCES  users(id)
	)`

	_, err = DB.Exec(createRegistrationTable)
	if err != nil {
		panic("could not create registration table")
	}

}
