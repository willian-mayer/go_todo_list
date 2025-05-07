package main

import (
	"fmt"
	"go-todo-list/todo"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func main() {
	// Inicializar la base de datos
	todo.InitDB("todos.db")
	defer todo.DB.Close()

	// Crear router
	r := mux.NewRouter()

	// Definir rutas
	r.HandleFunc("/todo", todo.GetTodosHandler).Methods("GET")           // Obtener todas las tareas
	r.HandleFunc("/todo/{id}", todo.GetTodoByIDHandler).Methods("GET")   // Obtener una tarea por ID
	r.HandleFunc("/todo", todo.CreateTodoHandler).Methods("POST")        // Crear una tarea
	r.HandleFunc("/todo/{id}", todo.UpdateTodoHandler).Methods("PUT")    // Actualizar una tarea
	r.HandleFunc("/todo/{id}", todo.DeleteTodoHandler).Methods("DELETE") // Eliminar una tarea

	// Configurar CORS
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"}, // En producción usa tu dominio específico
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})

	// Iniciar servidor con CORS
	handler := c.Handler(r)

	fmt.Println("Servidor escuchando en :8080...")
	log.Fatal(http.ListenAndServe(":8080", handler))
}
