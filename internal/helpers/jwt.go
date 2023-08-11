package helpers

import (
	"errors"
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

func ValidateToken(signedToken string) error {
	token, err := jwt.ParseWithClaims(
		signedToken,
		&JWTClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(jwtSecretKey), nil
		},
	)
	if err != nil {
		return err
	}
	claims, ok := token.Claims.(*JWTClaim)
	if !ok {
		err = errors.New("couldn't parse claims")
	}
	if claims.ExpiresAt.Unix() < time.Now().UTC().Unix() {
		err = errors.New("token expired")
		return err
	}
	return nil
}
