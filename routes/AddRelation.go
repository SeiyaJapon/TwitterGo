package routes

import (
	"context"
	"github.com/SeiyaJapon/golang/TwitterGo/db"
	"github.com/SeiyaJapon/golang/TwitterGo/models"
	"github.com/aws/aws-lambda-go/events"
)

func AddRelation(ctx context.Context, request events.APIGatewayProxyRequest, claim models.Claim) models.RestApi {
	var response models.RestApi

	response.Status = 400

	ID := request.QueryStringParameters["id"]

	if len(ID) < 1 {
		response.Message = "ID required"

		return response
	}

	var user models.Relation

	user.UserID = claim.ID.Hex()
	user.UserRelationID = ID

	status, err := db.InsertRelation(user)

	if err != nil {
		response.Message = "Get error when inserting relation: " + err.Error()

		return response
	}

	if !status {
		response.Message = "Get error when inserting relation"

		return response
	}

	response.Status = 200
	response.Message = "Success adding relation"

	return response
}
