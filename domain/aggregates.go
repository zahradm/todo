package domain

import (
	"errors"

	"github.com/google/uuid"
)

type TodoList struct {
	id    uuid.UUID
	title string
	tasks []*Task
}

func NewTodoList(title string) (*TodoList, error) {
	if title == "" {
		return nil, errors.New("title list is required")
	}
	return &TodoList{
		id:    uuid.New(),
		title: title,
		tasks: []*Task{},
	}, nil
}

func (tl *TodoList) GetId() uuid.UUID  { return tl.id }
func (tl *TodoList) GetTitle() string  { return tl.title }
func (tl *TodoList) GetTasks() []*Task { return tl.tasks }

func (tl *TodoList) AddTask(id uuid.UUID, title string, description string, completed bool) (*TodoList, error) {
	task, error := NewTask(title, description, completed)
	if error != nil {
		return nil, errors.New("error creating task")
	}
	tl.tasks = append(tl.tasks, &task)
	return tl, nil
}

func (tl *TodoList) GetTask(id uuid.UUID) (*Task, error) {
	for _, task := range tl.tasks {
		if (*task).GetID() == id {
			return task, nil
		}
	}
	return nil, errors.New("task not found")
}
