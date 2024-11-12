package application

import (
	"admin/microservice/infra/dto"
	"admin/microservice/internal/domain/repositoty"
)

type IRoleService interface {
	CreateRoleService(data dto.Role) error
}
type roleService struct {
	roleRepo repositoty.IRoleRepo
}

func NewRoleService(
	roleRepo repositoty.IRoleRepo) IRoleService {
	return roleService{
		roleRepo: roleRepo,
	}
}
func (a roleService) CreateRoleService(data dto.Role) error {
	return a.roleRepo.CreateRole(data)
}
