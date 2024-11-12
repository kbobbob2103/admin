package application

import (
	"admin/microservice/infra/dto"
	"admin/microservice/internal/domain/repositoty"
)

type IEmployeeService interface {
	CreateEmployeeService(data dto.Employee) error
	UpdateEmployeeService(data dto.Employee) error
	FindOneEmployee(employeeId string) (dto.Employee, error)
	FindAll(query dto.QueryEmployee) ([]dto.Employee, error)
	UpdateStatus(id string, status bool) error
	UpdateStatusEmployee(id string, statusEmployee dto.StatusEmployee) error
	FindCount(query dto.QueryEmployee) (int64, error)
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
func (a employeeService) UpdateEmployeeService(data dto.Employee) error {
	employee, err := a.employeeRepo.FindOneByID(data.EmployeeID)
	if err != nil {
		return err
	}
	employee.UpdateEmployee(data)
	return a.employeeRepo.UpdateOne(employee)
}
func (a employeeService) FindOneEmployee(employeeId string) (dto.Employee, error) {
	return a.employeeRepo.FindOneByID(employeeId)
}
func (a employeeService) FindAll(query dto.QueryEmployee) ([]dto.Employee, error) {
	roleIds := make([]string, 0)
	if query.RoleID != "" {
		roleIds = append(roleIds, query.RoleID)
	}
	return a.employeeRepo.FindAll(
		[]string{},
		roleIds,
		query.Search,
		query.Page,
		query.Limit,
	)
}

func (a employeeService) UpdateStatus(id string, status bool) error {
	return a.employeeRepo.UpdateStatus(id, status)
}
func (a employeeService) UpdateStatusEmployee(id string, statusEmployee dto.StatusEmployee) error {
	return a.employeeRepo.UpdateStatusEmployee(id, statusEmployee)
}
func (a employeeService) FindCount(query dto.QueryEmployee) (int64, error) {
	roleIds := make([]string, 0)
	if query.RoleID != "" {
		roleIds = append(roleIds, query.RoleID)
	}
	return a.employeeRepo.FindCount(
		[]string{},
		roleIds,
	)
}