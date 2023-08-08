package helpers

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/mjaliz/gotracktime/internal/constants"
	"github.com/mjaliz/gotracktime/internal/inputs"
	"os"
	"time"
)

var jwtSecretKey = []byte(os.Getenv("JWT_KEY"))

type JWTClaim struct {
	Username string `json:"username"`
	Email    string `json:"email"`
	jwt.RegisteredClaims
}

func GenerateJWT(u *inputs.User) (string, error) {
	expiredAt := time.Now().UTC().Add(constants.JWTExpireDuration)
	claims := &JWTClaim{
		Username: u.Name,
		Email:    u.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expiredAt),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(jwtSecretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
