package todo

import (
	"database/sql"
	"log"
	"time"

	"github.com/google/uuid"
	_ "github.com/mattn/go-sqlite3"
)

var DB *sql.DB

// Inicializar la conexión a la base de datos
func InitDB(dataSourceName string) {
	var err error
	DB, err = sql.Open("sqlite3", dataSourceName)
	if err != nil {
		log.Fatal(err)
	}

	// Crear la tabla si no existe
	createTableSQL := `
	CREATE TABLE IF NOT EXISTS todos (
		id TEXT PRIMARY KEY,
		title TEXT,
		content TEXT,
		created_at DATETIME,
		updated_at DATETIME
	);
	`
	_, err = DB.Exec(createTableSQL)
	if err != nil {
		log.Fatal(err)
	}
}

// Obtener todas las tareas
func GetAllTodos() ([]Todo, error) {
	rows, err := DB.Query("SELECT id, title, content, created_at, updated_at FROM todos")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var todos []Todo
	for rows.Next() {
		var t Todo
		if err := rows.Scan(&t.ID, &t.Title, &t.Content, &t.CreatedAt, &t.UpdatedAt); err != nil {
			return nil, err
		}
		todos = append(todos, t)
	}
	return todos, nil
}

// Obtener una tarea por ID (UUID)
func GetById(id string) (Todo, error) {
	row := DB.QueryRow("SELECT id, title, content, created_at, updated_at FROM todos WHERE id = ?", id)

	var t Todo
	err := row.Scan(&t.ID, &t.Title, &t.Content, &t.CreatedAt, &t.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			return Todo{}, nil
		}
		return Todo{}, err
	}
	return t, nil
}

// Crear una nueva tarea
func CreateTodo(todo *Todo) error {
	todo.ID = uuid.New().String()
	todo.CreatedAt = time.Now()
	todo.UpdatedAt = time.Now()

	query := `INSERT INTO todos (id, title, content, created_at, updated_at)
	          VALUES (?, ?, ?, ?, ?)`
	_, err := DB.Exec(query, todo.ID, todo.Title, todo.Content, todo.CreatedAt, todo.UpdatedAt)
	return err
}

// Actualizar una tarea existente por UUID
func UpdateTodo(id string, title, content string) (Todo, error) {
	now := time.Now()
	_, err := DB.Exec("UPDATE todos SET title = ?, content = ?, updated_at = ? WHERE id = ?", title, content, now, id)
	if err != nil {
		return Todo{}, err
	}

	// Devolver la tarea actualizada
	todo, err := GetById(id)
	if err != nil {
		return Todo{}, err
	}
	return todo, nil
}

// Eliminar una tarea por UUID
func DeleteTodo(id string) error {
	_, err := DB.Exec("DELETE FROM todos WHERE id = ?", id)
	return err
}
