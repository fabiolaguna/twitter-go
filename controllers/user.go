package controllers

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/fabiolaguna/twitter-go/models"
	"github.com/fabiolaguna/twitter-go/services"
)

func Create(ctx context.Context) models.Response {
	fmt.Println("Creating user")
	return services.Create(ctx)
}

func Login(ctx context.Context) models.Response {
	fmt.Println("User is loging")
	return services.Login(ctx)
}

func Profile(request events.APIGatewayProxyRequest) models.Response {
	fmt.Println("Getting profile")
	return services.Profile(request)
}
