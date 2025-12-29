# 08 - Go Fiber CRM

This Go program implements a simple CRM (Customer Relationship Management) application using the Fiber web framework and GORM ORM with SQLite database. It manages leads with basic CRUD operations.

### 🔧 Imports

```go
import (
	"log"

	"github.com/RanitManik/go-projects/08-go-fiber-crm/database"
	"github.com/RanitManik/go-projects/08-go-fiber-crm/lead"
	"github.com/gofiber/fiber/v2"
)
```

- `log`: For logging.
- Custom `database` and `lead` packages.
- `github.com/gofiber/fiber/v2`: Fast web framework.

### 🎬 Data Structures

```go
type Lead struct {
	gorm.Model
	Name    string `json:"name"`
	Company string `json:"company"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
}
```

- `Lead` represents a CRM lead with contact information.

### 📦 Database

```go
func Connect() {
	db, err := gorm.Open(sqlite.Open("./data/leads.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database:", err)
	}

	DB = db
}
```

- Connects to SQLite database file `./data/leads.db`.

### 🧩 Handlers

#### ✅ Get All Leads

```go
func GetLeads(c *fiber.Ctx) error {
	var leads []Lead
	if err := database.DB.Find(&leads).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(leads)
}
```

- Retrieves all leads from database.

#### 🔍 Get a Lead by ID

```go
func GetLead(c *fiber.Ctx) error {
	id := c.Params("id")
	var lead Lead
	// ... find and return lead ...
}
```

- Finds lead by ID.

#### ➕ Create a New Lead

```go
func NewLead(c *fiber.Ctx) error {
	lead := new(Lead)
	if err := c.BodyParser(lead); err != nil {
		return c.Status(400).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if lead.Name == "" || lead.Email == "" {
		return c.Status(400).SendString("name and email are required")
	}

	if err := database.DB.Create(lead).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.Status(201).JSON(lead)
}
```

- Parses JSON body, validates required fields, creates lead.

#### ❌ Delete a Lead

```go
func DeleteLead(c *fiber.Ctx) error {
	id := c.Params("id")
	var lead Lead
	// ... find and delete lead ...
	return c.JSON(fiber.Map{
		"message": "lead deleted",
	})
}
```

- Deletes lead by ID.

### 🚀 Main Function

```go
func main() {
	database.Connect()
	database.DB.AutoMigrate(&lead.Lead{})

	app := fiber.New()
	setupRoutes(app)

	log.Fatal(app.Listen(":8080"))
}

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/leads", lead.GetLeads)
	app.Get("/api/v1/leads/:id", lead.GetLead)
	app.Post("/api/v1/leads", lead.NewLead)
	app.Delete("/api/v1/leads/:id", lead.DeleteLead)
}
```

- Connects database, migrates schema, sets up routes, starts server.

### 📝 Summary of API Endpoints

| Method | Endpoint            | Description         |
|--------|---------------------|---------------------|
| GET    | `/api/v1/leads`     | List all leads      |
| GET    | `/api/v1/leads/:id` | Get a lead by ID    |
| POST   | `/api/v1/leads`     | Create a new lead   |
| DELETE | `/api/v1/leads/:id` | Delete a lead by ID |