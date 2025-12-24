package lead

import (
	"errors"

	"github.com/RanitManik/go-projects/08-go-fiber-crm/database"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type Lead struct {
	gorm.Model
	Name    string `json:"name"`
	Company string `json:"company"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
}

func GetLeads(c *fiber.Ctx) error {
	var leads []Lead
	if err := database.DB.Find(&leads).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(leads)
}

func GetLead(c *fiber.Ctx) error {
	id := c.Params("id")

	var lead Lead
	if err := database.DB.First(&lead, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(404).SendString("Lead not found")
		}
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(lead)
}

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

func DeleteLead(c *fiber.Ctx) error {
	id := c.Params("id")

	var lead Lead
	if err := database.DB.First(&lead, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return c.Status(404).SendString("Lead not found")
		}
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	if err := database.DB.Delete(&lead, id).Error; err != nil {
		return c.Status(500).JSON(fiber.Map{
			"error": err.Error(),
		})
	}

	return c.JSON(fiber.Map{
		"message": "lead deleted",
	})
}
