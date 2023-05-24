package token

import (
	"testing"
	"time"

	"github.com/golang-jwt/jwt"
	"github.com/julysNICK/stock_system/utils"
	"github.com/stretchr/testify/require"
)

func TestJwtMaker(t *testing.T) {
	make, err := NewJwtMaker(utils.RandomString(32))

	require.NoError(t, err)

	userName := utils.RandomName()

	duration := time.Minute

	issueAt := time.Now()

	expiresAt := issueAt.Add(duration)

	token, payload, err := make.CreateToken(userName, duration)
	require.NoError(t, err)

	require.NotEmpty(t, token)

	require.Equal(t, userName, payload.Username)

	require.Equal(t, payload.IssuedAt.Unix(), issueAt.Unix())

	require.Equal(t, payload.ExpiredAt.Unix(), expiresAt.Unix())

}

func TestExpiredJwtMaker(t *testing.T) {
	make, err := NewJwtMaker(utils.RandomString(32))

	require.NoError(t, err)

	token, _, err := make.CreateToken(utils.RandomName(), -time.Minute)

	require.NoError(t, err)

	_, err = make.VerifyToken(token)

	require.Error(t, err)
}

func TestInvalidJWTTokenAlgNone(t *testing.T) {
	payload, err := NewPayload(utils.RandomName(), time.Minute)
	require.NoError(t, err)

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodNone, payload)
	token, err := jwtToken.SignedString(jwt.UnsafeAllowNoneSignatureType)
	require.NoError(t, err)

	maker, err := NewJwtMaker(utils.RandomString(32))
	require.NoError(t, err)

	payload, err = maker.VerifyToken(token)
	require.Error(t, err)
	require.EqualError(t, err, ErrInvalidToken.Error())
	require.Nil(t, payload)
}
