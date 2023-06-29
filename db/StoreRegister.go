package db

import (
	"context"
	"github.com/SeiyaJapon/golang/TwitterGo/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func StoreRegister(user models.User) (string, bool, error) {
	ctx := context.TODO()

	db := MongoCN.Database(Database)
	respCollection := db.Collection("users")

	user.Password, _ = EncryptPassword(user.Password)

	result, err := respCollection.InsertOne(ctx, user)

	if err != nil {
		return "", false, err
	}

	ObjID, _ := result.InsertedID.(primitive.ObjectID)

	return ObjID.String(), true, nil
}
