package injector

import "admin/microservice/interfaces/api/controller"

var (
	PermissionController = controller.NewPermissionController(PermissionService)
	EmployeeController   = controller.NewEmployeeController(EmployeeService)
	NavigationController = controller.NewNavigationController(NavigationBarService)
)
