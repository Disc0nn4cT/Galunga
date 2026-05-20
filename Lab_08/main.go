package main

import (
	"fmt"
	"log"

	"lab8/models"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func sendJSON(c *fiber.Ctx, status int, data interface{ MarshalJSON() ([]byte, error) }) error {
	bytes, err := data.MarshalJSON()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).SendString("Помилка генерації JSON")
	}
	c.Type("json")
	return c.Status(status).Send(bytes)
}

func sendError(c *fiber.Ctx, status int, msg string) error {
	resp := models.ErrorResponse{Error: msg}
	return sendJSON(c, status, &resp)
}

func connectDatabase() {
	dsn := "host=localhost user=postgres password=secret dbname=contacts_db port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Помилка підключення до БД!\n", err)
	}

	db.AutoMigrate(&models.Contact{})
	DB = db
	fmt.Println("Успішно підключено до PostgreSQL!")
}

func main() {
	connectDatabase()
	app := fiber.New()

	app.Get("/contacts", func(c *fiber.Ctx) error {
		var contacts models.ContactList
		DB.Find(&contacts)
		return sendJSON(c, fiber.StatusOK, &contacts)
	})

	app.Get("/contacts/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		var contact models.Contact

		if err := DB.First(&contact, id).Error; err != nil {
			return sendError(c, fiber.StatusNotFound, "Контакт не знайдено")
		}
		return sendJSON(c, fiber.StatusOK, &contact)
	})

	app.Post("/contacts", func(c *fiber.Ctx) error {
		var contact models.Contact

		if err := contact.UnmarshalJSON(c.Body()); err != nil {
			return sendError(c, fiber.StatusBadRequest, "Недійсний JSON")
		}

		if contact.Name == "" || contact.Phone == "" {
			return sendError(c, fiber.StatusBadRequest, "Ім'я та телефон обов'язкові")
		}

		DB.Create(&contact)
		return sendJSON(c, fiber.StatusCreated, &contact)
	})

	app.Put("/contacts/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		var contact models.Contact

		if err := DB.First(&contact, id).Error; err != nil {
			return sendError(c, fiber.StatusNotFound, "Контакт не знайдено")
		}

		var updateData models.Contact
		if err := updateData.UnmarshalJSON(c.Body()); err != nil {
			return sendError(c, fiber.StatusBadRequest, "Недійсний JSON")
		}

		contact.Name = updateData.Name
		contact.Phone = updateData.Phone
		DB.Save(&contact)

		return sendJSON(c, fiber.StatusOK, &contact)
	})

	app.Delete("/contacts/:id", func(c *fiber.Ctx) error {
		id := c.Params("id")
		var contact models.Contact

		if err := DB.First(&contact, id).Error; err != nil {
			return sendError(c, fiber.StatusNotFound, "Контакт не знайдено")
		}

		DB.Delete(&contact)
		return c.SendStatus(fiber.StatusNoContent)
	})

	fmt.Println("Сервер запущено на http://localhost:3000")
	app.Listen(":3000")
}
