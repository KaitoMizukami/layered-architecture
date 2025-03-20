package main

import (
	"net/http"

	"github.com/KaitoMizukami/todo-app/handler"
	"github.com/KaitoMizukami/todo-app/repository"
	"github.com/KaitoMizukami/todo-app/service"
	"github.com/go-chi/chi/v5"
)

func main() {
	inMemoryRepo := repository.NewInMemoryRepository()
	todoService := service.NewTodoService(inMemoryRepo)
	todoHandler := handler.NewUserHandler(todoService)

	r := chi.NewRouter()

	r.Get("/todos", todoHandler.GetAllTodo)
	r.Get("/todos/{id}", todoHandler.GetTodoById)

	http.ListenAndServe(":8000", r)
}
