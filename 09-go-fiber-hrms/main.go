package main

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

type MongoInstance struct {
	Client *mongo.Client
	DB     *mongo.Database
}

var mg MongoInstance

const dbName = "fiber-hrms"
const mongoURI = "mongodb://localhost:27017"

type Employee struct {
	ID     primitive.ObjectID `json:"id,omitempty" bson:"_id,omitempty"`
	Name   string             `json:"name"`
	Salary float64            `json:"salary"`
	Age    int                `json:"age"`
}

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

func createEmployee(c *fiber.Ctx) error {
	collection := mg.DB.Collection("employees")

	var employee Employee

	if err := c.BodyParser(&employee); err != nil {
		return c.Status(fiber.StatusBadRequest).
			SendString("invalid JSON body")
	}

	if employee.Name == "" {
		return c.Status(fiber.StatusBadRequest).
			SendString("name is required")
	}
	if employee.Age <= 0 {
		return c.Status(fiber.StatusBadRequest).
			SendString("age must be greater than 0")
	}

	result, err := collection.InsertOne(c.Context(), employee)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			SendString(err.Error())
	}

	employee.ID = result.InsertedID.(primitive.ObjectID)

	return c.Status(fiber.StatusCreated).JSON(employee)
}

func getEmployee(c *fiber.Ctx) error {
	id := c.Params("id")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			SendString("invalid employee id")
	}

	var employee Employee

	err = mg.DB.Collection("employees").
		FindOne(c.Context(), bson.M{"_id": objectID}).
		Decode(&employee)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.Status(fiber.StatusNotFound).
				SendString("employee not found")
		}
		return c.Status(fiber.StatusInternalServerError).
			SendString(err.Error())
	}

	return c.JSON(employee)
}

func updateEmployee(c *fiber.Ctx) error {
	id := c.Params("id")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("invalid employee id")
	}

	var updateData Employee
	if err := c.BodyParser(&updateData); err != nil {
		return c.Status(fiber.StatusBadRequest).SendString("invalid JSON body")
	}

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

	if len(update) == 0 {
		return c.Status(fiber.StatusBadRequest).SendString("no valid fields to update")
	}

	result, err := mg.DB.Collection("employees").UpdateOne(
		c.Context(),
		bson.M{"_id": objectID},
		bson.M{"$set": update},
	)

	if err != nil {
		c.Status(fiber.StatusInternalServerError).SendString(err.Error())
	}

	if result.MatchedCount == 0 {
		return c.Status(fiber.StatusNotFound).SendString("employee not found")
	}

	var updatedEmployee Employee
	err = mg.DB.Collection("employees").
		FindOne(c.Context(), bson.M{"_id": objectID}).
		Decode(&updatedEmployee)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			SendString(err.Error())
	}

	return c.JSON(updatedEmployee)
}

func deleteEmployee(c *fiber.Ctx) error {
	id := c.Params("id")

	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).
			SendString("invalid employee id")
	}

	result, err := mg.DB.Collection("employees").DeleteOne(
		c.Context(),
		bson.M{"_id": objectID},
	)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).
			SendString(err.Error())
	}

	if result.DeletedCount == 0 {
		return c.Status(fiber.StatusNotFound).
			SendString("employee not found")
	}

	return c.SendStatus(fiber.StatusNoContent)
}
