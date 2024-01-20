package tokenManager

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

var ErrInvalidToken = errors.New("invalid token")

type TokenManager struct {
	accessTokenExpiredDuration  time.Duration
	accessTokenSecret           string
	refreshTokenExpiredDuration time.Duration
	refreshTokenSecret          string
}

type Claims struct {
	Username string `json:"username"`
	jwt.RegisteredClaims
}

func NewTokenManager() *TokenManager {
	accessTokenExpDuration := os.Getenv("ACCESS_TOKEN_EXPIRED_DURATION")
	durationInt, err := strconv.Atoi(accessTokenExpDuration)
	if err != nil {
		fmt.Println(err)
	}
	accessTokenTime := time.Duration(durationInt)
	accessTokenSecret := os.Getenv("ACCESS_TOKEN_SECRET")

	refreshTokenExpiredDuration := os.Getenv("REFRESH_TOKEN_EXPIRED_DURATION")
	refreshDurationInt, err := strconv.Atoi(refreshTokenExpiredDuration)
	if err != nil {
		fmt.Println(err)
	}
	refreshTokenTime := time.Duration(refreshDurationInt)
	refreshTokenSecret := os.Getenv("REFRESH_TOKEN_SECRET")

	fmt.Printf("%s, %s, %s, %s", accessTokenTime, refreshTokenTime, accessTokenSecret, refreshTokenSecret)

	return &TokenManager{
		accessTokenExpiredDuration:  accessTokenTime,
		refreshTokenExpiredDuration: refreshTokenTime,
		accessTokenSecret:           accessTokenSecret,
		refreshTokenSecret:          refreshTokenSecret,
	}
}

func (s *TokenManager) GenerateAccessToken(username string) (string, error) {
	claims := Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(s.accessTokenExpiredDuration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(s.accessTokenSecret))
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(claims)

	return signedToken, err

}

func (s *TokenManager) GenerateRefreshToken(username string) (string, error) {
	claims := &Claims{
		Username: username,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(s.refreshTokenExpiredDuration)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(s.refreshTokenSecret))

	return signedToken, err
}

func (s *TokenManager) Parse(accessToken string) (string, error) {
	token, err := jwt.Parse(accessToken, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %s", t.Header["alg"])
		}
		return []byte(s.accessTokenSecret), nil
	})
	if err != nil {
		return "", fmt.Errorf("cannot parse token")
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", fmt.Errorf("cannot get claims from token")
	}
	return claims["username"].(string), nil
}
