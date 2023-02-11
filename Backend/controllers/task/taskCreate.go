package task

import (
	"context"
	"fmt"

	"github.com/Daizaikun/Go-Fiber/models"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"

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

	file, err := c.FormFile("document")

	if file != nil && err == nil {

		c.SaveFile(file, fmt.Sprintf("./uploads/%s", uuid.New().String()))

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
