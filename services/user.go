package services

import (
	"context"
	"encoding/json"
	"fmt"

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
		fmt.Println("Error mapping body: " + response.Message)
		return response
	}

	if len(user.Email) == 0 {
		response.Message = "Email is required"
		fmt.Println(response.Message)
		return response
	}

	if len(user.Password) < 6 {
		response.Message = "Password with more than 5 characters is required"
		fmt.Println(response.Message)
		return response
	}

	_, userExists, _ := dao.UserRegisteredCheck(user.Email)
	if userExists {
		response.Message = "The email sent is already registered"
		fmt.Println(response.Message)
		return response
	}

	_, status, err := dao.Insert(user)
	if err != nil {
		response.Status = 500
		response.Message = "Error has occurred creating user: " + err.Error()
		fmt.Println(response.Message)
		return response
	}

	if !status {
		response.Status = 500
		response.Message = "The user could not be created"
		fmt.Println(response.Message)
		return response
	}

	response.Status = 200
	response.Message = "User created"
	fmt.Println(response.Message)
	return response
}
