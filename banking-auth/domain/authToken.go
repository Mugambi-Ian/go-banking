package domain

import (
	"go-banking/banking-auth/utils"

	"github.com/dgrijalva/jwt-go"
)

type AuthToken struct {
	token *jwt.Token
}

func (t AuthToken) NewAccessToken() (string, *utils.AppError) {
	signedString, err := t.token.SignedString([]byte(HMAC_SAMPLE_SECRET))
	if err != nil {
		utils.LogError("Failed while signing access token: " + err.Error())
		return "", utils.NewUnexpectedError("cannot generate access token")
	}
	return signedString, nil
}

func (t AuthToken) newRefreshToken() (string, *utils.AppError) {
	c := t.token.Claims.(AccessTokenClaims)
	refreshClaims := c.RefreshTokenClaims()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, refreshClaims)
	signedString, err := token.SignedString([]byte(HMAC_SAMPLE_SECRET))
	if err != nil {
		utils.LogError("Failed while signing refresh token: " + err.Error())
		return "", utils.NewUnexpectedError("cannot generate refresh token")
	}
	return signedString, nil
}

func NewAuthToken(claims AccessTokenClaims) AuthToken {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return AuthToken{token: token}
}

func NewAccessTokenFromRefreshToken(refreshToken string) (string, *utils.AppError) {
	token, err := jwt.ParseWithClaims(refreshToken, &RefreshTokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(HMAC_SAMPLE_SECRET), nil
	})
	if err != nil {
		return "", utils.NewAuthenticationError("invalid or expired refresh token")
	}
	r := token.Claims.(*RefreshTokenClaims)
	accessTokenClaims := r.AccessTokenClaims()
	authToken := NewAuthToken(accessTokenClaims)

	return authToken.NewAccessToken()
}
