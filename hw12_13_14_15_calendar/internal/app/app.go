package app

import (
	"context"

	"github.com/bon3o/otus-hw-01/hw12_13_14_15_calendar/internal/storage"
)

type App struct {
	log Logger
	s   Storage
}
type Logger interface { // TODO
}

type Storage interface {
	Create(e storage.Event) error
	Update(e storage.Event) error
	Delete(e storage.Event) error
	FindAll() ([]storage.Event, error)
	GetByID(id string) (storage.Event, error)
}

func New(logger Logger, storage Storage) *App {
	return &App{
		log: logger,
		s:   storage,
	}
}

func (a *App) CreateEvent(ctx context.Context, id, title string) error {
	return a.s.Create(storage.Event{ID: id, Title: title})
}

// TODO
