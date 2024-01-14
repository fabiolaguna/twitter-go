package dao

import (
	"context"

	"github.com/fabiolaguna/twitter-go/configurations"
	"github.com/fabiolaguna/twitter-go/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func CreateTweet(tweet models.CreatedTweet) (string, bool, error) {
	ctx := context.TODO()
	db := configurations.MongoConnection.Database(configurations.DatabaseName)
	col := db.Collection("tweet")

	register := bson.M{
		"userid":  tweet.UserId,
		"message": tweet.Message,
		"date":    tweet.Date,
	}

	result, err := col.InsertOne(ctx, register)
	if err != nil {
		return "", false, err
	}

	objId, _ := result.InsertedID.(primitive.ObjectID)
	return objId.String(), true, nil
}
