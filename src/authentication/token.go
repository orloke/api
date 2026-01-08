package authentication

import (
	"os"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

func CreateToken(userID uint64, email string, nick string) (string, error) {
	permissions := jwt.MapClaims{}
	permissions["authorized"] = true
	permissions["exp"] = time.Now().Add(time.Hour * 6).Unix()
	permissions["userId"] = userID
	permissions["email"] = email
	permissions["nick"] = nick

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, permissions)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}
