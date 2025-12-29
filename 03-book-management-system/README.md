# 03 - MySQL Book Management System

A RESTful API built with Go that demonstrates basic CRUD (Create, Read, Update, Delete) operations for a book resource using MySQL and GORM. This project uses the `gorilla/mux` router for handling HTTP requests.

## Features

- List all books
- Get a single book by ID
- Create a new book
- Update an existing book
- Delete a book

## Project Structure

- `cmd/main.go`: Main application logic for the CRUD API
- `pkg/config/`: Database configuration
- `pkg/controllers/`: HTTP handlers for book operations
- `pkg/models/`: Book model and database operations
- `pkg/routes/`: Route definitions
- `go.mod`, `go.sum`: Go module files for dependency management

## How to Run

1. Navigate to the project directory:
   ```sh
   cd "03-book-management-system"
   ```
2. Download dependencies:
   ```sh
   go mod tidy
   ```
3. Set up MySQL database and update config
4. Run the server:
   ```sh
   go run cmd/main.go
   ```
5. The API will be available at [http://localhost:8080](http://localhost:8080)

## API Endpoints

- `GET    /book/` - List all books
- `GET    /book/{id}` - Get a book by ID
- `POST   /book/` - Create a new book (JSON body)
- `PUT    /book/{id}` - Update a book by ID (JSON body)
- `DELETE /book/{id}` - Delete a book by ID

## Example Book JSON

```json
{
  "name": "Book Title",
  "author": "Author Name",
  "publication": "Publisher"
}
```

---

This project is part of my Go learning journey!