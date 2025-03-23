package domain

import (
	"errors"

	"github.com/google/uuid"
)

type task struct {
	id          uuid.UUID
	name        string
	description string
	completed   bool
}

type Task interface {
	GetID() uuid.UUID
	GetName() string
	GetDescription() string
	IsCompleted() bool
	Complete() error
}

func NewTask(title, description string, completed bool) (Task, error) {
	if title == "" {
		return nil, errors.New("title is required")
	}
	return &task{
		id:          uuid.New(),
		name:        title,
		description: description,
		completed:   completed,
	}, nil
}
func (t *task) GetID() uuid.UUID       { return t.id }
func (t *task) GetName() string        { return t.name }
func (t *task) GetDescription() string { return t.description }
func (t *task) IsCompleted() bool      { return t.completed }

func (t *task) Complete() error {
	if t.completed {
		return errors.New("Task is already completed")
	}
	t.completed = true
	return nil
}
