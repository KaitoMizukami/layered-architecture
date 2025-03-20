package service

import (
	"errors"

	"github.com/KaitoMizukami/todo-app/models"
	"github.com/KaitoMizukami/todo-app/repository"
)

type TodoService interface {
	GetTodos() []models.Todo
	GetTodoByID(id int) (models.Todo, error)
	CreateTodo(title string) (models.Todo, error)
	UpdateStatus(id int) (models.Todo, error)
}

type todoService struct {
	pr repository.PersistentRepository
}

func NewTodoService(repo repository.PersistentRepository) *todoService {
	return &todoService{
		pr: repo,
	}
}

func (ts *todoService) GetTodos() []models.Todo {
	return ts.pr.GetTodos()
}

func (ts *todoService) GetTodoByID(id int) (models.Todo, error) {
	return ts.pr.GetTodoByID(id)
}

func (ts *todoService) CreateTodo(title string) (models.Todo, error) {
	if len(title) == 0 {
		return models.Todo{}, errors.New("title length should be greater than 1")
	}

	return ts.pr.CreateTodo(title)
}

func (ts *todoService) UpdateStatus(id int) (models.Todo, error) {
	return ts.pr.UpdateStatus(id)
}
