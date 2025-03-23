package repository

import "github.com/zahradm/todo/domain"

type TodoListRepository interface {
	Save(domain.TodoList) error
	FindByID(id string) (domain.TodoList, error)
}
