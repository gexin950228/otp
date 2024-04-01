package middleware

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"otp/logSource"
	"strings"
	"time"
)

type Claims struct {
	ID       int
	UserName string `form:"username" json:"userName" gorm:"userName"`
	jwt.StandardClaims
}

var (
	secret     = []byte("jwt-otp")
	effectTime = time.Hour * 24
)

func GenerateToken(claims Claims) string {
	claims.ExpiresAt = time.Now().Add(effectTime).Unix()
	sign, err := jwt.NewWithClaims(jwt.SigningMethodES256, claims).SignedString(secret)
	if err != nil {
		logSource.Log.Error("初始化token失败")
	}
	return sign
}

func ParseToken() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		tokenString := ctx.GetHeader("user")
		if strings.Contains(ctx.Request.RequestURI, "user") {
			ctx.Abort()
			return
		}
		token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
			return secret, nil
		})
		if err != nil {
			logSource.Log.Error("token验证失败")
			ctx.Redirect(http.StatusMovedPermanently, "http://127.0.0.1:8080/?uri=http://localhost:8080/user/login/")
		}
		_, ok := token.Claims.(*Claims)
		if !ok {
			logSource.Log.Panic("token校验出错")
		}
		ctx.Next()
	}
}
