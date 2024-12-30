package pkg

import (
	"fmt"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var secretKey = []byte(os.Getenv("SECRET_KEY"))

type CustomClaimsJwt struct {
	Id string `json:"id"`
	jwt.RegisteredClaims
}

func JwtGenerateAccessToken(id string) (string, error) {
	claims := CustomClaimsJwt{
		id,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add((time.Hour * 24) * 7)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err := token.SignedString(secretKey)
	return accessToken, err
}
func JwtVerify(tokenString string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomClaimsJwt{}, func(t *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return "", fmt.Errorf("cannot parse token: %s", err)
	}
	claims, ok := token.Claims.(*CustomClaimsJwt)
	if !ok {
		return "", fmt.Errorf("invalid token")
	}
	if claims.ExpiresAt.Time.Before(time.Now()) {
		return "", fmt.Errorf("token is expired")
	}
	id := claims.Id
	return id, nil
}
