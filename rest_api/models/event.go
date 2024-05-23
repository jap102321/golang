package models

import (
	"time"

	"example.com/rest-api/db"
)

type Event struct {
	ID          int64
	Title       string    `binding:"required"`
	Description string    `binding:"required"`
	Location    string    `binding:"required"`
	DateTime    time.Time `binding:"required"`
	UserId      int64
}

var events = []Event{}

func (e *Event) Save() error {
	query := `INSERT INTO events(name, description, location, dateTime,user_id) 
	VALUES(
		?, ?, ?, ?, ?
	)
	`
	statement, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}

	defer statement.Close()

	res, err := statement.Exec(e.Title, e.Description, e.Location, e.DateTime, e.UserId)

	if err != nil {
		return err
	}

	id, err := res.LastInsertId()

	e.ID = id

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
		err := rows.Scan(&event.ID, &event.Title, &event.Description, &event.Location, &event.DateTime, &event.UserId)

		if err != nil {
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
	err := row.Scan(&event.ID, &event.Title, &event.Description, &event.Location, &event.DateTime, &event.UserId)

	if err != nil {
		return nil, err
	}

	return &event, nil
}

func (e Event) UpdateEvent() error {
	query := `UPDATE events 
	SET name = ?, description = ?, location = ?, dateTime = ?
	WHERE id = ?
	`

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.Title, e.Description, e.Location, e.DateTime, e.ID)

	return err
}

func (e Event) DeleteEvent() error {
	query := "DELETE FROM events WHERE id = ?"

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.ID)

	return err
}

func (e Event) Register(userId int64) error {
	query := "INSERT INTO registrations(event_id, user_id) values (?, ?)"

	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.ID, userId)

	return err
}

func (e Event) CancelRegistration(userId int64) error {
	query := "DELETE FROM registration WHERE event_id = ? AND user_id = ?"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return err
	}

	defer stmt.Close()

	_, err = stmt.Exec(e.ID, userId)

	return err
}
