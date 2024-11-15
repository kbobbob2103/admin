package injector

import "admin/microservice/interfaces/api/controller"

var (
	PermissionController = controller.NewPermissionController(PermissionService)
	EmployeeController   = controller.NewEmployeeController(EmployeeService)
	NavigationController = controller.NewNavigationController(NavigationBarService)
	RoleController       = controller.NewRoleController(RoleService)
	RankController       = controller.NewRankController(RankService)
)
