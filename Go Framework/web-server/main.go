package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
	
	
)



func main() {
	var router *gin.Engine = gin.Default()

	router.GET("/count", filecount)

	router.Run(":8001")
}
func filecount(context *gin.Context) {
    context.IndentedJSON(http.StatusOK, gin.H{"message": "ok"})
}

