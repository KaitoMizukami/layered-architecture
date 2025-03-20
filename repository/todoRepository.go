package repository

import (
	"fmt"
	"time"

	"github.com/KaitoMizukami/todo-app/models"
)

type PersistentRepository interface {
	GetTodos() []models.Todo
	GetTodoByID(id int) (models.Todo, error)
	CreateTodo(title string) (models.Todo, error)
	UpdateStatus(id int) (models.Todo, error)
}

type inMemoryRepository struct {
	todos []models.Todo
	id    int
}

func NewInMemoryRepository() *inMemoryRepository {
	return &inMemoryRepository{
		todos: []models.Todo{},
		id:    1,
	}
}

func (imr *inMemoryRepository) GetTodos() []models.Todo {
	return imr.todos
}

func (imr *inMemoryRepository) GetTodoByID(id int) (models.Todo, error) {
	for _, todo := range imr.todos {
		if todo.ID == id {
			return todo, nil
		}
	}

	return models.Todo{}, fmt.Errorf("item with id %d not found", id)
}

func (imr *inMemoryRepository) CreateTodo(title string) (models.Todo, error) {
	newTodo := models.Todo{
		ID:           imr.id,
		Title:        title,
		HasCompleted: false,
		CreatedAt:    time.Now(),
	}
	imr.todos = append(imr.todos, newTodo)
	imr.id += 1
	return newTodo, nil
}

func (imr *inMemoryRepository) UpdateStatus(id int) (models.Todo, error) {
	for i, todo := range imr.todos {
		if todo.ID == id {
			imr.todos[i].HasCompleted = !todo.HasCompleted
			return todo, nil
		}
	}

	return models.Todo{}, fmt.Errorf("item with id %d not found", id)
}
