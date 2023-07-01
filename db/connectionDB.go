package db

import (
	"context"
	"fmt"

	"github.com/SeiyaJapon/golang/TwitterGo/models"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var MongoCN *mongo.Client
var Database string

func ConnectDB(ctx context.Context) error {
	username := ctx.Value(models.Key("username")).(string)
	password := ctx.Value(models.Key("password")).(string)
	host := ctx.Value(models.Key("host")).(string)
	connStr := fmt.Sprintf("mongodb+srv://%s:%s@%s/?retryWrites=true&w=majority", username, password, host)

	var clientOptions = options.Client().ApplyURI(connStr)

	client, err := mongo.Connect(ctx, clientOptions)

	if err != nil {
		fmt.Println(err.Error())

		return err
	}

	err = client.Ping(ctx, nil)

	if err != nil {
		fmt.Println(err.Error())

		return err
	}

	fmt.Println("DB Connection Success")

	MongoCN = client
	Database = ctx.Value(models.Key("database")).(string)

	return nil
}

func DBConnected() bool {
	err := MongoCN.Ping(context.TODO(), nil)

	return err == nil
}
