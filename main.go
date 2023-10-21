package main

import (
	"context"
	"os"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	lambda "github.com/aws/aws-lambda-go/lambda"
	"github.com/fabiolaguna/twitter-go/configurations"
	"github.com/fabiolaguna/twitter-go/handlers"
	"github.com/fabiolaguna/twitter-go/models"
	"github.com/fabiolaguna/twitter-go/services"
)

func main() {
	lambda.Start(LambdaExecute)
}

func LambdaExecute(ctx context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	var res *events.APIGatewayProxyResponse

	configurations.AwsInitializer()

	if !validateParams() {
		res = &events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Environment variables error. SecretName, BucketName and UrlPrefix not found",
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}

		return res, nil
	}

	SecretModel, err := services.GetSecret(os.Getenv("SecretName"))

	if err != nil {
		res = &events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Error reading secret: " + err.Error(),
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}

		return res, nil
	}

	path := strings.Replace(request.PathParameters["twittergo"], os.Getenv("UrlPrefix"), "", -1)

	configurations.Ctx = context.WithValue(configurations.Ctx, models.Key("path"), path)
	configurations.Ctx = context.WithValue(configurations.Ctx, models.Key("method"), request.HTTPMethod)
	configurations.Ctx = context.WithValue(configurations.Ctx, models.Key("user"), SecretModel.Username)
	configurations.Ctx = context.WithValue(configurations.Ctx, models.Key("password"), SecretModel.Password)
	configurations.Ctx = context.WithValue(configurations.Ctx, models.Key("host"), SecretModel.Host)
	configurations.Ctx = context.WithValue(configurations.Ctx, models.Key("database"), SecretModel.Database)
	configurations.Ctx = context.WithValue(configurations.Ctx, models.Key("jwtSign"), SecretModel.JWTSign)
	configurations.Ctx = context.WithValue(configurations.Ctx, models.Key("body"), request.Body)
	configurations.Ctx = context.WithValue(configurations.Ctx, models.Key("bucketName"), os.Getenv("BucketName"))

	// Database connection
	configurations.ConnectDB(configurations.Ctx)
	if err != nil {
		res = &events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Error connecting database: " + err.Error(),
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}

		return res, nil
	}

	response := handlers.RequestHandlers(configurations.Ctx, request)
	if response.CustomResponse == nil {
		res = &events.APIGatewayProxyResponse{
			StatusCode: response.Status,
			Body:       response.Message,
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}

		return res, nil
	} else {
		return response.CustomResponse, nil
	}
}

func validateParams() bool {
	_, existsSecretName := os.LookupEnv("SecretName")
	_, existsBucketName := os.LookupEnv("BucketName")
	_, existsUrlPrefix := os.LookupEnv("UrlPrefix")

	if !existsSecretName || !existsBucketName || !existsUrlPrefix {
		return false
	}

	return true
}
