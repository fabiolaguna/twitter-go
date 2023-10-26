package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id        primitive.ObjectID `bson:"_id,omitempty" json:"id"`
	Name      string             `bson:"name" json:"name,omitempty"`
	Surname   string             `bson:"surname" json:"surname,omitempty"`
	Birthdate time.Time          `bson:"birthdate" json:"birthdate,omitempty"`
	Email     string             `bson:"email" json:"email"`
	Password  string             `bson:"password" json:"password,omitempty"`
	Avatar    string             `bson:"avatar" json:"avatar,omitempty"`
	Banner    string             `bson:"banner" json:"banner,omitempty"`
	Biography string             `bson:"biography" json:"biography,omitempty"`
	Ubication string             `bson:"ubication" json:"ubication,omitempty"`
	Website   string             `bson:"website" json:"website,omitempty"`
}
