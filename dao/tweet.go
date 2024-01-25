package dao

import (
	"context"

	"github.com/fabiolaguna/twitter-go/configurations"
	"github.com/fabiolaguna/twitter-go/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
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

func GetTweets(id string, page int64) ([]*models.TweetsResponse, bool) {
	ctx := context.TODO()
	db := configurations.MongoConnection.Database(configurations.DatabaseName)
	col := db.Collection("tweet")
	var results []*models.TweetsResponse

	condition := bson.M{
		"userid": id,
	}
	options := options.Find()
	options.SetLimit(20)
	options.SetSort(bson.D{{Key: "fecha", Value: -1}})
	options.SetSkip((page - 1) * 20)

	registries, err := col.Find(ctx, condition, options)
	if err != nil {
		return results, false
	}

	for registries.Next(ctx) {
		var registry models.TweetsResponse
		err := registries.Decode(&registry)
		if err != nil {
			return results, false
		}
		results = append(results, &registry)
	}

	return results, true
}

func DeleteTweet(id string, userId string) error {
	ctx := context.TODO()
	db := configurations.MongoConnection.Database(configurations.DatabaseName)
	col := db.Collection("tweet")

	objectId, _ := primitive.ObjectIDFromHex(id)
	condition := bson.M{
		"_id":    objectId,
		"userid": userId,
	}

	_, err := col.DeleteOne(ctx, condition)
	return err
}
