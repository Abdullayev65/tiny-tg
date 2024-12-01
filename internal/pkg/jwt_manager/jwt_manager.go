package jwt_manager

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"strconv"
	"strings"
	"time"
)

type JwtManager struct {
	signingKey     []byte
	expiryInterval time.Duration
}

func New(signingKey string, expiryInterval time.Duration) *JwtManager {
	return &JwtManager{signingKey: []byte(signingKey), expiryInterval: expiryInterval}
}

func (t *JwtManager) Generate(id int) (string, error) {
	return t.generate(strconv.Itoa(id))
}

func (t *JwtManager) Parse(tokenStr string) (int, error) {
	tokenStr = removeBearerIfExists(tokenStr)

	subject, err := t.parse(tokenStr)
	if err != nil {
		return 0, err
	}

	return strconv.Atoi(subject)
}

func (t *JwtManager) generate(sub string) (string, error) {
	now := time.Now()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &jwt.StandardClaims{
		IssuedAt:  now.Unix(),
		ExpiresAt: now.Add(t.expiryInterval).Unix(),
		Subject:   sub,
	})
	return token.SignedString(t.signingKey)
}

func (t *JwtManager) parse(tokenStr string) (string, error) {
	token, err := jwt.ParseWithClaims(tokenStr, &jwt.StandardClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return t.signingKey, nil
	})
	if err != nil {
		return "", err
	}

	claims, ok := token.Claims.(*jwt.StandardClaims)
	if !ok {
		return "", errors.New("token claims are not type of  jwt.StandardClaims")
	}
	return claims.Subject, nil
}

func removeBearerIfExists(token string) string {
	arr := strings.Split(token, " ")
	if len(arr) < 2 {
		return token
	}
	if len(arr) == 2 && strings.EqualFold("Bearer", arr[0]) {
		return arr[1]
	}
	return token
}
