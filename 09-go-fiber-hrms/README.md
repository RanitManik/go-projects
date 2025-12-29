# 09 - Go Fiber HRMS

An HRMS (Human Resource Management System) built with Go Fiber and MongoDB for managing employee records. This project demonstrates full CRUD operations with a NoSQL database.

## Features

- List all employees
- Get a single employee by ID
- Create a new employee
- Update an existing employee
- Delete an employee
- MongoDB integration

## Project Structure

- `main.go`: Main application with all handlers and route setup
- `go-fiber-hrms.postman_collection.json`: Postman collection for API testing
- `go.mod`, `go.sum`: Go module files for dependency management

## Prerequisites

- MongoDB running locally on `mongodb://localhost:27017`
- Database name: `fiber-hrms`
- Collection: `employees`

## How to Run

1. Navigate to the project directory:
   ```sh
   cd "09-go-fiber-hrms"
   ```
2. Start MongoDB (if not running):
   ```sh
   mongod
   ```
3. Download dependencies:
   ```sh
   go mod tidy
   ```
4. Run the server:
   ```sh
   go run main.go
   ```
5. The API will be available at [http://localhost:8080](http://localhost:8080)

## API Endpoints

- `GET    /employees` - List all employees
- `GET    /employee/{id}` - Get an employee by ID
- `POST   /employee` - Create a new employee (JSON body)
- `PATCH  /employee/{id}` - Update an employee (JSON body)
- `DELETE /employee/{id}` - Delete an employee

## Example Employee JSON

```json
{
  "name": "John Doe",
  "salary": 50000.0,
  "age": 30
}
```

## Database

The application connects to MongoDB and uses the `employees` collection. Employee IDs are MongoDB ObjectIDs.

---

This project is part of my Go learning journey!