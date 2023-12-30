package services

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/aws/aws-lambda-go/events"
	"github.com/fabiolaguna/twitter-go/configurations/jwt"
	"github.com/fabiolaguna/twitter-go/dao"
	"github.com/fabiolaguna/twitter-go/models"
)

func Create(ctx context.Context) models.Response {
	var user models.User
	var response models.Response
	response.Status = 400

	body := ctx.Value(models.Key("body")).(string)
	err := json.Unmarshal([]byte(body), &user)
	if err != nil {
		response.Message = err.Error()
		fmt.Println("[user service][method:Create] Error mapping body: " + response.Message)
		return response
	}

	if len(user.Email) == 0 {
		response.Message = "Email is required"
		fmt.Println("[user service][method:Create] " + response.Message)
		return response
	}

	if len(user.Password) < 6 {
		response.Message = "Password with more than 5 characters is required"
		fmt.Println("[user service][method:Create] " + response.Message)
		return response
	}

	_, userExists, _ := dao.UserRegisteredCheck(user.Email)
	if userExists {
		response.Message = "The email sent is already registered"
		fmt.Println("[user service][method:Create] " + response.Message)
		return response
	}

	_, status, err := dao.Insert(user)
	if err != nil {
		response.Status = 500
		response.Message = "[user service][method:Create] Error has occurred creating user: " + err.Error()
		fmt.Println(response.Message)
		return response
	}

	if !status {
		response.Status = 500
		response.Message = "The user could not be created"
		fmt.Println("[user service][method:Create] " + response.Message)
		return response
	}

	response.Status = 200
	response.Message = "User created"
	fmt.Println("[user service][method:Create] " + response.Message)
	return response
}

func Login(ctx context.Context) models.Response {
	var user models.User
	var response models.Response
	response.Status = 400

	body := ctx.Value(models.Key("body")).(string)
	err := json.Unmarshal([]byte(body), &user)
	if err != nil {
		response.Message = err.Error()
		fmt.Println("[user service][method:Login] Error mapping body: " + response.Message)
		return response
	}

	if len(user.Email) == 0 {
		response.Message = "Email is required"
		fmt.Println("[user service][method:Login] " + response.Message)
		return response
	}

	userData, exists := dao.Login(user.Email, user.Password)
	if !exists {
		response.Message = "Incorrect email and/or password"
		fmt.Println("[user service][method:Login] " + response.Message)
		return response
	}

	jwt, err := jwt.Generate(ctx, userData)
	if err != nil {
		response.Message = err.Error()
		fmt.Println("[user service][method:Login] Error generating token: " + response.Message)
		return response
	}

	loginResponse := models.LoginResponse{
		Token: jwt,
	}

	token, err := json.Marshal(loginResponse)
	if err != nil {
		response.Message = err.Error()
		fmt.Println("[user service][method:Login] Error formatting token: " + response.Message)
		return response
	}

	cookie := &http.Cookie{
		Name:    "token",
		Value:   jwt,
		Expires: time.Now().Add(24 * time.Hour),
	}
	finalCookie := cookie.String()

	proxyResponse := &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(token),
		Headers: map[string]string{
			"Content-Type":                "application/json",
			"Access-Control-Allow-Origin": "*",
			"Set-Cookie":                  finalCookie,
		},
	}

	response.Status = 200
	response.Message = string(token)
	response.CustomResponse = proxyResponse

	return response
}

func Profile(request events.APIGatewayProxyRequest) models.Response {
	var response models.Response
	response.Status = 400

	id := request.QueryStringParameters["id"]
	if len(id) < 1 {
		response.Message = "Id param is required"
		fmt.Println("[user service][method:Profile] " + response.Message)
		return response
	}

	profile, err := dao.Profile(id)
	if err != nil {
		response.Message = err.Error()
		fmt.Println("[user service][method:Profile] Profile not found: " + response.Message)
		return response
	}

	jsonResponse, err := json.Marshal(profile)
	if err != nil {
		response.Status = 500
		response.Message = err.Error()
		fmt.Println("[user service][method:Profile] Error formatting token: " + response.Message)
		return response
	}

	response.Status = 200
	response.Message = string(jsonResponse)
	return response
}
