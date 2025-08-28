package authentication

import (
	"os"
	"time"

	"github.com/golang-jwt/jwt"
)

func GenerateJWTtoken(email string, userID int64) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"email":  email,
		"userID": userID,
		"exp":    time.Now().Add(time.Hour * 2).Unix(),
	})

	return token.SignedString([]byte(os.Getenv("JWTSECRETKEY")))
}
