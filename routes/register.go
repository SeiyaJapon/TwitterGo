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
	fmt.Println("1")
	err := json.Unmarshal([]byte(body), &user)
	fmt.Println("2")

	if err != nil {
		fmt.Println("3")
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

	fmt.Println("4")
	_, found, _ := db.ExistUser(user.Email)
	fmt.Println("5")

	if found {
		response.Message = "Already exists an user with this email"

		fmt.Println(response.Message)

		return response
	}

	fmt.Println("6")
	_, status, err := db.StoreRegister(user)
	fmt.Println("7")

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
