package routes

import (
	"github.com/SeiyaJapon/golang/TwitterGo/db"
	"github.com/SeiyaJapon/golang/TwitterGo/models"
	"github.com/aws/aws-lambda-go/events"
)

func DeleteTweet(request events.APIGatewayProxyRequest, claim models.Claim) models.RestApi {
	var response models.RestApi

	response.Status = 400

	ID := request.QueryStringParameters["id"]

	if len(ID) < 1 {
		response.Message = "ID required"

		return response
	}

	err := db.DeleteTweet(ID, claim.ID.Hex())

	if err != nil {
		response.Message = "Error deleting tweet: " + err.Error()

		return response
	}

	response.Message = "Success deleting tweet"
	response.Status = 200

	return response
}
