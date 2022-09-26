package main

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go/v4"
)

var mySigningKey = []byte("mysupersecrete")

func GenerateJWT() (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)

	claims["authorized"] = true
	claims["user"] = "Smith Golang"
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()

	tokenString, err := token.SignedString(mySigningKey)

	if err != nil {
		err = fmt.Errorf("something went wrong %s", err.Error())
		return "", err
	}
	return tokenString, nil
}

func main() {
	fmt.Println("-------Testing JWT-------")

	tokenString, err := GenerateJWT()
	if err != nil {
		fmt.Println("Error generating token string")
	}

	fmt.Println(tokenString)
}
