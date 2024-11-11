package application

import (
	"admin/microservice/infra/dto"
	"admin/microservice/internal/domain/repositoty"
)

type IPermissionService interface {
	CreatePermissionService(data dto.Permission) error
}
type permissionService struct {
	permissionRepo repositoty.IPermissionRepo
}

func NewPermissionService(
	permissionRepo repositoty.IPermissionRepo) IPermissionService {
	return permissionService{
		permissionRepo: permissionRepo,
	}
}
func (p permissionService) CreatePermissionService(data dto.Permission) error {
	return p.permissionRepo.CreatePermission(data)
}
