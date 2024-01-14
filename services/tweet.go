package services

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/fabiolaguna/twitter-go/dao"
	"github.com/fabiolaguna/twitter-go/models"
)

func CreateTweet(ctx context.Context, claim models.Claim) models.Response {
	var message models.Tweet
	var response models.Response
	response.Status = 400

	userId := claim.Id.Hex()
	body := ctx.Value(models.Key("body")).(string)
	err := json.Unmarshal([]byte(body), &message)
	if err != nil {
		response.Message = err.Error()
		fmt.Println("[tweet service][method:CreateTweet] Error mapping body: " + response.Message)
		return response
	}

	register := models.CreatedTweet{
		UserId:  userId,
		Message: message.Message,
		Date:    time.Now(),
	}

	_, status, err := dao.CreateTweet(register)
	if err != nil {
		response.Status = 500
		response.Message = "[tweet service][method:CreateTweet] Error has occurred creating tweet: " + err.Error()
		fmt.Println(response.Message)
		return response
	}

	if !status {
		response.Status = 500
		response.Message = "The tweet could not be created"
		fmt.Println("[tweet service][method:CreateTweet] " + response.Message)
		return response
	}

	response.Status = 200
	response.Message = "Tweet created successfully"
	return response
}
