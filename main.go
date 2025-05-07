package main

import (
	"fmt"
	"go-todo-list/todo"
	"log"
	"net/http"
)

func main() {
	// Inicializar la base de datos
	todo.InitDB("todos.db")
	defer todo.DB.Close()

	// Configurar las rutas
	http.HandleFunc("/todos", todo.GetTodosHandler)         // Obtener todas las tareas
	http.HandleFunc("/todo", todo.GetTodoByIDHandler)       // Obtener una tarea por ID
	http.HandleFunc("/todo/create", todo.CreateTodoHandler) // Crear una tarea
	http.HandleFunc("/todo/update", todo.UpdateTodoHandler) // Actualizar una tarea
	http.HandleFunc("/todo/delete", todo.DeleteTodoHandler) // Eliminar una tarea

	// Arrancar el servidor
	fmt.Println("Servidor escuchando en :8080...")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
