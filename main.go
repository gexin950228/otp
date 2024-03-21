package main

import (
	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"net/http"
	"otp/router"
	"otp/sessionInit"
	"time"
)

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
	s.ListenAndServe()
}
u