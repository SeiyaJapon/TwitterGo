package secretmanager

import (
	"encoding/json"
	"fmt"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/secretsmanager"

	"github.com/SeiyaJapon/golang/TwitterGo/awsgo"
	"github.com/SeiyaJapon/golang/TwitterGo/models"
)

func GetSecret(secretName string) (models.Secret, error) {
	var dataSecret models.Secret

	fmt.Println("> Request Secret" + secretName)

	svc := secretsmanager.NewFromConfig(awsgo.Cfg)

	key, err := svc.GetSecretValue(awsgo.Ctx, &secretsmanager.GetSecretValueInput{
		SecretId: aws.String(secretName),
	})

	if err != nil {
		fmt.Println(err.Error())

		return dataSecret, err
	}

	json.Unmarshal([]byte(*key.SecretString), &dataSecret)

	fmt.Println(" > Secret OK " + secretName)

	return dataSecret, nil
}
