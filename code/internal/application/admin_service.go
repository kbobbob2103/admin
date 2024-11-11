package application

import (
	"admin/microservice/infra/dto"
	"admin/microservice/internal/domain/repositoty"
)

type IEmployeeService interface {
	CreateEmployeeService(data dto.Employee) error
}
type employeeService struct {
	employeeRepo repositoty.IEmployeeRepo
}

func NewEmployeeService(
	employeeRepo repositoty.IEmployeeRepo) IEmployeeService {
	return employeeService{
		employeeRepo: employeeRepo,
	}
}
func (a employeeService) CreateEmployeeService(data dto.Employee) error {
	return a.employeeRepo.CreateEmployee(data)
}
