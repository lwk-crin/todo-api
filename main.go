package main

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
)

type Todo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}

var Todos []Todo

func main() {
	Todos = []Todo{}
	Todos = append(Todos, Todo{ID: 1, Title: "Study", Completed: true})
	Todos = append(Todos, Todo{ID: 2, Title: "Cook", Completed: false})
	Todos = append(Todos, Todo{ID: 3, Title: "Sleep", Completed: true})

	http.HandleFunc("GET /todo", getTodos)
	http.HandleFunc("GET /todo/{id}", getTodoById)
	http.HandleFunc("POST /todo", createTodo)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
}

func writeError(w http.ResponseWriter, status int, err string, message string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)

	resp := ErrorResponse{
		Error:   err,
		Message: message,
	}

	json.NewEncoder(w).Encode(resp)
}

func getTodos(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(Todos)
}

func getTodoById(w http.ResponseWriter, r *http.Request) {
	for _, todo := range Todos {
		convID, err := strconv.Atoi(r.PathValue("id"))
		if err != nil {
			writeError(w, http.StatusBadRequest, "Invalid ID", "Todo ID must be a valid positive integer")
			return
		}
		if todo.ID == convID {
			json.NewEncoder(w).Encode(todo)
			return
		}
	}

	writeError(w, http.StatusNotFound, "ID Not Found", "The requested Todo ID was not found")
}

func createTodo(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Println(err)
		writeError(w, http.StatusInternalServerError, "Request Body Reading Failed", "Failed to read the request body")
		return
	}

	var m Todo
	if err := json.Unmarshal(body, &m); err != nil {
		log.Println(err)
		writeError(w, http.StatusInternalServerError, "Request Body Unmarshaling failed", "Failed to unmarshal/decode the request body")
		return
	}

	//TODO: Add Validation logic here

	Todos = append(Todos, m)
	w.WriteHeader(http.StatusCreated)
	w.Write(body)
}
