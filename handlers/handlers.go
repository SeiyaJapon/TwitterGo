package handlers

import (
	"context"
	"fmt"
	"github.com/SeiyaJapon/golang/TwitterGo/jwt"

	"github.com/aws/aws-lambda-go/events"

	"github.com/SeiyaJapon/golang/TwitterGo/models"
	"github.com/SeiyaJapon/golang/TwitterGo/routes"
)

func Handlers(ctx context.Context, request events.APIGatewayProxyRequest) models.RestApi {
	fmt.Println("Processing " + ctx.Value(models.Key("path")).(string) + " > " + ctx.Value(models.Key("method")).(string))

	var response models.RestApi
	response.Status = 400

	isSuccess, statusCode, msg, claim := validateAuthorization(ctx, request)

	if !isSuccess {
		response.Status = statusCode
		response.Message = msg

		return response
	}

	switch ctx.Value(models.Key("method")).(string) {
	case "POST":
		switch ctx.Value(models.Key("path")).(string) {
		case "register":
			return routes.Register(ctx)
		case "login":
			return routes.Login(ctx)
		case "tweet":
			return routes.AddTweet(ctx, claim)
		case "uploadAvatar":
			return routes.UploadImage(ctx, "A", request, claim)
		case "uploadBanner":
			return routes.UploadImage(ctx, "B", request, claim)
		}
	case "GET":
		switch ctx.Value(models.Key("path")).(string) {
		case "profile":
			return routes.Profile(request)
		case "getTweets":
			return routes.GetTweets(request)
		}
	case "PUT":
		switch ctx.Value(models.Key("path")).(string) {
		case "updateProfile":
			return routes.UpdateProfile(ctx, claim)
		}
	case "DELETE":
		switch ctx.Value(models.Key("path")).(string) {
		case "deleteTweet":
			return routes.DeleteTweet(request, claim)
		}
	}

	response.Message = "Method Invalid"

	return response
}

func validateAuthorization(ctx context.Context, request events.APIGatewayProxyRequest) (bool, int, string, models.Claim) {
	path := ctx.Value(models.Key("path")).(string)

	if path == "register" || path == "login" || path == "getAvatar" || path == "getBanner" {
		return true, 200, "", models.Claim{}
	}

	token := request.Headers["Authorization"]

	if len(token) == 0 {
		return false, 401, "Required token", models.Claim{}
	}

	claim, allOK, msg, err := jwt.ProcessToken(token, ctx.Value(models.Key("jwtSign")).(string))

	if !allOK {
		if err != nil {
			fmt.Println("Token error: " + err.Error())

			return false, 401, err.Error(), models.Claim{}
		} else {
			fmt.Println("Token error: " + msg)

			return false, 401, msg, models.Claim{}
		}
	}

	fmt.Println("Token OK")

	return true, 200, msg, *claim
}
