package tests

import (
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/assert"
	"github.com/zahradm/todo/domain"
	"github.com/zahradm/todo/infrastructure"
)

func TestTodoListAggregate(t *testing.T) {
	repo := infrastructure.NewMemoryTodoListRepository()
	tl, _ := domain.NewTodoList("Shopping")
	assert.Equal(t, "Shopping", tl.GetTitle())
	assert.Empty(t, tl.GetTasks())

	todoID := uuid.New()
	_, _ = tl.AddTask(todoID, "Buy milk", "1L of milk", false)
	assert.Len(t, tl.GetTasks(), 1)
	assert.Equal(t, "Buy milk", (*tl.GetTasks()[0]).GetName())

	err := repo.Save(tl)
	assert.NoError(t, err)

	loaded, _ := repo.FindByID(tl.GetId())
	assert.Equal(t, tl.GetId(), loaded.GetId())
	assert.Len(t, loaded.GetTasks(), 1)
}
