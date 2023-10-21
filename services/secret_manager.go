package services

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"
	"github.com/fabiolaguna/twitter-go/configurations"
	"github.com/fabiolaguna/twitter-go/models"
)

func GetSecret(secretName string) (models.Secret, error) {
	var secretData models.Secret
	fmt.Println("Asking for secret: " + secretName)

	svc := secretsmanager.NewFromConfig(configurations.Config)
	key, err := svc.GetSecretValue(configurations.Ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretName),
	})

	if err != nil {
		fmt.Println("Error when getting secret: " + err.Error())
		return secretData, err
	}

	json.Unmarshal([]byte(*key.SecretString), &secretData)
	fmt.Println("Secret reading OK: " + secretName)

	return secretData, nil
}
