package task

import (
	"context"

	"github.com/Daizaikun/Go-Fiber/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func Update(c *fiber.Ctx, coll *mongo.Collection) error {

	clave := c.Params("id")

	id, err := primitive.ObjectIDFromHex(clave)

	if err != nil {
		return c.JSON(&fiber.Map{
			"data": "No documents found",
		})
	}

	var task models.Task

	err = c.BodyParser(&task)

	if err != nil {
		return c.SendString(err.Error())
	}

	err = models.UpdateTask(&task)

	if err != nil {
		return c.SendString(err.Error())
	}

	filter := bson.D{{Key: "_id", Value: id}}
	update := bson.D{{Key: "$set", Value: task}}

	result, err := coll.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.JSON(&fiber.Map{
				"data": "No documents found",
			})
		}
	}

	return c.JSON(&fiber.Map{
		"data": result,
	})
}
