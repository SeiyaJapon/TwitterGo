package routes

import (
	"context"
	"encoding/json"
	"github.com/SeiyaJapon/golang/TwitterGo/db"
	"github.com/SeiyaJapon/golang/TwitterGo/models"
)

func UpdateProfile(ctx context.Context, claim models.Claim) models.RestApi {
	var response models.RestApi

	response.Status = 400

	var user models.User

	body := ctx.Value(models.Key("body")).(string)
	err := json.Unmarshal([]byte(body), &user)

	if err != nil {
		response.Message = "Incorrect Data: " + err.Error()
	}

	status, err := db.UpdateRegister(user, claim.ID)

	if err != nil || !status {
		response.Message = "Get error updating user: " + err.Error()

		return response
	}

	response.Status = 200
	response.Message = "Success updating profile"

	return response
}
