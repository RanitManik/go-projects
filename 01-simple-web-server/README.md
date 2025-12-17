# 01 - A Simple Web Server

A basic web server built with Go that demonstrates serving static files, handling GET and POST requests, and simple form processing.

## Features

- Serve static HTML files
- Handle GET and POST requests
- Simple form submission and processing

## Project Structure

- `main.go`: Main application logic for the web server
- `static/index.html`: Example static file served at the root
- `static/form.html`: HTML form for submitting data to the server

## How to Run

1. Navigate to the project directory:
   ```sh
   cd "01 - A Simple Web Server"
   ```
2. Run the server:
   ```sh
   go run main.go
   ```
3. Open your browser and visit:
   - [http://localhost:8080/](http://localhost:8080/) for the static page
   - [http://localhost:8080/form](http://localhost:8080/form) for the form (submit with POST)
   - [http://localhost:8080/hello](http://localhost:8080/hello) for a simple hello endpoint

---

This project is part of my Go learning journey!
