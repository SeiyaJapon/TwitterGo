package main

import (
	"context"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"

	"github.com/SeiyaJapon/golang/TwitterGo/awsgo"
	"github.com/SeiyaJapon/golang/TwitterGo/db"
	"github.com/SeiyaJapon/golang/TwitterGo/handlers"
	"github.com/SeiyaJapon/golang/TwitterGo/models"
	"github.com/SeiyaJapon/golang/TwitterGo/secretmanager"

	"os"
)

func main() {
	lambda.Start(execLambda)
}

func execLambda(theContext context.Context, request events.APIGatewayProxyRequest) (*events.APIGatewayProxyResponse, error) {
	var resp *events.APIGatewayProxyResponse

	awsgo.InitAWS()

	if !ValidateParams() {
		resp = &events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Error en las variables de entorno. Deben incluir 'SecretName', 'BucketName', 'UrlPrefix'",
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}

		return resp, nil
	}

	SecretModel, err := secretmanager.GetSecret(os.Getenv("SecretName"))

	if err != nil {
		resp = &events.APIGatewayProxyResponse{
			StatusCode: 400,
			Body:       "Error en la lectura de Secret " + err.Error(),
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}

		return resp, nil
	}

	path := strings.Replace(request.PathParameters["twittergo"], os.Getenv("UrlPrefix"), "", -1)

	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("path"), path)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("method"), request.HTTPMethod)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("user"), SecretModel.Username)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("password"), SecretModel.Password)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("host"), SecretModel.Host)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("database"), SecretModel.Database)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("jwtsign"), SecretModel.JWTSign)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("body"), request.Body)
	awsgo.Ctx = context.WithValue(awsgo.Ctx, models.Key("bucktName"), os.Getenv("BucketName"))

	err = db.ConnectDB(awsgo.Ctx)

	if err != nil {
		resp = &events.APIGatewayProxyResponse{
			StatusCode: 500,
			Body:       "Error al conectar la DB " + err.Error(),
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}

		return resp, nil
	}

	restAPI := handlers.Handlers(awsgo.Ctx, request)

	if restAPI.CustomResp == nil {
		resp = &events.APIGatewayProxyResponse{
			StatusCode: restAPI.Status,
			Body:       restAPI.Message,
			Headers: map[string]string{
				"Content-Type": "application/json",
			},
		}

		return resp, nil
	} else {
		return restAPI.CustomResp, nil
	}
}

func ValidateParams() bool {
	_, getParam := os.LookupEnv("SecretName")
	if !getParam {
		return getParam
	}

	_, getParam = os.LookupEnv("BucketName")
	if !getParam {
		return getParam
	}

	_, getParam = os.LookupEnv("UrlPrefix")
	if !getParam {
		return getParam
	}

	return getParam
}
