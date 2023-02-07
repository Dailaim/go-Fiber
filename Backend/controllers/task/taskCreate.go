package task

import (
	"context"

	"github.com/Daizaikun/Go-Fiber/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/mongo"
)

func Create(c *fiber.Ctx, coll *mongo.Collection) error {

	var task models.Task

	err := c.BodyParser(&task)

	if err != nil {
		return c.SendString(err.Error())
	}

	err = models.NewTask(&task)

	if err != nil {
		return c.SendString(err.Error())
	}

	result, err := coll.InsertOne(context.TODO(), task)

	if err != nil {
		return c.JSON(&fiber.Map{
			"data": "Nulo",
		})
	}

	return c.JSON(&fiber.Map{
		"data": result,
	})

}
