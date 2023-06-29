package routes

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/SeiyaJapon/golang/TwitterGo/db"
	"github.com/SeiyaJapon/golang/TwitterGo/models"
)

func Register(ctx context.Context) models.RestApi {
	var user models.User
	var response models.RestApi

	response.Status = 400

	fmt.Println("Go into Register")

	body := ctx.Value(models.Key("body")).(string)
	err := json.Unmarshal([]byte(body), &user)

	if err != nil {
		response.Message = err.Error()

		fmt.Println(response.Message)

		return response
	}

	if 0 == len(user.Email) {
		response.Message = "Not email empty"

		fmt.Println(response.Message)

		return response
	}

	if 0 == len(user.Password) {
		response.Message = "Password must not empty and more than 6 characters"

		fmt.Println(response.Message)

		return response
	}

	_, found, _ := db.ExistUser(user.Email)

	if found {
		response.Message = "Already exists an user with this email"

		fmt.Println(response.Message)

		return response
	}

	_, status, err := db.StoreRegister(user)

	if err != nil {
		response.Message = "Error registering user: " + err.Error()

		fmt.Println(response.Message)

		return response
	}

	if !status {
		response.Message = "Error registering user"

		fmt.Println(response.Message)

		return response
	}

	response.Status = 200
	response.Message = "Register success"

	fmt.Println(response.Message)

	return response
}
