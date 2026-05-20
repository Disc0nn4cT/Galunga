package main

import (
	"encoding/json"
	"fmt"
	"lab7/models"
	"strconv"
	"sync"

	"github.com/gofiber/fiber/v2"
)

var (
	notes  = make(map[int]models.Note)
	nextID = 1
	mu     sync.Mutex
)

func main() {
	app := fiber.New()

	app.Get("/notes", func(c *fiber.Ctx) error {
		mu.Lock()
		defer mu.Unlock()

		result := []models.Note{}
		for _, note := range notes {
			result = append(result, note)
		}
		return c.Status(fiber.StatusOK).JSON(result)
	})

	app.Get("/notes/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{Error: "Некоректний ID"})
		}

		mu.Lock()
		note, exists := notes[id]
		mu.Unlock()

		if !exists {
			return c.Status(fiber.StatusNotFound).JSON(models.ErrorResponse{Error: "Нотатку не знайдено"})
		}

		return c.Status(fiber.StatusOK).JSON(note)
	})

	app.Post("/notes", func(c *fiber.Ctx) error {
		var note models.Note

		if err := json.Unmarshal(c.Body(), &note); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{Error: "Недійсний JSON"})
		}

		if note.Title == "" {
			return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{Error: "Title не може бути порожнім"})
		}

		mu.Lock()
		note.ID = nextID
		notes[nextID] = note
		nextID++
		mu.Unlock()

		return c.Status(fiber.StatusCreated).JSON(note)
	})

	app.Put("/notes/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{Error: "Некоректний ID"})
		}

		var updatedNote models.Note
		if err := json.Unmarshal(c.Body(), &updatedNote); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{Error: "Недійсний JSON"})
		}

		if updatedNote.Title == "" {
			return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{Error: "Title не може бути порожнім"})
		}

		mu.Lock()
		defer mu.Unlock()

		if _, exists := notes[id]; !exists {
			return c.Status(fiber.StatusNotFound).JSON(models.ErrorResponse{Error: "Нотатку не знайдено"})
		}

		updatedNote.ID = id
		notes[id] = updatedNote

		return c.Status(fiber.StatusOK).JSON(updatedNote)
	})

	app.Delete("/notes/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(models.ErrorResponse{Error: "Некоректний ID"})
		}

		mu.Lock()
		defer mu.Unlock()

		if _, exists := notes[id]; !exists {
			return c.Status(fiber.StatusNotFound).JSON(models.ErrorResponse{Error: "Нотатку не знайдено"})
		}

		delete(notes, id)
		return c.SendStatus(fiber.StatusNoContent)
	})

	fmt.Println("🚀 Сервер запущено на http://localhost:3000")
	app.Listen(":3000")
}
