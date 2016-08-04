package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"log"
	"github.com/adriendomoison/ActivityTrackerAPI/DTO"
	"github.com/adriendomoison/ActivityTrackerAPI/service"
	"github.com/adriendomoison/ActivityTrackerAPI/translate"
	"github.com/nicksnyder/go-i18n/i18n"
	"strconv"
	"github.com/adriendomoison/ActivityTrackerAPI/utils"
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
	
	initRouts()
	
	app.Run(":" + port)
}

func initRouts() {
	app.POST("/api/users", notImplemented)
	app.DELETE("/api/users/:userId", notImplemented)
	app.POST("/api/users/:userId/messages", notImplemented)
	app.POST("/api/events", createEvent)
	app.GET("/api/events/:eventId", notImplemented)
	app.PUT("/api/events/:eventId", notImplemented)
	app.DELETE("/api/events/:eventId", notImplemented)
	app.POST("/api/events/:eventId/subscribe", notImplemented)
	app.DELETE("/api/events/:eventId/subscribe", notImplemented)
	app.POST("/api/programs", createProgram)
	app.GET("/api/programs", retrieveAllPrograms)
	app.GET("/api/programs/:programId", retrieveProgram)
	app.DELETE("/api/programs/:programId", deleteProgram)
	app.GET("/api/programs/:programId/info", notImplemented)
	app.GET("/api/programs/:programId/events", notImplemented)
	app.GET("/api/programs/:programId/enrolled-students", notImplemented)
	app.GET("/api/programs/:programId/enrolled-students/:enrolled-studentsId/status", notImplemented)
	app.GET("/api/justifications", notImplemented)
	app.POST("/api/justifications", notImplemented)
	app.GET("/api/justifications/:justificationId", notImplemented)
	app.PUT("/api/justifications/:justificationId/validatation", notImplemented)
	app.DELETE("/api/justifications/:justificationId/validatation", notImplemented)
	app.GET("/api/langues", retrieveLanguages)
	app.PUT("/api/langues/update", updateLanguages)
}

func createUser(c *gin.Context) {
	
}

func deleteUser(c *gin.Context) {
	
}

func createMessage(c *gin.Context) {
	
}

func createEvent(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	event := &DTO.Event{}
	if err := c.BindJSON(event); err != nil {
		log.Println(err.Error())
	}
	log.Println(event)
	c.JSON(http.StatusOK, event)
}

func retrieveEvent(c *gin.Context) {
	
}

func updateEvent(c *gin.Context) {
	
}

func deleteEvent(c *gin.Context) {
	
}

func subscribeEvent(c *gin.Context) {
	
}

func unsubscribeEvent(c *gin.Context) {
	
}

func createProgram(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	program := &DTO.Program{}
	if err := c.BindJSON(program); err != nil {
		log.Println(err.Error())
	}
	if err := program.CheckField(); err != "" {
		c.JSON(http.StatusBadRequest, DTO.UserMsg{Information: translate.T("programCreationUsage"), Error: err})
	} else {
		c.JSON(defineUserMessage(service.CreateProgram(program)))
	}
}

func retrieveAllPrograms(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	programs, _ := service.RetrieveAllPrograms()
	c.JSON(http.StatusOK, programs)
}

func retrieveProgram(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	id, err := strconv.Atoi(c.Params.ByName("programId"))
	if err != nil {
		log.Println(err)
	}
	program, rMsg := service.RetrieveProgram(id)
	if (rMsg.Status == 0) {
		c.JSON(http.StatusOK, program)
	} else if rMsg.Status == 1 {
		c.JSON(http.StatusBadRequest, DTO.UserMsg{Error:rMsg.Msg})
	} else if rMsg.Status == -1 {
		c.JSON(http.StatusInternalServerError, DTO.UserMsg{Error:rMsg.Msg})
	}
}

func deleteProgram(c *gin.Context) {
	c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
	id, err := strconv.Atoi(c.Params.ByName("programId"))
	if err != nil {
		log.Println(err)
	}
	if rMsg := service.DeleteProgram(id); rMsg.Status == 0 {
		c.JSON(http.StatusOK, DTO.UserMsg{Information:rMsg.Msg})
	} else if rMsg.Status == 1 {
		c.JSON(http.StatusBadRequest, DTO.UserMsg{Error:rMsg.Msg})
	} else if rMsg.Status == -1 {
		c.JSON(http.StatusInternalServerError, DTO.UserMsg{Error:rMsg.Msg})
	}
}

func retrieveProgramInfo(c *gin.Context) {
	
}

func retriveProgramsEvent(c *gin.Context) {
	
}

func retrieveEnrolledStudents(c *gin.Context) {
	
}

func retrieveStatusOfEnreolledStudent(c *gin.Context) {
	
}

func retrieveJustifications(c *gin.Context) {
	
}

func createJustification(c *gin.Context) {
	
}

func retrieveJustification(c *gin.Context) {
	
}

func acceptJustification(c *gin.Context) {
	
}

func refuseJustification(c *gin.Context) {
	
}

func retrieveLanguages(c *gin.Context) {
	c.JSON(http.StatusOK, i18n.LanguageTags())
}

func updateLanguages(c *gin.Context) {
	cookieLang, err := c.Cookie("lang")
	if err != nil {
		log.Println(err)
	}
	acceptLang := c.Request.Header.Get("Accept-Language")
	translate.ChangeLang(cookieLang, acceptLang)
}

func notImplemented(c *gin.Context) {
	c.Writer.WriteHeader(http.StatusNotImplemented)
}

func defineUserMessage(r utils.ReturnMsg) (int, DTO.UserMsg) {
	if r.Status == 0 {
		return http.StatusCreated, DTO.UserMsg{Information: r.Msg}
	} else if r.Status == 1 {
		return http.StatusBadRequest, DTO.UserMsg{Error: r.Msg}
	} else {
		return http.StatusInternalServerError, DTO.UserMsg{Error: r.Msg}
	}
}