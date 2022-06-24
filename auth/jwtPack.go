package auth

import (
	"FinalProjectAssignment/config"
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/ilyakaznacheev/cleanenv"
)

var cfg config.Config

func GenerateJWT(user_id string) (string, error) {
	_ = cleanenv.ReadConfig(".env", &cfg)
	var JwtKey = []byte(cfg.JwtKey)
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = user_id
	claims["expired"] = time.Now().Add(time.Hour * 12).Unix()

	tokenString, err := token.SignedString(JwtKey)

	if err != nil {
		fmt.Errorf("Something Went Wrong: %s", err.Error())
		return "", err
	}
	return tokenString, nil
}

func ValidationToken(createdToken string) (*jwt.Token, error) {
	_ = cleanenv.ReadConfig(".env", &cfg)
	var JwtKey = []byte(cfg.JwtKey)
	token, err := jwt.Parse(createdToken, func(t *jwt.Token) (interface{}, error) {
		_, ok := t.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("Token Anuthorized")
		}
		return JwtKey, nil
	})
	if err != nil {
		return token, err
	}
	return token, nil
}
