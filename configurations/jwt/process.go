package jwt

import (
	"errors"
	"strings"

	"github.com/fabiolaguna/twitter-go/models"
	jwt "github.com/golang-jwt/jwt/v5"
)

var Email string
var UserId string

func TokenProccesing(token string, sign string) (*models.Claim, bool, string, error) {
	key := []byte(sign)
	var claims models.Claim

	splitToken := strings.Split(token, "Bearer")
	if len(splitToken) != 2 {
		return &claims, false, "", errors.New("Invalid token format")
	}

	token = strings.TrimSpace(splitToken[1])

	tokenParsed, err := jwt.ParseWithClaims(token, &claims, func(tk *jwt.Token) (interface{}, error) {
		return key, nil
	})
	if err == nil {
		// Chequeo contra DB
	}

	if !tokenParsed.Valid {
		return &claims, false, string(""), errors.New("Invalid token")
	}

	return &claims, true, string(""), err
}
