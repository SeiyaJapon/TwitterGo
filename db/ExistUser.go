package db

import (
	"context"
	"github.com/SeiyaJapon/golang/TwitterGo/models"
	"go.mongodb.org/mongo-driver/bson"
)

func ExistUser(email string) (models.User, bool, string) {
	ctx := context.TODO()

	db := MongoCN.Database(Database)
	respCollection := db.Collection("users")

	condition := bson.M{"email": email}

	var result models.User

	err := respCollection.FindOne(ctx, condition).Decode(&result)
	ID := result.ID.Hex()

	if err != nil {
		return result, false, ID
	}

	return result, true, ID
}
