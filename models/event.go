package models

import "time"

const EVENT_TABLE = "events"

type Event struct {
	Id           int       `sql:"id"`
	name         string    `sql:"name"`
	description  string    `sql:"description"`
	proposedTime time.Time `sql:"proposed_time"`
	actualTime   time.Time `sql:"actual_time"`
	createdAt    time.Time `sql:"created_at"`
	createdBy    int       `sql:"created_by"`
}
