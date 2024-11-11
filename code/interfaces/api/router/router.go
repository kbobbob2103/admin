package router

import (
	"admin/microservice/injector"
	"github.com/gin-gonic/gin"
)

func ApplicationV1Routes(engine *gin.Engine) {

	routeAmin := engine.Group("/api/admin/v1/admin")
	adminCTL := injector.EmployeeController
	routeAmin.POST("", adminCTL.CreateEmployeeController)

	route := engine.Group("/api/admin/v1/permission")
	permissionCTL := injector.PermissionController
	route.GET("test", permissionCTL.Test)
	route.POST("", permissionCTL.CreatePermissionController)

	routeNavigation := engine.Group("/api/admin/v1/navigation_bar")
	navigationCTL := injector.NavigationController
	routeNavigation.POST("", navigationCTL.CreateNavigationController)
}
