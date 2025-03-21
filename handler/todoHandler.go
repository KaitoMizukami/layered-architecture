package handler

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/KaitoMizukami/todo-app/models"
	"github.com/KaitoMizukami/todo-app/service"
	"github.com/go-chi/chi/v5"
)

type TodoHandler struct {
	ts service.TodoService
}

func NewTodoHandler(ts service.TodoService) *TodoHandler {
	return &TodoHandler{
		ts: ts,
	}
}

func (th *TodoHandler) GetAllTodo(w http.ResponseWriter, r *http.Request) {
	todos := th.ts.GetTodos()
	json.NewEncoder(w).Encode(todos)
}

func (th *TodoHandler) GetTodoById(w http.ResponseWriter, r *http.Request) {
	strId := chi.URLParam(r, "id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		http.Error(w, "id should be integer", http.StatusBadRequest)
		return
	}

	todo, err := th.ts.GetTodoByID(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(todo)
}

func (th *TodoHandler) CreateTodo(w http.ResponseWriter, r *http.Request) {
	var reqTodo models.Todo
	json.NewDecoder(r.Body).Decode(&reqTodo)

	newTodo, err := th.ts.CreateTodo(reqTodo.Title)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	json.NewEncoder(w).Encode(newTodo)
}

func (th *TodoHandler) UpdateStatus(w http.ResponseWriter, r *http.Request) {
	strId := chi.URLParam(r, "id")
	id, err := strconv.Atoi(strId)
	if err != nil {
		http.Error(w, "id should be integer", http.StatusBadRequest)
		return
	}

	todo, err := th.ts.UpdateStatus(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(todo)
}
