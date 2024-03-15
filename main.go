package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"otp/logSource"
	"otp/router"
	"otp/sessionInit"
	"time"
)

func Hello(ctx *gin.Context) {
	ctx.String(http.StatusOK, "Success!")
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/ping", Hello)
	return r
}

func main() {
	engine := gin.Default()
	router.Router(engine)
	s := &http.Server{
		Addr:              ":8080",
		Handler:           engine,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      5 * time.Second,
	}
	sessionStore := sessionInit.InitSession()
	engine.Use(sessions.Sessions("opt_session", sessionStore))
	err := s.ListenAndServe()
	if err != nil {
		logSource.Log.Panic(err.Error())
		return
	}
}
