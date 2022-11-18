package helpers

import (
	"dibagi/config"
	"errors"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
)

func GenerateToken(id string, username string) (string, error) {
	claims := jwt.MapClaims{
		"id":        id,
		"user_name": username,
		"exp":       time.Now().Add(time.Hour * 24).Unix(),
	}

	parseToken := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := parseToken.SignedString([]byte(config.SECRET_KEY))
	if err != nil {
		return "", err
	}

	return signedToken, nil
}

func VerifyToken(ctx *gin.Context) (interface{}, error) {
	headerToken := GetRequestHeaders(ctx).Authorization
	if headerToken == "" {
		return nil, errors.New("token is empty")
	}

	haveBearer := strings.HasPrefix(headerToken, "Bearer")
	if !haveBearer {
		return nil, errors.New("no bearer token found")
	}

	stringToken := strings.Split(headerToken, " ")[1]

	token, err := jwt.Parse(stringToken, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("failed to parse token, expected jwt.SigningMethodHMAC")
		}
		return []byte(config.SECRET_KEY), nil
	})

	if err != nil {
		return nil, errors.New("failed to parse token, token invalid or expired")
	}

	v, ok := token.Claims.(jwt.MapClaims)

	if !ok || !token.Valid {
		return nil, errors.New("failed to claims or token not valid")
	}

	return v, nil
}
