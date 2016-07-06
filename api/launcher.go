package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"os"
	"log"
)

var app *gin.Engine

func InitApi() {
	port := os.Getenv("PORT")

	if port == "" {
		log.Println("Local deployement, use port 5000")
		port = "5000"
	} else {
		log.Println("Heroku deployement, use port " + port)
	}

	app = gin.Default()
	app.GET("/api/user/:id", func (c *gin.Context) {
		id, _ := strconv.Atoi(c.Params.ByName("id"))
		c.JSON(http.StatusOK, gin.H{"id": id, "working": true})
	})
	app.Run(":" + port)
}
