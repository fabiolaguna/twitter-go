package services

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
	"github.com/fabiolaguna/twitter-go/dao"
	"github.com/fabiolaguna/twitter-go/models"
)

func GetProfile(request events.APIGatewayProxyRequest) models.Response {
	var response models.Response
	response.Status = 400

	id := request.QueryStringParameters["id"]
	if len(id) < 1 {
		response.Message = "Id param is required"
		fmt.Println("[profile service][method:GetProfile] " + response.Message)
		return response
	}

	profile, err := dao.GetProfile(id)
	if err != nil {
		response.Message = err.Error()
		fmt.Println("[profile service][method:GetProfile] Profile not found: " + response.Message)
		return response
	}

	jsonResponse, err := json.Marshal(profile)
	if err != nil {
		response.Status = 500
		response.Message = err.Error()
		fmt.Println("[profile service][method:GetProfile] Error formatting token: " + response.Message)
		return response
	}

	response.Status = 200
	response.Message = string(jsonResponse)
	return response
}

func UpdateProfile(ctx context.Context, claims models.Claim) models.Response {
	var response models.Response
	response.Status = 400

	var user models.User
	body := ctx.Value(models.Key("body")).(string)
	err := json.Unmarshal([]byte(body), &user)
	if err != nil {
		response.Message = "Incorrect body data: " + err.Error()
		fmt.Println("[profile service][method:UpdateProfile] " + response.Message)
		return response
	}

	status, err := dao.UpdateProfile(user, claims.Id.Hex())
	if err != nil {
		response.Message = "Error trying to update profile: " + err.Error()
		fmt.Println("[profile service][method:UpdateProfile] " + response.Message)
		return response
	}
	if !status {
		response.Message = "Profile couldn't be updated"
		fmt.Println("[profile service][method:UpdateProfile] " + response.Message)
		return response
	}

	response.Status = 200
	response.Message = "Profile updated successfully"
	return response
}
