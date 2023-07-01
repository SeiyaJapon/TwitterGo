package routes

import (
	"context"
	"encoding/json"
	"github.com/SeiyaJapon/golang/TwitterGo/db"
	"github.com/SeiyaJapon/golang/TwitterGo/jwt"
	"github.com/SeiyaJapon/golang/TwitterGo/models"
)

func Login(ctx context.Context) models.RestApi {
	var user models.User
	var response models.RestApi

	response.Status = 400

	body := ctx.Value(models.Key("body")).(string)
	err := json.Unmarshal([]byte(body), &user)

	if err != nil {
		response.Message = "User and/or password incorrect: " + err.Error()

		return response
	}

	if len(user.Email) == 0 {
		response.Message = "Email required"

		return response
	}

	userData, exists := db.tryLogin(user.Email, user.Password)

	if !exists {
		response.Message = "User and/or password incorrect: " + err.Error()

		return response
	}

	jwtKey, err := jwt.GenerateJWT(ctx, user)

	if err != nil {
		response.Message = "Error generating token: " + err.Error()

		return response
	}

}
