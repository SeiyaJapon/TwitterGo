package routes

import (
	"github.com/SeiyaJapon/golang/TwitterGo/db"
	"github.com/SeiyaJapon/golang/TwitterGo/models"
	"github.com/aws/aws-lambda-go/events"
)

func DeleteRelation(request events.APIGatewayProxyRequest, claim models.Claim) models.RestApi {
	var response models.RestApi

	response.Status = 400

	ID := request.QueryStringParameters["id"]

	if len(ID) < 1 {
		response.Message = "ID required"

		return response
	}

	var relation models.Relation

	relation.UserID = claim.ID.Hex()
	relation.UserRelationID = ID

	status, err := db.DeleteRelation(relation)

	if err != nil {
		response.Message = "Get error deleting relation: " + err.Error()

		return response
	}

	if !status {
		response.Message = "Get error deleting relation"

		return response
	}

	response.Status = 200
	response.Message = "Success deleting relation"

	return response
}
