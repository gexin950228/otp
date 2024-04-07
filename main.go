package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"otp/router"
	"otp/sessionInit"
	"time"
)

func Cors(ctx *gin.Context) {
	ctx.Header("Access-Control-Allow-Origin", "*")
	ctx.Next()
}

func main() {
	engine := gin.Default()
	engine.LoadHTMLGlob("templates/**/*")
	engine.Static("/static", "static")
	router.Router(engine)
	s := &http.Server{
		Addr:              ":8080",
		Handler:           engine,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      5 * time.Second,
	}
	sessionStore := sessionInit.InitSession()
	engine.Use(sessions.Sessions("opt_session", sessionStore))
	engine.Use(Cors)
	s.ListenAndServe()
}
