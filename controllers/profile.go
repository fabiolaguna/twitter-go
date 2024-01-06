package controllers

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/fabiolaguna/twitter-go/models"
	"github.com/fabiolaguna/twitter-go/services"
)

func GetProfile(request events.APIGatewayProxyRequest) models.Response {
	fmt.Println("Getting profile")
	return services.GetProfile(request)
}

func UpdateProfile(ctx context.Context, claims models.Claim) models.Response {
	fmt.Println("Updating profile")
	return services.UpdateProfile(ctx, claims)
}
