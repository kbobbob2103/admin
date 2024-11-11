package router

import (
	"admin/microservice/injector"
	"github.com/gin-gonic/gin"
)

func ApplicationV1Routes(engine *gin.Engine) {
	route := engine.Group("/api/admin/v1/admin")
	billCTL := injector.PermissionController
	route.GET("test", billCTL.Test)
	route.POST("", billCTL.CreatePermissionController)
}
