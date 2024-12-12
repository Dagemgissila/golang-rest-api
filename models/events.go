package models

import (
	"restapi.com/dagem/db"
)

type Event struct {
	Id          int64
	Name        string `binding:"required"`
	Description string `binding:"required"`
	Location    string `binding:"required"`
	DateTime    string
	UserId      int64
}

var events = []Event{}

func (e *Event) Save() error {
	query := `INSERT INTO events(name,description,location,dateTime) VALUES(?,?,?,?)`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	results, err := stmt.Exec(e.Name, e.Description, e.Location, e.DateTime)
	if err != nil {
		return err
	}
	id, err := results.LastInsertId()
	if err != nil {
		return err
	}
	e.Id = id
	return err
}

func GetAllEvents() ([]Event, error) {
	query := "SELECT * FROM events"
	rows, err := db.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var events []Event

	for rows.Next() {
		var event Event
		err := rows.Scan(&event.Id, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)
		if err != nil {
			return nil, err
		}
		events = append(events, event)
	}
	return events, nil
}

func GetEventById(id int64) (*Event, error) {
	query := "SELECT * FROM events WHERE id=?"
	row := db.DB.QueryRow(query, id)
	var event Event
	err := row.Scan(&event.Id, &event.Name, &event.Description, &event.Location, &event.DateTime, &event.UserId)
	if err != nil {
		return nil, err
	}
	return &event, nil
}

func (event Event) Update() error {
	query := `
	UPDATE events
	SET name=?,description=?,location=?,dateTime=?
	WHERE id =?
	`
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(event.Name, event.Description, event.Location, event.DateTime, event.Id)
	return err
}

func (event Event) Delete() error {
	query := "DELETE FROM events WHERE id=?"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(event.Id)
	return err
}