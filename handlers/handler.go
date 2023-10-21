package handlers

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/fabiolaguna/twitter-go/models"
)

func RequestHandlers(ctx context.Context, request events.APIGatewayProxyRequest) models.Response {
	fmt.Println("Proccesing [PATH: " + ctx.Value(models.Key("path")).(string) + " | HTTP METHOD: " + ctx.Value(models.Key("method")).(string) + "]")

	var response models.Response

	switch ctx.Value(models.Key("method")).(string) {
	case "POST":
		switch ctx.Value(models.Key("path")).(string) {

		}
		//
	case "GET":
		switch ctx.Value(models.Key("path")).(string) {

		}
		//
	case "PUT":
		switch ctx.Value(models.Key("path")).(string) {

		}
		//
	case "DELETE":
		switch ctx.Value(models.Key("path")).(string) {

		}
		//
	}

	response.Status = 404
	response.Message = "the requested url was not found"
	return response
}
