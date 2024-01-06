package dao

import (
	"context"

	"github.com/fabiolaguna/twitter-go/configurations"
	"github.com/fabiolaguna/twitter-go/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func GetProfile(id string) (models.User, error) {
	ctx := context.TODO()
	db := configurations.MongoConnection.Database(configurations.DatabaseName)
	col := db.Collection("users")

	var profile models.User
	objectId, _ := primitive.ObjectIDFromHex(id)

	condition := bson.M{
		"_id": objectId,
	}

	err := col.FindOne(ctx, condition).Decode(&profile)
	profile.Password = ""
	if err != nil {
		return profile, err
	}

	return profile, nil
}

func UpdateProfile(user models.User, id string) (bool, error) {
	ctx := context.TODO()
	db := configurations.MongoConnection.Database(configurations.DatabaseName)
	col := db.Collection("users")

	record := make(map[string]interface{})
	if len(user.Name) > 0 {
		record["name"] = user.Name
	}
	if len(user.Surname) > 0 {
		record["surname"] = user.Surname
	}
	if len(user.Avatar) > 0 {
		record["avatar"] = user.Avatar
	}
	if len(user.Banner) > 0 {
		record["banner"] = user.Banner
	}
	if len(user.Biography) > 0 {
		record["biography"] = user.Biography
	}
	if len(user.Ubication) > 0 {
		record["ubication"] = user.Ubication
	}
	if len(user.Website) > 0 {
		record["website"] = user.Website
	}
	record["birthdate"] = user.Birthdate

	stringRecord := bson.M{
		"$set": record,
	}

	objectId, _ := primitive.ObjectIDFromHex(id)
	filter := bson.M{
		"_id": bson.M{
			"$eq": objectId,
		},
	}

	_, err := col.UpdateOne(ctx, filter, stringRecord)
	if err != nil {
		return false, err
	}

	return true, nil
}
