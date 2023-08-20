package token

import (
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestPasetoMaker(t *testing.T) {
	maker, err := NewPasetoMaker("slkfjsdlifj3403dfsfsdfsdfdsfsdfs")
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

func TestExpiredPasetoToken(t *testing.T) {
	maker, err := NewPasetoMaker("slkfjsdlifj3403dfsfsdfsdfdsfsdfs")
	require.NoError(t, err)

	email := "mrph@gmail.com"
	token, err := maker.CreateToken(email, -time.Minute)
	require.NoError(t, err)
	require.NotEmpty(t, token)

	payload, err := maker.VerifyToken(token)
	require.Error(t, err)
	require.ErrorContains(t, err, ErrExpiredToken.Error())
	require.Nil(t, payload)
}
