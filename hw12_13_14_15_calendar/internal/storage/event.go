package storage

import (
	"errors"
	"time"
)

var (
	ErrEventNotFound = errors.New("event not found with id '%s'")
	ErrAlreadyExist  = errors.New("event with id '%s' already exist")
)

type Event struct {
	ID          string        `db:"id"`
	Title       string        `db:"title"`
	Description string        `db:"description"`
	StartAt     time.Time     `db:"start_at"`
	EndAt       time.Time     `db:"end_at"`
	UserID      int           `db:"user_id"`
	RemindFor   time.Duration `db:"remind_for"`
}
