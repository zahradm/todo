package application

import (
	"github.com/google/uuid"
	"github.com/zahradm/todo/domain"
)

type TodoListService struct {
	repo          domain.TodoListRepository
	domainService *domain.TodoDomainService
}

func NewTodoListService(repo domain.TodoListRepository, domainService *domain.TodoDomainService) *TodoListService {
	return &TodoListService{repo: repo, domainService: domainService}
}

func (s *TodoListService) CreateTodoList(title string) (*domain.TodoList, error) {
	tl, err := domain.NewTodoList(title)
	if err != nil {
		return nil, err
	}
	if err := s.repo.Save(tl); err != nil {
		return nil, err
	}
	return tl, nil
}

func (s *TodoListService) AddTaskToList(listId uuid.UUID, title string, description string, completed bool) (*domain.TodoList, error) {
	tl, err := s.repo.FindByID(listId)
	if err != nil {
		return nil, err
	}
	taskID := uuid.New()
	todo, err := tl.AddTask(taskID, title, description, completed)
	if err != nil {
		return nil, err
	}
	if err := s.repo.Save(todo); err != nil {
		return nil, err
	}
	return todo, nil
}

func (s *TodoListService) CompleteTask(listID, taskID uuid.UUID) error {
	tl, err := s.repo.FindByID(listID)
	if err != nil {
		return err
	}
	task, err := tl.GetTask(taskID)
	if err != nil {
		return err
	}
	if err := (*task).Complete(); err != nil {
		return err
	}
	return s.repo.Save(tl)
}
