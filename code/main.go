package main

import (
	"admin/microservice/interfaces/api/router"
	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	router.ApplicationV1Routes(r)
	startServer(r)

}

func startServer(router *gin.Engine) {
	router.Run(":3000")
}
