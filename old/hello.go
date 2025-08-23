package hello

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"sync"
)

type Todo struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
	Done bool   `json:"done"`
}

var (
	todos  = []Todo{}
	nextID = 1
	mutex  sync.Mutex
)

func main() {
	http.HandleFunc("/todos", handleTodos)
	http.HandleFunc("/todos/", handleTodoByID)

	fmt.Println("Server started at http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// GET /todos — список задач
// POST /todos — добавить задачу
func handleTodos(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		mutex.Lock()
		defer mutex.Unlock()
		json.NewEncoder(w).Encode(todos)

	case "POST":
		var newTodo Todo
		if err := json.NewDecoder(r.Body).Decode(&newTodo); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		mutex.Lock()
		newTodo.ID = nextID
		nextID++
		todos = append(todos, newTodo)
		mutex.Unlock()

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(newTodo)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

// GET /todos/{id} — получить задачу
// PUT /todos/{id} — обновить задачу
// DELETE /todos/{id} — удалить задачу
func handleTodoByID(w http.ResponseWriter, r *http.Request) {
	idStr := r.URL.Path[len("/todos/"):]
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return
	}

	mutex.Lock()
	defer mutex.Unlock()

	index := -1
	for i, t := range todos {
		if t.ID == id {
			index = i
			break
		}
	}

	if index == -1 {
		http.Error(w, "Not found", http.StatusNotFound)
		return
	}

	switch r.Method {
	case "GET":
		json.NewEncoder(w).Encode(todos[index])

	case "PUT":
		var updated Todo
		if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		updated.ID = id
		todos[index] = updated
		json.NewEncoder(w).Encode(updated)

	case "DELETE":
		todos = append(todos[:index], todos[index+1:]...)
		w.WriteHeader(http.StatusNoContent)

	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}
