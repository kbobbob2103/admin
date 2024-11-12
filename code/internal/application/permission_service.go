package application

import (
	"admin/microservice/infra/dto"
	"admin/microservice/internal/domain/repositoty"
)

type IPermissionService interface {
	CreatePermissionService(data dto.Permission) error
	UpdatePermissionService(data dto.Permission) error
	FinaOnePermission(id string) (dto.Permission, error)
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
func (p permissionService) UpdatePermissionService(data dto.Permission) error {
	permission, err := p.permissionRepo.FindOnePermissionByID(data.PermissionID)
	if err != nil {
		return err
	}
	permission.UpdatePermission(data)
	return p.permissionRepo.UpdatePermission(permission)
}
func (p permissionService) FinaOnePermission(id string) (dto.Permission, error) {
	return p.permissionRepo.FindOnePermissionByID(id)
}
