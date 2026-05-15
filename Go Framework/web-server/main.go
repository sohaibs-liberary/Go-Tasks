package main

import (
	"net/http"
	"web-server/route"
	"github.com/gin-gonic/gin"
)



func main() {
	var router *gin.Engine = gin.Default()

	router.GET("/count", filecount)

	router.Run(":8001")
}
func filecount(context *gin.Context) {
	route.Filecontext()
    context.IndentedJSON(http.StatusOK, gin.H{"Status": "Done"})
}

