package authentication

import (
	"context"
	"duval/internal/configuration"
	"duval/pkg/database"
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strings"
)

type Auth struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Token struct {
	UserId     uint `json:"user_id"`
	UserLevel  uint `json:"user_level"`
	UserStatus uint `json:"user_status"`
	jwt.RegisteredClaims
}

var userCtxKey = &contextKey{name: "Authorization"}

type contextKey struct {
	name string
}

func GetTokenString(userId uint) (str string, err error) {
	var tok Token
	err = database.Get(&tok, `SELECT u.id as 'user_id', u.status as 'user_status' FROM user u WHERE u.id = ?`, userId)
	if err != nil {
		return str, err
	}

	err = database.Get(&tok, `SELECT auth.level as 'user_level' FROM authorization auth WHERE auth.user_id = ?`, userId)
	if err != nil {
		return str, err
	}

	str, err = NewAccessToken(tok)
	if err != nil {
		return str, err
	}

	return str, err
}

func GetTokenDataFromContext(ctx context.Context) (tok *Token, err error) {
	tokenString, ok := ctx.Value(userCtxKey).(string)

	if !ok || len(strings.TrimSpace(tokenString)) == 0 {
		return nil, errors.New("bad header value given")
	}

	bearer := strings.Split(tokenString, " ")
	if len(bearer) != 2 || bearer[0] != "Bearer" {
		return nil, errors.New("incorrectly formatted authorization header")
	}

	tok, err = ParseAccessToken(bearer[1])
	if err != nil {
		return nil, err
	}

	return tok, nil
}

/*

	MIDDLEWARE

*/

func AuthMiddleware() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token := r.Header.Get("Authorization")
			// put it in context
			ctx := context.WithValue(r.Context(), userCtxKey, token)
			// and call the next with our new context
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

/*

	UTILS

*/

func NewAccessToken(claims Token) (string, error) {
	accessToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return accessToken.SignedString([]byte(configuration.App.TokenSecret))
}

func ParseAccessToken(accessToken string) (*Token, error) {
	parsedAccessToken, err := jwt.ParseWithClaims(accessToken, &Token{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(configuration.App.TokenSecret), nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := parsedAccessToken.Claims.(*Token); ok && parsedAccessToken.Valid {
		return claims, nil
	} else {
		return nil, errors.New("invalid token")
	}
}

func ParseRefreshToken(refreshToken string) *jwt.RegisteredClaims {
	parsedRefreshToken, _ := jwt.ParseWithClaims(refreshToken, &jwt.RegisteredClaims{}, func(token *jwt.Token) (any, error) {
		return []byte(configuration.App.TokenSecret), nil
	})

	return parsedRefreshToken.Claims.(*jwt.RegisteredClaims)
}

func NewRefreshToken(claims jwt.RegisteredClaims) (string, error) {
	refreshToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return refreshToken.SignedString([]byte(configuration.App.TokenSecret))
}
