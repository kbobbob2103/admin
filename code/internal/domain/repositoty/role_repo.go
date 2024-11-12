package repositoty

import "admin/microservice/infra/dto"

type IRoleRepo interface {
	CreateRole(data dto.Role) error
}
