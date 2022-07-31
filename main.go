package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func setupRouter() *gin.Engine {
	gin.SetMode(gin.ReleaseMode)
	
	r := gin.Default()
	r.LoadHTMLGlob("templates/*.html")

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		os.Getenv("username"): os.Getenv("password"),
	}))
	authorized.GET("/", welcome_page)
	authorized.Static("/static", "./static")

	return r
}

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	fmt.Println(os.Getenv("welcomemessage"))

	r := setupRouter()
	r.Run("0.0.0.0:80")
}

func welcome_page(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}
