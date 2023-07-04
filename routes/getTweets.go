package routes

import (
	"encoding/json"
	"github.com/SeiyaJapon/golang/TwitterGo/db"
	"github.com/SeiyaJapon/golang/TwitterGo/models"
	"github.com/aws/aws-lambda-go/events"
	"strconv"
)

func GetTweets(request events.APIGatewayProxyRequest) models.RestApi {
	var response models.RestApi

	response.Status = 400

	ID := request.QueryStringParameters["id"]
	page := request.QueryStringParameters["page"]

	if len(ID) < 1 {
		response.Message = "ID required"

		return response
	}

	if len(page) < 1 {
		page = "1"
	}

	pag, err := strconv.Atoi(page)

	if err != nil {
		response.Message = "Parameter 'page' must be int greater than 0"

		return response
	}

	tweets, success := db.GetTweets(ID, int64(pag))

	if !success {
		response.Message = "Error getting tweets"

		return response
	}

	jsonResponse, err := json.Marshal(tweets)

	if err != nil {
		response.Status = 500
		response.Message = "Get error at formating data to JSON"

		return response
	}

	response.Status = 200
	response.Message = string(jsonResponse)

	return response
}
