package domain

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/marcelomoresco/go-sales/internal/application/constants"
	"github.com/marcelomoresco/go-sales/internal/configuration/errors"
)

func (ud *UserDomain) GenerateToken() (string, *errors.KasError) {
	secret := os.Getenv(constants.JWT_SECRET_KEY)

	claims := jwt.MapClaims{
		"id":    ud.Id,
		"email": ud.Email,
		"name":  ud.Name,
		"exp":   time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", errors.NewInternalServerError(
			fmt.Sprintf("error trying to generate jwt token, err=%s", err.Error()))
	}

	return tokenString, nil
}