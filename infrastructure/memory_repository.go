package infrastructure

import (
	"errors"
	"sync"

	"github.com/zahradm/todo/domain"

	"github.com/google/uuid"
)

type MemoryTodoListRepository struct {
	mu   sync.RWMutex
	list map[uuid.UUID]*domain.TodoList
}

func NewMemoryTodoListRepository() *MemoryTodoListRepository {
	return &MemoryTodoListRepository{
		list: make(map[uuid.UUID]*domain.TodoList)}
}

func (r *MemoryTodoListRepository) Save(tl *domain.TodoList) error {
	r.mu.Lock()
	defer r.mu.Unlock()
	r.list[tl.GetId()] = tl
	return nil
}

func (r *MemoryTodoListRepository) FindByID(id uuid.UUID) (*domain.TodoList, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()
	tl, ok := r.list[id]
	if !ok {
		return nil, errors.New("list not found")
	}
	return tl, nil
}
