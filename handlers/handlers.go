package handlers

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"

	"github.com/SeiyaJapon/golang/TwitterGo/models"
)

func Handlers(ctx context.Context, request events.APIGatewayProxyRequest) models.RestApi {
	fmt.Println("Processing " + ctx.Value(models.Key("path")).(string) + " > " + ctx.Value(models.Key("method")).(string))

	var response models.RestApi
	response.Status = 400

	switch ctx.Value(models.Key("method")).(string) {
	case "POST":
		switch ctx.Value(models.Key("method")).(string) {

		}
		//
	case "GET":
		switch ctx.Value(models.Key("method")).(string) {

		}
		//
	case "PUT":
		switch ctx.Value(models.Key("method")).(string) {

		}
		//
	case "DELETE":
		switch ctx.Value(models.Key("method")).(string) {

		}
		//
	}

	response.Message = "Method Invalid"

	return response
}
