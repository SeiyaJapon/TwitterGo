package routes

import (
	"context"
	"encoding/json"
	"github.com/SeiyaJapon/golang/TwitterGo/db"
	"github.com/SeiyaJapon/golang/TwitterGo/jwt"
	"github.com/SeiyaJapon/golang/TwitterGo/models"
	"github.com/aws/aws-lambda-go/events"
	"net/http"
	"time"
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

	userData, exists := db.TryLogin(user.Email, user.Password)

	if !exists {
		response.Message = "User and/or password incorrect: " + err.Error()

		return response
	}

	jwtKey, err := jwt.GenerateJWT(ctx, userData)

	if err != nil {
		response.Message = "Error generating token: " + err.Error()

		return response
	}

	loginResponse := models.LoginResponse{
		Token: jwtKey,
	}

	token, errResponseLogin := json.Marshal(loginResponse)

	if errResponseLogin != nil {
		response.Message = "Error formating token: " + errResponseLogin.Error()

		return response
	}

	cookie := &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: time.Now().Add(24 * time.Hour),
	}
	cookieString := cookie.String()

	customResponse := &events.APIGatewayProxyResponse{
		StatusCode: 200,
		Body:       string(token),
		Headers: map[string]string{
			"Content-Type":                "application/json",
			"Access-Control-Allow-Origin": "*",
			"Set-Cookie":                  cookieString,
		},
	}

	response.Status = 200
	response.Message = string(token)
	response.CustomResp = customResponse

	return response
}
