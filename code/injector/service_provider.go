package injector

import "admin/microservice/internal/application"

var (
	PermissionService = application.NewPermissionService(PermissionRepository)
)