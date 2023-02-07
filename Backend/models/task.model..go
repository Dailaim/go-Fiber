package models

import (
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Task struct {
	ID          primitive.ObjectID `bson:"_id,omitempty" validate:"required"`
	CreatedAt   time.Time          `bson:"created_at,omitempty" validate:"required"`
	UpdatedAt   time.Time          `bson:"updated_at,omitempty" validate:"required"`
	Title       string             `json:"title" bson:"title" validate:"required,unique"`
	Description string             `json:"description" bson:"description" validate:"required" `
	Done        bool               `bson:"done" json:"done"`
}

func NewTask(Data *Task) error {
	Data.ID = primitive.NewObjectID()
	Data.CreatedAt = time.Now()
	Data.UpdatedAt = time.Now()
	Data.Done = false

	return verification(Data)

}

func UpdateTask(Data *Task) error {

	Data.UpdatedAt = time.Now()

	return verification(Data)

}

func verification(Data *Task) error {
	if Data.Title == "" {
		return fmt.Errorf("title is required")
	}

	if Data.Description == "" {
		return fmt.Errorf("description is required")
	}

	return nil
}
