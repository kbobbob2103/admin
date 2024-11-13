package application

import (
	"admin/microservice/infra/dto"
	"admin/microservice/internal/domain/repositoty"
	"fmt"
)

type IRankBarService interface {
	CreateRankService(data dto.Rank) error
	FindAllRankBarService() ([]dto.Rank, error)
}
type RankBarService struct {
	ranksRepo repositoty.IRankRepo
}

func NewRankBarService(
	ranksRepo repositoty.IRankRepo) IRankBarService {
	return RankBarService{
		ranksRepo: ranksRepo,
	}
}
func (a RankBarService) CreateRankService(data dto.Rank) error {
	ranks, err := a.ranksRepo.FindAllRank()
	if err != nil {
		return err
	}
	for _, r := range ranks {
		if (data.Min >= r.Min && data.Min <= r.Max) || (data.Max >= r.Min && data.Max <= r.Max) ||
			(r.Min >= data.Min && r.Min <= data.Max) || (r.Max >= data.Min && r.Max <= data.Max) {
			return fmt.Errorf("ช่วงเวลา Min และ Max ของข้อมูลใหม่ทับกับช่วงที่มีอยู่แล้ว")
		}
	}
	return a.ranksRepo.CreateRank(data)
}
func (a RankBarService) FindAllRankBarService() ([]dto.Rank, error) {
	return a.ranksRepo.FindAllRank()
}
