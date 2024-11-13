package repositoty

import "admin/microservice/infra/dto"

type IEmployeeRepo interface {
	FindAll(
		employeeIds,
		roleId []string,
		search string,
		page,
		limit int,
	) ([]dto.Employee, error)
	FindOneUserName(userName string) (dto.Employee, error)
	FindOneByID(id string) (dto.Employee, error)
	CreateEmployee(data dto.Employee) error
	UpdateOne(data dto.Employee) error
	UpdateStatus(id string, status bool) error
	UpdateStatusEmployee(id string, statusEmployee dto.StatusEmployee) error
	FindCount(
		employeeIds,
		roleId []string,
	) (int64, error)
}
