package configurations

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/config"
)

const awsRegion = "us-east-1"

var Ctx context.Context
var Config aws.Config
var err error

func AwsInitializer() {
	Ctx = context.TODO()
	Config, err = config.LoadDefaultConfig(Ctx, config.WithDefaultRegion(awsRegion))

	if err != nil {
		panic("Error initializing AWS configuration: " + err.Error())
	}
}
