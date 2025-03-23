package domain

import (
	"github.com/google/uuid"
)

type TodoListRepository interface {
	FindByID(id uuid.UUID) (*TodoList, error)
	Save(tl *TodoList) error
}

type TodoDomainService struct {
	repo TodoListRepository
}

func NewTodoDomainService(repo TodoListRepository) *TodoDomainService {
	return &TodoDomainService{repo: repo}
}

func (s *TodoDomainService) IsTitleUnique(title string) (bool, error) {
	return true, nil
}
