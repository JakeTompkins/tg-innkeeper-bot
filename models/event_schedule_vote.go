package models

import "time"

const EVENT_SCHEDULE_VOTE_TABLE = "event_schedule_votes"

type EventScheduleVote struct {
	Id        int       `sql:"id"`
	EventId   int       `sql:"event_id"`
	UserId    int       `sql:"user_id"`
	TimeStart time.Time `sql:"time_start"`
	TimeEnd   time.Time `sql:"time_end"`
}
