package db

import (
	"context"
	"github.com/SeiyaJapon/golang/TwitterGo/models"
)

func DeleteRelation(relation models.Relation) (bool, error) {
	ctx := context.TODO()
	db := MongoCN.Database(Database)
	relationCollection := db.Collection("relation")

	_, err := relationCollection.DeleteOne(ctx, relation)

	if err != nil {
		return false, err
	}

	return true, nil
}
