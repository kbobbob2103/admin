package repositoty

import "admin/microservice/infra/dto"

type IPermissionRepo interface {
	CreatePermission(data dto.Permission) error
	UpdatePermission(data dto.Permission) error
	FindOnePermissionByID(id string) (dto.Permission, error)
}
