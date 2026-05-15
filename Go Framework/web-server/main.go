package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	
	
)



func main() {
	var router *gin.Engine = gin.Default()

	router.GET("/people", filecontext())

	router.Run(":8001")
}

func getPeople(context *gin.Context) {
	context.IndentedJSON(http.StatusOK, filecontext())
}
