package routes

import (
	"encoding/json"
	"fmt"
	"github.com/SeiyaJapon/golang/TwitterGo/db"

	"github.com/SeiyaJapon/golang/TwitterGo/models"
	"github.com/aws/aws-lambda-go/events"
)

func Profile(request events.APIGatewayProxyRequest) models.RestApi {
	var response models.RestApi

	response.Status = 200

	fmt.Println("Go into Profile")

	ID := request.QueryStringParameters["id"]

	if len(ID) < 1 {
		response.Message = "ID required"

		return response
	}

	profile, err := db.FindProfileById(ID)

	if err != nil {
		response.Message = "Error finding user: " + err.Error()

		return response
	}

	respJson, err := json.Marshal(profile)

	if err != nil {
		response.Status = 500
		response.Message = "Error formatting data as JSON: " + err.Error()

		return response
	}

	response.Message = string(respJson)

	return response
}
