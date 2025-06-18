# 02 - Build A CRUD API

This Go program implements a simple **RESTful API** for managing a list of movies using the Gorilla Mux router. Let's break it down step by step:

### 🔧 Imports

```go
import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"slices"
)
```

- `encoding/json`: For encoding/decoding JSON.
- `fmt`, `log`: For output and logging.
- `math/rand`, `strconv`: Used to generate a random ID.
- `net/http`: To build the web server.
- `github.com/gorilla/mux`: For routing with path parameters.
- `slices`: Provides operations on slices, including deletion (Go 1.21+).

### 🎬 Data Structures

```go
type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}
```

- `Movie` represents a movie record.
- `Director` is a nested structure inside `Movie`.
- JSON tags (`json:"..."`) define how fields are named in JSON.

### 📦 Global Variables

```go
var movies []Movie
```

- In-memory list of movies. Initially populated in `main()`.

### 🧩 Handlers

#### ✅ Get All Movies

```go
func getAllMovies(w http.ResponseWriter, r *http.Request)
```

- Responds with the entire `movies` list in JSON format.

#### 🔍 Get a Movie by ID

```go
func getMovie(w http.ResponseWriter, r *http.Request)
```

- Extracts `id` from URL path and finds the matching movie.

#### ➕ Create a New Movie

```go
func createMovie(w http.ResponseWriter, r *http.Request)
```

- Decodes request body into a `Movie`, generates a random `ID`, appends it to `movies`.

#### ✏️ Update a Movie

```go
func updateMovie(w http.ResponseWriter, r *http.Request)
```

- Finds movie by `id`, removes the old version using `slices.Delete`, and appends the new version with the same ID.

#### ❌ Delete a Movie

```go
func deleteMovie(w http.ResponseWriter, r *http.Request)
```

- Finds and removes the movie by `id` using `slices.Delete`.

### 🚀 Main Function

```go
func main() {
	r := mux.NewRouter()

	// Seed with two sample movies
	movies = append(movies, Movie{...})
	movies = append(movies, Movie{...})

	// Route mappings
	r.HandleFunc("/movies", getAllMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Println("Starting Server At PORT 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
```

- Registers all routes and starts the HTTP server on port `8080`.

### 📝 Summary of API Endpoints

| Method | Endpoint         | Description             |
|--------|------------------|-------------------------|
| GET    | `/movies`        | List all movies         |
| GET    | `/movies/{id}`   | Get a movie by ID       |
| POST   | `/movies`        | Add a new movie         |
| PUT    | `/movies/{id}`   | Update an existing movie|
| DELETE | `/movies/{id}`   | Delete a movie          |
