package task

import (
	"context"

	"github.com/Daizaikun/Go-Fiber/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetAll(c *fiber.Ctx, coll *mongo.Collection) error {

	var tasks []models.Task

	result, err := coll.Find(context.TODO(), bson.M{})

	if err != nil {
		return c.JSON(&fiber.Map{
			"data": "Nulo",
		})
	}

	for result.Next(context.TODO()) {
		var task models.Task
		result.Decode(&task)
		tasks = append(tasks, task)
	}

	return c.JSON(&fiber.Map{
		"data": tasks,
	})

}
