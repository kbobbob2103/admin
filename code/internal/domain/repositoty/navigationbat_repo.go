package repositoty

import "admin/microservice/infra/dto"

type INavigationRepo interface {
	CreateNavigation(data dto.NavigationBar) error
}
