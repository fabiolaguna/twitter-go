package controllers

import (
	"context"
	"fmt"

	"github.com/fabiolaguna/twitter-go/models"
	"github.com/fabiolaguna/twitter-go/services"
)

func CreateTweet(ctx context.Context, claim models.Claim) models.Response {
	fmt.Println("Creating tweet")
	return services.CreateTweet(ctx, claim)
}
