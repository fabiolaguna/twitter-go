package controllers

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/fabiolaguna/twitter-go/models"
	"github.com/fabiolaguna/twitter-go/services"
)

func CreateTweet(ctx context.Context, claim models.Claim) models.Response {
	fmt.Println("Creating tweet")
	return services.CreateTweet(ctx, claim)
}

func GetTweets(request events.APIGatewayProxyRequest) models.Response {
	fmt.Println("Getting tweets")
	return services.GetTweets(request)
}

func DeleteTweet(request events.APIGatewayProxyRequest, claim models.Claim) models.Response {
	fmt.Println("Deleting tweet")
	return services.DeleteTweet(request, claim)
}
