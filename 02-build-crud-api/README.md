# 02 - Build A CRUD API

A simple RESTful API built with Go that demonstrates basic CRUD (Create, Read, Update, Delete) operations for a movie resource. This project uses the `gorilla/mux` router for handling HTTP requests.

## Features

- List all movies
- Get a single movie by ID
- Create a new movie
- Update an existing movie
- Delete a movie

## Project Structure

- `main.go`: Main application logic for the CRUD API
- `go.mod`, `go.sum`: Go module files for dependency management

## How to Run

1. Navigate to the project directory:
   ```sh
   cd "02 - Build A CRUD API"
   ```
2. Download dependencies:
   ```sh
   go mod tidy
   ```
3. Run the server:
   ```sh
   go run main.go
   ```
4. The API will be available at [http://localhost:8080](http://localhost:8080)

## API Endpoints

- `GET    /movies` - List all movies
- `GET    /movies/{id}` - Get a movie by ID
- `POST   /movies` - Create a new movie (JSON body)
- `PUT    /movies/{id}` - Update a movie by ID (JSON body)
- `DELETE /movies/{id}` - Delete a movie by ID

## Example Movie JSON

```json
{
  "isbn": "12345",
  "title": "Example Movie",
  "director": {
    "firstname": "John",
    "lastname": "Doe"
  }
}
```

---

This project is part of my Go learning journey!
