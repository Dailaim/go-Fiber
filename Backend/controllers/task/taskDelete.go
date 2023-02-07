package task

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func Delete(c *fiber.Ctx, coll *mongo.Collection) error {

	clave := c.Params("id")

	id, err := primitive.ObjectIDFromHex(clave)

	if err != nil {
		return c.JSON(&fiber.Map{
			"data": "Invalid id",
		})
	}

	filter := bson.D{{Key: "_id", Value: id}}

	result, err := coll.DeleteOne(context.TODO(), filter)

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
