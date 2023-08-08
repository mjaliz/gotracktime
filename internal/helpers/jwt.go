package helpers

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/mjaliz/gotracktime/internal/models"
	"os"
	"time"
)

var jwtSecretKey = []byte(os.Getenv("JWT_KEY"))

type JWTClaim struct {
	UserID uint   `json:"userid"`
	Email  string `json:"email"`
	jwt.RegisteredClaims
}

func GenerateJWT(u *models.User, expiredAt time.Time) (string, error) {
	claims := &JWTClaim{
		UserID: u.ID,
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
