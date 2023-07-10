package routes

import (
	"bytes"
	"context"
	"encoding/base64"
	"io"
	"mime"
	"mime/multipart"
	"strings"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/s3/s3manager"

	"github.com/SeiyaJapon/golang/TwitterGo/db"
	"github.com/SeiyaJapon/golang/TwitterGo/models"
)

type readSeeker struct {
	io.Reader
}

func (readSeeker *readSeeker) Seek(offset int64, whence int) (int64, error) {
	return 0, nil
}

func UploadImage(ctx context.Context, uploadType string, request events.APIGatewayProxyRequest, claim models.Claim) models.RestApi {
	var response models.RestApi

	response.Status = 400

	IDUser := claim.ID.Hex()
	bucket := aws.String(ctx.Value(models.Key("bucketName")).(string))

	var filename string
	var user models.User

	switch uploadType {
	case "A":
		filename = "avatars/" + IDUser + ".jpg"
		user.Avatar = filename
	case "B":
		filename = "banners/" + IDUser + ".jpg"
		user.Banner = filename
	}

	mediaType, params, err := mime.ParseMediaType(request.Headers["content-type"])

	if err != nil {
		response.Status = 500
		response.Message = err.Error()

		return response
	}

	if strings.HasPrefix(mediaType, "multipart/") {
		body, err := base64.StdEncoding.DecodeString(request.Body)

		if err != nil {
			response.Status = 500
			response.Message = err.Error()

			return response
		}

		multipartRead := multipart.NewReader(bytes.NewReader(body), params["boundary"])
		part, err := multipartRead.NextPart()

		if err != nil && err != io.EOF {
			response.Status = 500
			response.Message = err.Error()

			return response
		}

		if err != io.EOF {
			if part.FileName() != "" {
				buffer := bytes.NewBuffer(nil)

				if _, err := io.Copy(buffer, part); err != nil {
					response.Status = 500
					response.Message = err.Error()

					return response
				}

				myNewSession, err := session.NewSession(&aws.Config{
					Region: aws.String("eu-west-1"),
				})

				if err != nil {
					response.Status = 500
					response.Message = err.Error()

					return response
				}

				uploader := s3manager.NewUploader(myNewSession)
				_, err = uploader.Upload(&s3manager.UploadInput{
					Bucket: bucket,
					Key:    aws.String(filename),
					Body:   &readSeeker{buffer},
				})

				if err != nil {
					response.Status = 500
					response.Message = err.Error()

					return response
				}
			}
		}

		status, err := db.UpdateRegister(user, IDUser)

		if err != nil || !status {
			response.Status = 400
			response.Message = "Error updating user: " + err.Error()

			return response
		}
	} else {
		response.Message = "Request is not an image"
		response.Status = 400

		return response
	}

	response.Status = 200
	response.Message = "Success uploading image"

	return response
}
