package models

import "github.com/aws/aws-lambda-go/events"

type Response struct {
	Status         int
	Message        string
	CustomResponse *events.APIGatewayProxyResponse
}
