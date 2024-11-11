package repositoty

import "admin/microservice/infra/dto"

type IEmployeeRepo interface {
	CreateEmployee(data dto.Employee) error
	UpdateOne(data dto.Employee) error
	UpdateStatus(id string, status bool) error
	UpdateStatusEmployee(id string, statusEmployee dto.StatusEmployee) error
}
