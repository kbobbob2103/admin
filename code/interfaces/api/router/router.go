package router

import (
	"admin/microservice/injector"
	"github.com/gin-gonic/gin"
)

func ApplicationV1Routes(engine *gin.Engine) {

	routeAmin := engine.Group("/api/admin/v1/employee")
	adminCTL := injector.EmployeeController
	routeAmin.POST("", adminCTL.CreateEmployeeController)
	routeAmin.PATCH("", adminCTL.UpdateEmployeeController)
	routeAmin.PATCH("status", adminCTL.UpdateStatusController)
	routeAmin.DELETE(":employee_id", adminCTL.DeleteController)
	routeAmin.GET(":employee_id", adminCTL.GetOneEmployeeController)
	routeAmin.GET("", adminCTL.GetAllEmployeeController)
	routeAmin.GET("count", adminCTL.GetCountController)

	route := engine.Group("/api/admin/v1/permission")
	permissionCTL := injector.PermissionController
	route.POST("", permissionCTL.CreatePermissionController)

	routeNavigation := engine.Group("/api/admin/v1/navigation_bar")
	navigationCTL := injector.NavigationController
	routeNavigation.POST("", navigationCTL.CreateNavigationController)
	routeNavigation.GET("", navigationCTL.GetAllController)

	routeRole := engine.Group("/api/admin/v1/role")
	roleCTL := injector.RoleController
	routeRole.POST("", roleCTL.CreateRoleController)
	routeRole.GET("", roleCTL.GetAllRoleController)
	routeRole.GET(":role_id", roleCTL.GetOneRoleController)
	routeRole.PATCH("", roleCTL.UpdateRoleController)
}
