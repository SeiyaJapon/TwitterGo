package jwt

import (
	"errors"
	"github.com/SeiyaJapon/golang/TwitterGo/db"
	"strings"

	jwt "github.com/golang-jwt/jwt/v5"

	"github.com/SeiyaJapon/golang/TwitterGo/models"
)

var Email string
var IDUsuario string

func ProcessToken(tk string, JWTSign string) (*models.Claim, bool, string, error) {
	myKey := []byte(JWTSign)
	var claims models.Claim

	splitToken := strings.Split(tk, "Bearer")

	if len(splitToken) != 2 {
		return &claims, false, string(""), errors.New("Invalid format token")
	}

	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, &claims, func(token *jwt.Token) (interface{}, error) {
		return myKey, nil
	})

	if err == nil {
		// check DB
		_, found, _ := db.ExistUser(claims.Email)

		if found {
			Email = claims.Email
			IDUsuario = claims.ID.Hex()
		}

		return &claims, found, IDUsuario, nil
	}

	if !tkn.Valid {
		return &claims, false, string(""), errors.New("Token Invalid")
	}

	return &claims, false, string(""), err
}
