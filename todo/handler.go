package todo

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

// Obtener todas las tareas
func GetTodosHandler(w http.ResponseWriter, r *http.Request) {
	todos, err := GetAllTodos()
	if err != nil {
		http.Error(w, "Error al obtener tareas", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(todos)
}

// Obtener una tarea por ID
func GetTodoByIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	todo, err := GetById(id)
	if err != nil {
		http.Error(w, "Error al obtener tarea", http.StatusInternalServerError)
		return
	}
	if todo.ID == "" {
		http.NotFound(w, r)
		return
	}
	json.NewEncoder(w).Encode(todo)
}

// Crear nueva tarea
func CreateTodoHandler(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	if err := json.NewDecoder(r.Body).Decode(&todo); err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

	err := CreateTodo(&todo)
	if err != nil {
		http.Error(w, "Error al crear tarea", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(todo)
}

// Actualizar una tarea
func UpdateTodoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	var input Todo
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		http.Error(w, "Datos inválidos", http.StatusBadRequest)
		return
	}

	updatedTodo, err := UpdateTodo(id, input.Title, input.Content, input.IsDone)
	if err != nil {
		http.Error(w, "Error al actualizar tarea", http.StatusInternalServerError)
		return
	}
	json.NewEncoder(w).Encode(updatedTodo)
}

// Eliminar una tarea
func DeleteTodoHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	if err := DeleteTodo(id); err != nil {
		http.Error(w, "Error al eliminar tarea", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
