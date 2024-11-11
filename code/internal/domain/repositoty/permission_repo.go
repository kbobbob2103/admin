package repositoty

import "admin/microservice/infra/dto"

type IPermissionRepo interface {
	CreatePermission(data dto.Permission) error
}
