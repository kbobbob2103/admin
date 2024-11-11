package application

import (
	"admin/microservice/infra/dto"
	"admin/microservice/internal/domain/repositoty"
)

type INavigationBarService interface {
	CreateNavigationService(data dto.NavigationBar) error
}
type navigationBarService struct {
	navigationBarRepo repositoty.INavigationRepo
}

func NewNavigationBarService(
	navigationBarRepo repositoty.INavigationRepo) INavigationBarService {
	return navigationBarService{
		navigationBarRepo: navigationBarRepo,
	}
}
func (a navigationBarService) CreateNavigationService(data dto.NavigationBar) error {
	return a.navigationBarRepo.CreateNavigation(data)
}
