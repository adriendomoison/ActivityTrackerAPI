package api

import (
	"github.com/gin-gonic/gin"
)

var app *gin.Engine

func InitApi() {
	app = gin.Default()
	app.Run(":8100")
}
