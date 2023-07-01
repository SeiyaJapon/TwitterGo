package jwt

import (
	"context"
	"time"

	"github.com/SeiyaJapon/golang/TwitterGo/models"
	jwt "github.com/golang-jwt/jwt/v5"
)

func GenerateJWT(ctx context.Context, user models.User) (string, error) {
	jwtSign := ctx.Value(models.Key("jwtSign")).(string)
	myKey := []byte(jwtSign)

	payload := jwt.MapClaims{
		"email":    user.Email,
		"name":     user.Name,
		"surname":  user.Surname,
		"birthday": user.Birthday,
		"bio":      user.Bio,
		"location": user.Location,
		"website":  user.Website,
		"_id":      user.ID.Hex(),
		"expire":   time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(myKey)

	if err != nil {
		return tokenStr, err
	}

	return tokenStr, nil
}
