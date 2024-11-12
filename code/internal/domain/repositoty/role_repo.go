package repositoty

import "admin/microservice/infra/dto"

type IRoleRepo interface {
	CreateRole(data dto.Role) error
	UpdateRole(data dto.Role) error
	FindAll(
		roleId []string,
		search string,
		page,
		limit int,
	) ([]dto.Role, error)
	FindOneRoleByID(id string) (dto.Role, error)
	Delete(id string) error
}
