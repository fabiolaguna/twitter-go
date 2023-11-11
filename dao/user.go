package dao

import (
	"context"

	"github.com/fabiolaguna/twitter-go/configurations"
	"github.com/fabiolaguna/twitter-go/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

func UserRegisteredCheck(email string) (models.User, bool, string) {
	db := configurations.MongoConnection.Database(configurations.DatabaseName)
	col := db.Collection("users")
	condition := bson.M{"email": email}

	var result models.User
	err := col.FindOne(context.TODO(), condition).Decode(&result)
	id := result.Id.Hex()

	if err != nil {
		return result, false, id
	}

	return result, true, id
}

func Insert(u models.User) (string, bool, error) {
	db := configurations.MongoConnection.Database(configurations.DatabaseName)
	col := db.Collection("users")

	u.Password, _ = encryptPassword(u.Password)
	result, err := col.InsertOne(context.TODO(), u)

	if err != nil {
		return "", false, err
	}

	objId, _ := result.InsertedID.(primitive.ObjectID)
	return objId.String(), true, nil
}

func encryptPassword(pass string) (string, error) {
	cost := 8
	bytes, err := bcrypt.GenerateFromPassword([]byte(pass), cost)

	if err != nil {
		return err.Error(), err
	}

	return string(bytes), nil
}
