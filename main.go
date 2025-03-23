package main

import (
	"fmt"

	"github.com/zahradm/todo/application"
	"github.com/zahradm/todo/domain"
	"github.com/zahradm/todo/infrastructure"
)

func main() {
	repo := infrastructure.NewMemoryTodoListRepository()
	domainSvc := domain.NewTodoDomainService(repo)
	appSvc := application.NewTodoListService(repo, domainSvc)

	listTodo, _ := appSvc.CreateTodoList("Shopping")
	task, _ := appSvc.AddTaskToList(listTodo.GetId(), "Buy milk", "1L of milk", false)
	appSvc.CompleteTask(listTodo.GetId(), task.GetId())

	tl, _ := repo.FindByID(listTodo.GetId())
	fmt.Printf("List: %s, Items: %d\n", tl.GetTitle(), len(tl.GetTasks()))
	for _, item := range tl.GetTasks() {
		fmt.Printf(" - %s (Completed: %v)\n", (*item).GetName(), (*item).IsCompleted())
	}
}
