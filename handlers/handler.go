package handlers

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/fabiolaguna/twitter-go/configurations/jwt"
	"github.com/fabiolaguna/twitter-go/controllers"
	"github.com/fabiolaguna/twitter-go/models"
)

func RequestHandlers(ctx context.Context, request events.APIGatewayProxyRequest) models.Response {
	fmt.Println("Proccesing [PATH: " + ctx.Value(models.Key("path")).(string) + " | HTTP METHOD: " + ctx.Value(models.Key("method")).(string) + "]")

	var response models.Response

	isValid, statusCode, msg, claim := authorization(ctx, request)

	if !isValid {
		response.Status = statusCode
		response.Message = msg
		return response
	}

	switch ctx.Value(models.Key("method")).(string) {
	case "POST":
		switch ctx.Value(models.Key("path")).(string) {
		case "register":
			return controllers.Create(ctx)
		case "login":
			return controllers.Login(ctx)
		case "tweet":
			return controllers.CreateTweet(ctx, claim)
		case "avatar":
			return controllers.UploadImage(ctx, "avatar", request, claim)
		case "banner":
			return controllers.UploadImage(ctx, "banner", request, claim)
		}
	case "GET":
		switch ctx.Value(models.Key("path")).(string) {
		case "profile":
			return controllers.GetProfile(request)
		case "tweets":
			return controllers.GetTweets(request)
		}
	case "PUT":
		switch ctx.Value(models.Key("path")).(string) {
		case "profile":
			return controllers.UpdateProfile(ctx, claim)
		}
		//
	case "DELETE":
		switch ctx.Value(models.Key("path")).(string) {
		case "tweet":
			return controllers.DeleteTweet(request, claim)
		}
	}

	response.Status = 404
	response.Message = "the requested url was not found"
	return response
}

func authorization(ctx context.Context, request events.APIGatewayProxyRequest) (bool, int, string, models.Claim) {
	path := ctx.Value(models.Key("path")).(string)

	if path == "register" || path == "login" {
		return true, 200, "", models.Claim{}
	}

	token := request.Headers["Authorization"]

	if len(token) == 0 {
		return false, 401, "Token is required", models.Claim{}
	}

	claim, isOk, msg, err := jwt.TokenProccesing(token, ctx.Value(models.Key("jwtSign")).(string))
	if !isOk {
		if err != nil {
			fmt.Println("Error in token " + err.Error())
			return false, 401, err.Error(), models.Claim{}
		} else {
			fmt.Println("Error in token " + msg)
			return false, 401, msg, models.Claim{}
		}
	}

	fmt.Println("Token is valid")
	return true, 200, msg, *claim
}
