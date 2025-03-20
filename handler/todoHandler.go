package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/KaitoMizukami/todo-app/service"
	"github.com/go-chi/chi/v5"
)

type UserHandler struct {
	ts service.TodoService
}

func NewUserHandler(ts service.TodoService) *UserHandler {
	return &UserHandler{
		ts: ts,
	}
}

func (uh *UserHandler) GetAllTodo(w http.ResponseWriter, r *http.Request) {
	todos := uh.ts.GetTodos()
	json.NewEncoder(w).Encode(todos)
}

func (uh *UserHandler) GetTodoById(w http.ResponseWriter, r *http.Request) {
	strId := chi.URLParam(r, "id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		http.Error(w, "id should be integer", http.StatusBadRequest)
		return
	}

	todo, err := uh.ts.GetTodoByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(todo)
}
