package configurations

import (
	"context"
	"fmt"

	"github.com/fabiolaguna/twitter-go/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoConnection *mongo.Client
var DatabaseName string

func ConnectDB(ctx context.Context) error {
	user := ctx.Value(models.Key("user")).(string)
	password := ctx.Value(models.Key("password")).(string)
	host := ctx.Value(models.Key("host")).(string)

	connectionString := fmt.Sprintf("mongodb+srv://%s:%s@%s/?retryWrites=true&w=majority", user, password, host)
	var clientOptions = options.Client().ApplyURI(connectionString)

	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		fmt.Println("Error connecting database: " + err.Error())
		return err
	}

	err = client.Ping(ctx, nil)
	if err != nil {
		fmt.Println("Error getting ping from database connection: " + err.Error())
		return err
	}

	fmt.Println("Successful connection to database")
	MongoConnection = client
	DatabaseName = ctx.Value(models.Key("database")).(string)

	return nil
}

func DBConnected() bool {
	err := MongoConnection.Ping(context.TODO(), nil)
	return err == nil
}
