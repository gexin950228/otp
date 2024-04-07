package middleware

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"math/rand"
	"time"
)

var (
	effectTime = time.Hour * 24
	secretKey  = []byte("qazcdewrsfagqtaaptwtagqtagsrqtag")
)

type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func GenerateToken(username string) string {
	fmt.Println(secretKey)
	_, err := rand.Read(secretKey)
	if err != nil {
		fmt.Printf("生成secretKey出错， %s\n", err.Error())
	}
	fmt.Printf("Generate Token Username: %s\n", username)
	claims := MyClaims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(effectTime).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(secretKey)
	if err != nil {
		fmt.Printf("Signed Error: %v\n", err)
	}
	return signedToken
}

func ParseToken(tokenString string) (*MyClaims, error) {
	tokenClaims, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*MyClaims); ok && tokenClaims.Valid {
			return claims, nil
		}
	}
	return nil, err
}
