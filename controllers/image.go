package controllers

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/fabiolaguna/twitter-go/models"
	"github.com/fabiolaguna/twitter-go/services"
)

func UploadImage(ctx context.Context, uploadType string, request events.APIGatewayProxyRequest, claim models.Claim) models.Response {
	fmt.Println("Uploading " + uploadType + " image")
	return services.UploadImage(ctx, uploadType, request, claim)
}
