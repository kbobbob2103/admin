package injector

import "admin/microservice/internal/application"

var (
	PermissionService    = application.NewPermissionService(PermissionRepository)
	EmployeeService      = application.NewEmployeeService(EmployeeRepository)
	NavigationBarService = application.NewNavigationBarService(NavigationBarRepository)
)
