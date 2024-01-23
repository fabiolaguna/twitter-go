package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type TweetsResponse struct {
	Id      primitive.ObjectID `bson:"_id" json:"_id,omitempty"`
	UserId  string             `bson:"userid" json:"userid,omitempty"`
	Message string             `bson:"message" json:"message,omitempty"`
	Date    time.Time          `bson:"date" json:"date,omitempty"`
}
