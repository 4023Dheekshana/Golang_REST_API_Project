package models

import (
	"log"
	"time"

	"dheek.com/restapi/db"
)

type Event struct {
	ID          int64     `json:"id"`
	Name        string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserID      int64     `json:"user_id"`
}

func (e *Event) Save() error {
	query := `
	INSERT INTO events(name, description, location, datetime, user_id)
	VALUES(?,?,?,?,?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		log.Printf("Error preparing query: %v", err)
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime, e.UserID)
	if err != nil {
		log.Printf("Error executing query: %v", err)
		return err
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Printf("Error getting last insert ID: %v", err)
		return err
	}
	e.ID = id
	return nil
}

func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM events"
	rows, err := db.DB.Query(query)
	if err != nil {
		log.Printf("Error querying events: %v", err)
		return nil, err
	}
	defer rows.Close()

	var events = []Event{}
	for rows.Next() {
		var event Event
		err := rows.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
		if err != nil {
			log.Printf("Error scanning event: %v", err)
			return nil, err
		}

		events = append(events, event)
	}
	return events, nil
}

func GetEventById(id int64) (*Event, error) {
	query := "SELECT * FROM events WHERE id = ?"
	row := db.DB.QueryRow(query, id)

	var event Event
	err := row.Scan(&event.ID, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserID)
	if err != nil {
		log.Printf("Error scanning a row:%v ", err)
		return nil, err
	}
	return &event, err
}

func (event *Event) Update() error {
	query := `
	UPDATE events
	SET name = ?, description = ?, location = ?, datetime = ?
	WHERE id = ?`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		log.Printf("Error preparing query: %v", err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.ID)
	if err != nil {
		log.Printf("Error executing the statement %v", err)
		return err
	}
	return nil
}

func (event *Event) Delete() error {
	query := `
	DELETE FROM events WHERE id = ?`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		log.Fatalf("Error preparing query: %v", err)
		return err
	}
	_, err = stmt.Exec(event.ID)
	return err
}

func (event *Event) Register(userId int64) error {
	query := "INSERT INTO registrations(event_id, user_id) VALUES(?, ?)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		log.Fatalf("Error preparing statement %v", err)
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(event.ID, userId)
	return err
}

func (event *Event) CancelRegister(userId int64) error {
	query := "DELETE FROM REGISTRATIONS WHERE event_id = ? AND user_id = ?"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(event.ID, userId)

	return err
}
