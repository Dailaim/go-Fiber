package router

import (
	"github.com/Daizaikun/Go-Fiber/controllers/task"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func Task(api fiber.Router, Database *mongo.Database) {

	coll := Database.Collection("task")

	api.Post("/", func(c *fiber.Ctx) error {
		return task.Create(c, coll)
	})

	api.Get("/", func(c *fiber.Ctx) error {
		return task.GetAll(c, coll)
	})

	api.Put("/:id", func(c *fiber.Ctx) error {
		return task.Update(c, coll)
	})

	api.Delete("/:id", func(c *fiber.Ctx) error {
		return task.Delete(c, coll)
	})

	api.Get("/:id", func(c *fiber.Ctx) error {
		return task.Get(c, coll)
	})

}
