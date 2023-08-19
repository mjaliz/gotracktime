package token

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestJWTMaker(t *testing.T) {
	maker, err := NewJWTMaker("slkfjsdlifj3403dfsfsdfsdfdsfsdfsf9dfsfdsfsfsddfdfdfsfsfsdfsdf483rjeslkfj")
	require.NoError(t, err)

	email := "mrph@gmail.com"
	duration := time.Minute

	issuedAt := time.Now().UTC()
	expiresAt := issuedAt.Add(duration)

	token, err := maker.CreateToken(email, duration)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.NoError(t, err)
	require.NotEmpty(t, payload)

	require.NotZero(t, payload.ID)
	require.Equal(t, email, payload.Email)
	require.WithinDuration(t, issuedAt, payload.IssuedAt.Time, time.Second)
	require.WithinDuration(t, expiresAt, payload.ExpiresAt.Time, time.Second)

}

func TestExpiredJWTToken(t *testing.T) {
	maker, err := NewJWTMaker("slkfjsdlifj3403dfsfsdfsdfdsfsdfsf9dfsfdsfsfsddfdfdfsfsfsdfsdf483rjeslkfj")
	require.NoError(t, err)

	email := "mrph@gmail.com"
	token, err := maker.CreateToken(email, -time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.Error(t, err)
	require.ErrorContains(t, err, jwt.ErrTokenExpired.Error())
	require.Nil(t, payload)
}

func TestInvalidJWTTokenAlgNone(t *testing.T) {
	payload, err := NewPayload("mrph@gmail.com", time.Minute)
	require.NoError(t, err)

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodNone, payload)
	token, err := jwtToken.SignedString(jwt.UnsafeAllowNoneSignatureType)
	require.NoError(t, err)

	maker, err := NewJWTMaker("slkfjsdlifj3403dfsfsdfsdfdsfsdfsf9dfsfdsfsfsddfdfdfsfsfsdfsdf483rjeslkfj")
	payload, err = maker.VerifyToken(token)
	require.Error(t, err)
	require.ErrorContains(t, err, ErrInavalidToken.Error())
	require.Nil(t, payload)
}
