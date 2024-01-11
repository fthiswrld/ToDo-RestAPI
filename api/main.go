package main

import (
	"github.com/gofiber/fiber/v2"
	"strconv"
)

type Task struct {
	ID    int    `json:"id"`
	Title string `json:"title"`
	Done  bool   `json:"done"`
}

var tasks []Task

func main() {
	app := fiber.New()

	app.Get("/api/tasks", func(c *fiber.Ctx) error {
		return c.JSON(tasks)
	})

	app.Post("/api/tasks", func(c *fiber.Ctx) error {
		var task Task
		err := c.BodyParser(&task)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "неверное тело задачи"})
		}
		task.ID = len(tasks) + 1
		tasks = append(tasks, task)
		return c.JSON(task)
	})

	app.Put("/api/tasks/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "неверный формат ID"})
		}
		var updatedTaskIndex = -1
		for k, v := range tasks {
			if v.ID == id {
				updatedTaskIndex = k
				break
			}
		}
		if updatedTaskIndex == -1 {
			return c.Status(404).JSON(fiber.Map{"error": "Задача не найдена"})
		}
		var updatedTask Task
		err = c.BodyParser(&updatedTask)
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Невозможно разобрать тело запроса"})
		}
		tasks[updatedTaskIndex] = updatedTask
		return c.JSON(updatedTask)
	})

	app.Delete("/api/tasks/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Неверный формат ID"})
		}
		var deletedTaskIndex = -1
		for k, v := range tasks {
			if v.ID == id {
				deletedTaskIndex = k
				break
			}
		}
		if deletedTaskIndex == -1 {
			return c.Status(404).JSON(fiber.Map{"error": "Задача не найдена"})
		}
		tasks = append(tasks[:deletedTaskIndex], tasks[deletedTaskIndex+1:]...)
		return c.JSON(tasks)
	})

	app.Get("/api/tasks/:id", func(c *fiber.Ctx) error {
		id, err := strconv.Atoi(c.Params("id"))
		if err != nil {
			return c.Status(400).JSON(fiber.Map{"error": "Неверный формат ID"})
		}
		var foundedTask Task
		for _, v := range tasks {
			if v.ID == id {
				foundedTask = v
				break
			}
		}
		if foundedTask.ID == 0 {
			return c.Status(404).JSON(fiber.Map{"error": "задача с ID не найдена"})
		}
		return c.JSON(foundedTask)
	})

	err := app.Listen(":3000")
	if err != nil {
		panic(err)
	}
}
