package jwt

import (
	"context"
	"time"

	"github.com/fabiolaguna/twitter-go/models"
	jwt "github.com/golang-jwt/jwt/v5"
)

func Generate(ctx context.Context, user models.User) (string, error) {
	jwtSign := ctx.Value(models.Key("jwtSign")).(string)
	key := []byte(jwtSign)

	payload := jwt.MapClaims{
		"email":     user.Email,
		"name":      user.Name,
		"surname":   user.Surname,
		"birthdate": user.Birthdate,
		"biography": user.Biography,
		"ubication": user.Ubication,
		"avatar":    user.Avatar,
		"banner":    user.Banner,
		"website":   user.Website,
		"_id":       user.Id.Hex(),
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	finalToken, err := token.SignedString(key)

	if err != nil {
		return finalToken, err
	}

	return finalToken, nil
}
