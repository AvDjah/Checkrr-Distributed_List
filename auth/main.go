package auth

import (
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"github.com/spf13/viper"
	"log"
	"time"
)

func GenerateJWT(userID int64, username string, role string) (string, error) {
	secretKey := viper.GetString("SECRET_KEY")

	if len(secretKey) == 0 {
		log.Panicln("Secret Not found")
	}
	myClaims := MyClaims{
		UserID:   userID,
		Username: username,
		Role:     role,
		Exp:      time.Now().Add(time.Hour * 1).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, myClaims)
	signedToken, err := token.SignedString([]byte(secretKey))
	if err != nil {
		return "", fmt.Errorf("failed to sign token: %w", err)
	}

	return signedToken, nil
}

func ParseJWT(tokenString string) (jwt.MapClaims, bool) {

	secretKey := viper.GetString("SECRET_KEY")

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(secretKey), nil
	})

	if err != nil {
		log.Panic(err)
		return nil, false
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		fmt.Println(claims["foo"], claims["nbf"])
		return claims, true
	} else {
		fmt.Println(err)
		return nil, false
	}
}
