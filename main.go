package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/go-sql-driver/mysql"
)

func Greet(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"MESSAGE": "GO RELEASER DEMO PARA MI AMORCITO :D"})
}

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {

		c.Header("Access-Control-Allow-Origin", "*")
		c.Header("Access-Control-Allow-Credentials", "true")
		c.Header("Access-Control-Allow-Headers", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization, accept, origin, Cache-Control, X-Requested-With")
		c.Header("Access-Control-Allow-Methods", "POST, HEAD, PATCH, OPTIONS, GET, PUT")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}

func main() {

	// Creating Gin server
	r := gin.Default()
	r.Use(CORSMiddleware())

	// Setting middlewares
	r.Use(gin.Logger())
	r.Use(gin.Recovery())

	// Routes
	r.GET("/", Greet)

	//Starting server on port 3000
	r.Run(":3000")
}
