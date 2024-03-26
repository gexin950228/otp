package user

import (
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func TokenHandler(userName string, w http.ResponseWriter, r *http.Request) {
	secretKey := []byte("hayvcbiyhyya")
	token := jwt.New(jwt.SigningMethodES256)
	claim := token.Claims.(jwt.MapClaims)
	claim["username"] = userName
	claim["exp"] = time.Now().Add(time.Hour * 24).Unix()
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		fmt.Println(err)
		return
	}
	w.Header().Set("Authorization", "Bearer"+tokenString)
	w.WriteHeader(http.StatusOK)
}

func CheckToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {

	}
}
