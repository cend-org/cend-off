package token

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

const SecretKey = "Secret_key"

type Token struct {
	UserId         int `json:"user_id,omitempty"`
	AccessLevel    int `json:"access_level,omitempty"`
	ExpirationDate struct {
		Value time.Time `json:"value,omitempty"`
	} `json:"expiration_date,omitempty"`
}

func New(token Token) (string, error) {
	tokenJson, err := json.Marshal(token)
	if err != nil {
		return "", err
	}

	claims := &jwt.MapClaims{
		"iss":  "issuer",
		"exp":  time.Now().Add(time.Hour * 24).Unix(),
		"data": string(tokenJson),
	}

	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return accessToken.SignedString([]byte(SecretKey))
}

func Refresh(claims jwt.Claims) (string, error) {
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return refreshToken.SignedString([]byte(SecretKey))
}

func Parse(accessToken string) *Token {
	var tokenValue Token

	parsedAccessToken, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		return []byte(SecretKey), nil
	})

	if err != nil {
		return nil
	}

	mapClaims := parsedAccessToken.Claims.(jwt.MapClaims)

	data := mapClaims["data"].(string)

	err = json.Unmarshal([]byte(data), &tokenValue)
	if err != nil {
		return nil
	}

	return &tokenValue
}

func GetFromContext(ctx context.Context) (*Token, error) {
	unparsedToken := ctx.Value("token")

	if tok, ok := unparsedToken.(Token); ok {
		return &tok, nil
	}

	return &Token{}, errors.New("cannot reach token data")
}
