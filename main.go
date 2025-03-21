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
	todoHandler := handler.NewTodoHandler(todoService)

	r := chi.NewRouter()

	r.Get("/todos", todoHandler.GetAllTodo)
	r.Get("/todos/{id}", todoHandler.GetTodoById)
	r.Post("/todos", todoHandler.CreateTodo)
	r.Put("/todos/{id}", todoHandler.UpdateStatus)

	http.ListenAndServe(":8000", r)
}
