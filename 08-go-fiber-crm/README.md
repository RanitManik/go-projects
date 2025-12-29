# 08 - Go Fiber CRM

A CRM (Customer Relationship Management) application built with Go Fiber that manages leads. This project uses Fiber as the web framework and GORM with SQLite for data persistence.

## Features

- List all leads
- Get a single lead by ID
- Create a new lead
- Delete a lead
- SQLite database storage

## Project Structure

- `main.go`: Main application and route setup
- `database/database.go`: Database connection and configuration
- `lead/lead.go`: Lead model and CRUD operations
- `data/`: Directory for SQLite database file
- `go-fiber-crm.postman_collection.json`: Postman collection for API testing
- `go.mod`, `go.sum`: Go module files

## How to Run

1. Navigate to the project directory:
   ```sh
   cd "08-go-fiber-crm"
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

- `GET    /api/v1/leads` - List all leads
- `GET    /api/v1/leads/{id}` - Get a lead by ID
- `POST   /api/v1/leads` - Create a new lead (JSON body)
- `DELETE /api/v1/leads/{id}` - Delete a lead by ID

## Example Lead JSON

```json
{
  "name": "John Doe",
  "company": "Example Corp",
  "email": "john@example.com",
  "phone": "+1234567890"
}
```

## Database

The application uses SQLite and creates a database file at `./data/leads.db`. The schema is automatically migrated on startup.

---

This project is part of my Go learning journey!