package main

import (
	"log"

	"github.com/RanitManik/go-projects/08-go-fiber-crm/database"
	"github.com/RanitManik/go-projects/08-go-fiber-crm/lead"
	"github.com/gofiber/fiber/v2"
)

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/leads", lead.GetLeads)
	app.Get("/api/v1/leads/:id", lead.GetLead)
	app.Post("/api/v1/leads", lead.NewLead)
	app.Delete("/api/v1/leads/:id", lead.DeleteLead)
}

func main() {
	database.Connect()

	database.DB.AutoMigrate(&lead.Lead{})

	app := fiber.New()
	setupRoutes(app)

	log.Fatal(app.Listen(":8080"))
}
