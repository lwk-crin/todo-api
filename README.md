# Go Todo API

A minimal **RESTful Todo API written in Go** using the standard `net/http` package.
This project is primarily a **learning project** focused on understanding how HTTP APIs work internally in Go without using external frameworks.

The goal of this repository is to demonstrate:

* Basic REST API design
* HTTP routing using the Go standard library
* JSON encoding/decoding
* Proper HTTP status codes
* Error responses in JSON
* Simple in-memory data storage

> ⚠️ This project is intentionally simple and **not production ready**. It stores data in memory and does not persist anything to disk or a database.

---

# Features

* Retrieve all todos
* Retrieve a todo by ID
* Create a new todo
* JSON error responses
* Proper HTTP status codes
* Simple in-memory storage

---

# Tech Stack

* **Language:** Go
* **HTTP:** `net/http`
* **JSON:** `encoding/json`
* **Storage:** In-memory slice

No external libraries are used.

---

# Project Structure

```
.
├── main.go
└── README.md
```

All API logic currently exists inside `main.go`.

---

# Todo Model

The API uses the following data structure:

```go
type Todo struct {
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}
```

Example JSON representation:

```json
{
  "id": 1,
  "title": "Study",
  "completed": true
}
```

---

# Error Response Format

All API errors return JSON in the following format:

```json
{
  "error": "Invalid ID",
  "message": "Todo ID must be a valid positive integer"
}
```

Error structure in Go:

```go
type ErrorResponse struct {
	Error   string `json:"error"`
	Message string `json:"message"`
}
```

---

# API Endpoints

## Get All Todos

**Endpoint**

```
GET /todo
```

**Description**

Returns all todos stored in memory.

**Example Response**

```json
[
  {
    "id": 1,
    "title": "Study",
    "completed": true
  },
  {
    "id": 2,
    "title": "Cook",
    "completed": false
  }
]
```

---

## Get Todo By ID

**Endpoint**

```
GET /todo/{id}
```

Example:

```
GET /todo/1
```

**Success Response**

```json
{
  "id": 1,
  "title": "Study",
  "completed": true
}
```

**Error Responses**

Invalid ID:

```json
{
  "error": "Invalid ID",
  "message": "Todo ID must be a valid positive integer"
}
```

ID not found:

```json
{
  "error": "ID Not Found",
  "message": "The requested Todo ID was not found"
}
```

---

## Create Todo

**Endpoint**

```
POST /todo
```

**Request Body**

```json
{
  "id": 4,
  "title": "Read a book",
  "completed": false
}
```

**Response**

Status code:

```
201 Created
```

Returns the created todo.

---

# Running the Project

### 1. Clone the repository

```bash
git clone https://github.com/yourusername/todo-api-go.git
cd todo-api-go
```

### 2. Run the server

```bash
go run main.go
```

The server will start on:

```
http://localhost:8080
```

---

# Example Requests

### Get all todos

```bash
curl http://localhost:8080/todo
```

### Get todo by id

```bash
curl http://localhost:8080/todo/1
```

### Create a todo

```bash
curl -X POST http://localhost:8080/todo \
-H "Content-Type: application/json" \
-d '{"id":4,"title":"Read","completed":false}'
```

---

# Current Limitations

This API is intentionally minimal and currently has several limitations:

* No request validation
* No automatic ID generation
* No database (data stored in memory)
* No concurrency protection
* No authentication
* No update/delete endpoints

These improvements are planned as the project evolves.

---

# Planned Improvements

Future enhancements may include:

* Input validation
* Auto-increment IDs
* Persistent storage (SQLite/Postgres)
* Update Todo (`PUT /todo/{id}`)
* Delete Todo (`DELETE /todo/{id}`)
* Middleware support
* Logging improvements
* Concurrency safety using mutexes
* API tests

---

# Learning Goals

This project is meant to help understand:

* How Go handles HTTP servers internally
* How routing works with `http.HandleFunc`
* JSON serialization and deserialization
* REST API conventions
* Error handling patterns in Go APIs

---

# Author

crin
