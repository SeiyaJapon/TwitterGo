package routes

import (
	"context"
	"encoding/json"
	"github.com/SeiyaJapon/golang/TwitterGo/db"
	"github.com/SeiyaJapon/golang/TwitterGo/models"
	"time"
)

func AddTweet(ctx context.Context, claim models.Claim) models.RestApi {
	var message models.Tweet
	var response models.RestApi

	response.Status = 400

	IDUser := claim.ID.Hex()
	body := ctx.Value(models.Key("body")).(string)
	err := json.Unmarshal([]byte(body), &message)

	if err != nil {
		response.Message = "Get an error decoding body: " + err.Error()

		return response
	}

	register := models.StoreTweet{
		UserID:  IDUser,
		Message: message.Message,
		Date:    time.Now(),
	}

	_, status, err := db.InsertTweet(register)

	if err != nil {
		response.Message = "Get an error inserting register: " + err.Error()

		return response
	}

	if !status {
		response.Message = "Unable to insert register"

		return response
	}

	response.Status = 200
	response.Message = "Success creating Tweet!"

	return response
}
