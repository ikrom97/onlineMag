package token

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"log"
	"onlineMag/models"
	"time"
)

type MyCustomClaims struct {
	ID    int64
	Name  string
	Phone int64
	Email string
	Role  string
	Login string
	jwt.StandardClaims
}

func CreateToken(user models.User) string {
	mySigningKey := []byte("AllYourBase")
	claims := MyCustomClaims{
		ID:    user.ID,
		Name:  user.Name,
		Phone: user.Phone,
		Email: user.Email,
		Role:  user.Role,
		Login: user.Login,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Unix() + 3600,
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString(mySigningKey)
	if err != nil {
		log.Println(err)

	}
	fmt.Printf("I am a token = %v\n", ss)
	return ss
}

func ParseToken(tokenString string) *MyCustomClaims {
	token, _ := jwt.ParseWithClaims(tokenString, &MyCustomClaims{}, func(token *jwt.Token) (interface{}, error) {
		return nil, nil
	})
	claims := token.Claims.(*MyCustomClaims)
	return claims
}
