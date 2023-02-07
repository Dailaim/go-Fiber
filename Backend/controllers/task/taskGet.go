package task

import (
	"context"

	"github.com/Daizaikun/Go-Fiber/models"
	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func Get(c *fiber.Ctx, coll *mongo.Collection) error {

	var task models.Task

	clave := c.Params("id")

	id, err := primitive.ObjectIDFromHex(clave)

	if err != nil {
		return c.JSON(&fiber.Map{
			"data": "Invalid id",
		})
	}

	filter := bson.D{{Key: "_id", Value: id}}

	err = coll.FindOne(context.TODO(), filter).Decode(&task)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return c.JSON(&fiber.Map{
				"data": "No documents found",
			})
		}
	}

	return c.JSON(&fiber.Map{
		"data": task,
	})

}
