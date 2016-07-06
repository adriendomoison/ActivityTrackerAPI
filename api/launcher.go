package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

var app *gin.Engine

func InitApi() {
	app = gin.Default()
	app.GET("/api/user/:id", func (c *gin.Context) {
		id, _ := strconv.Atoi(c.Params.ByName("id"))
		c.JSON(http.StatusOK, gin.H{"id": id, "working": true})
	})
	app.Run(":5000")
}
