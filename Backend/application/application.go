package application

import (
	"github.com/Daizaikun/Go-Fiber/router"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"

	"go.mongodb.org/mongo-driver/mongo"
	"github.com/gofiber/fiber/v2/middleware/cors"
	
)

func App(Database *mongo.Database) *fiber.App {

	app := fiber.New()

	app.Use(cors.New())

	app.Use(logger.New(logger.Config{
		Format: "[${ip}]:${port} ${status} - ${method} ${path}\n",
	}))

	app.Route("/task", func(api fiber.Router) {
		router.Task(api, Database)
	})

	return app

}
