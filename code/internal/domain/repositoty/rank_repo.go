package repositoty

import "admin/microservice/infra/dto"

type IRankRepo interface {
	CreateRank(data dto.Rank) error
	FindAllRank() ([]dto.Rank, error)
}
