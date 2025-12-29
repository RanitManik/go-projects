# 03 - MySQL Book Management System

This Go program implements a simple **RESTful API** for managing a list of books using MySQL and GORM. Let's break it down step by step:

### 🔧 Imports

```go
import (
	"fmt"
	"log"
	"net/http"

	"github.com/RanitManik/go-projects/03-book-management-system/pkg/routes"
	"github.com/gorilla/mux"
)
```

- `fmt`, `log`: For output and logging.
- `net/http`: To build the web server.
- `github.com/gorilla/mux`: For routing with path parameters.
- Custom routes package for registering routes.

### 🎬 Data Structures

```go
type Book struct {
	gorm.Model
	Name        string `gorm:"type:varchar(255)" json:"name"`
	Author      string `gorm:"type:varchar(255)" json:"author"`
	Publication string `gorm:"type:varchar(255)" json:"publication"`
}
```

- `Book` represents a book record with Name, Author, Publication.
- Uses GORM for ORM with MySQL.

### 📦 Global Variables

```go
var db *gorm.DB
```

- Global database connection.

### 🧩 Handlers

#### ✅ Get All Books

```go
func GetAllBooks(w http.ResponseWriter, r *http.Request)
```

- Responds with the entire books list in JSON format.

#### 🔍 Get a Book by ID

```go
func GetBookById(w http.ResponseWriter, r *http.Request)
```

- Extracts `id` from URL path and finds the matching book.

#### ➕ Create a New Book

```go
func CreateBook(w http.ResponseWriter, r *http.Request)
```

- Decodes request body into a `Book`, creates it in DB.

#### ✏️ Update a Book

```go
func UpdateBookById(w http.ResponseWriter, r *http.Request)
```

- Finds book by `id`, updates with new data.

#### ❌ Delete a Book

```go
func DeleteBookById(w http.ResponseWriter, r *http.Request)
```

- Finds and deletes the book by `id`.

### 🚀 Main Function

```go
func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)

	addr := ":8080"
	fmt.Println("Server running on http://localhost" + addr)

	log.Fatal(http.ListenAndServe(addr, r))
}
```

- Registers all routes and starts the HTTP server on port `8080`.

### 📝 Summary of API Endpoints

| Method | Endpoint         | Description             |
|--------|------------------|-------------------------|
| GET    | `/book/`         | List all books          |
| GET    | `/book/{id}`     | Get a book by ID        |
| POST   | `/book/`         | Add a new book          |
| PUT    | `/book/{id}`     | Update an existing book |
| DELETE | `/book/{id}`     | Delete a book           |