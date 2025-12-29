# 09 - Go Fiber HRMS

This Go program implements an HRMS (Human Resource Management System) using the Fiber web framework and MongoDB for data storage. It manages employee records with full CRUD operations.

### 🔧 Imports

```go
import (
	"context"
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)
```

- `context`, `log`, `time`: Standard libraries.
- `github.com/gofiber/fiber/v2`: Web framework.
- MongoDB driver packages for database operations.

### 🎬 Data Structures

```go
type Employee struct {
	ID     primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name   string             `json:"name"`
	Salary float64            `json:"salary"`
	Age    int                `json:"age"`
}
```

- `Employee` represents an employee with ID, Name, Salary, Age.

### 📦 Database Connection

```go
type MongoInstance struct {
	Client *mongo.Client
	DB     *mongo.Database
}

var mg MongoInstance

const dbName = "fiber-hrms"
const mongoURI = "mongodb://localhost:27017"

func Connect() error {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	client, err := mongo.Connect(
		ctx,
		options.Client().ApplyURI(mongoURI),
	)
	if err != nil {
		return err
	}

	if err := client.Ping(ctx, nil); err != nil {
		return err
	}

	mg = MongoInstance{
		Client: client,
		DB:     client.Database(dbName),
	}

	return nil
}
```

- Connects to local MongoDB instance.

### 🧩 Handlers

#### ✅ Get All Employees

```go
func getEmployees(c *fiber.Ctx) error {
	var employees []Employee

	cursor, err := mg.DB.Collection("employees").Find(c.Context(), bson.M{})
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}
	defer cursor.Close(c.Context())

	if err := cursor.All(c.Context(), &employees); err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	return c.JSON(employees)
}
```

- Retrieves all employees from MongoDB.

#### 🔍 Get an Employee by ID

```go
func getEmployee(c *fiber.Ctx) error {
	id := c.Params("id")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("invalid employee id")
	}

	var employee Employee
	err = mg.DB.Collection("employees").
		FindOne(c.Context(), bson.M{"_id": objectID}).
		Decode(&employee)
	// ... error handling and return ...
}
```

- Finds employee by ObjectID.

#### ➕ Create a New Employee

```go
func createEmployee(c *fiber.Ctx) error {
	var employee Employee

	if err := c.BodyParser(&employee); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("invalid JSON body")
	}

	if employee.Name == "" {
		return c.Status(fiber.StatusBadRequest).SendString("name is required")
	}
	if employee.Age <= 0 {
		return c.Status(fiber.StatusBadRequest).SendString("age must be greater than 0")
	}

	result, err := mg.DB.Collection("employees").InsertOne(c.Context(), employee)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	employee.ID = result.InsertedID.(primitive.ObjectID)

	return c.Status(fiber.StatusCreated).JSON(employee)
}
```

- Parses JSON, validates, inserts into MongoDB.

#### ✏️ Update an Employee

```go
func updateEmployee(c *fiber.Ctx) error {
	id := c.Params("id")
	// ... parse ID and body ...

	update := bson.M{}
	if updateData.Name != "" {
		update["name"] = updateData.Name
	}
	if updateData.Age > 0 {
		update["age"] = updateData.Age
	}
	if updateData.Salary > 0 {
		update["salary"] = updateData.Salary
	}

	result, err := mg.DB.Collection("employees").UpdateOne(
		c.Context(),
		bson.M{"_id": objectID},
		bson.M{"$set": update},
	)
	// ... error handling and return updated employee ...
}
```

- Updates employee fields selectively.

#### ❌ Delete an Employee

```go
func deleteEmployee(c *fiber.Ctx) error {
	id := c.Params("id")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("invalid employee id")
	}

	result, err := mg.DB.Collection("employees").DeleteOne(
		c.Context(),
		bson.M{"_id": objectID},
	)
	// ... error handling ...
	return c.SendStatus(fiber.StatusNoContent)
}
```

- Deletes employee by ID.

### 🚀 Main Function

```go
func main() {
	if err := Connect(); err != nil {
		log.Fatal(err)
	}

	defer func() {
		if err := mg.Client.Disconnect(context.Background()); err != nil {
			log.Println("Mongo disconnect error:", err)
		}
	}()

	app := fiber.New()

	app.Get("/employees", getEmployees)
	app.Get("/employee/:id", getEmployee)
	app.Post("/employee", createEmployee)
	app.Patch("/employee/:id", updateEmployee)
	app.Delete("/employee/:id", deleteEmployee)

	log.Fatal(app.Listen(":8080"))
}
```

- Connects to MongoDB, sets up routes, starts server.

### 📝 Summary of API Endpoints

| Method | Endpoint           | Description            |
|--------|--------------------|------------------------|
| GET    | `/employees`       | List all employees     |
| GET    | `/employee/:id`    | Get an employee by ID  |
| POST   | `/employee`        | Create a new employee  |
| PATCH  | `/employee/:id`    | Update an employee     |
| DELETE | `/employee/:id`    | Delete an employee     |