package utils

import (
	"fmt"
	"time"

	"github.com/Gabriel-Macedogmc/report-logs-golang/internal/interfaces"
	"github.com/golang-jwt/jwt/v5"
)

type JwtLogin struct{}

var secretKey = []byte("33f34dad4bc0ce9dc320863509aed43cab33a93a29752779ae0df6dbbea33e56")

func NewJwtLogin() interfaces.JsonWebToken {
	return &JwtLogin{}
}

func (s *JwtLogin) GenerateToken(userID uint) (map[string]interface{}, error) {
	claims := jwt.MapClaims{}
	claims["user_id"] = userID
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)

	if err != nil {
		return nil, err
	}

	tokenMap := make(map[string]interface{})
	tokenMap["token"] = tokenString
	return tokenMap, nil
}

func (s *JwtLogin) VerifyToken(tokenString string) (bool, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil {
		return false, err
	}

	if !token.Valid {
		return false, fmt.Errorf("invalid token")
	}

	return true, nil
}
