package controller

import (
	"admin/microservice/exception"
	"admin/microservice/infra/dto"
	"admin/microservice/interfaces/api/controller/res"
	"admin/microservice/internal/application"
	"github.com/gin-gonic/gin"
)

type employeeController struct {
	employeeService application.IEmployeeService
}

func NewEmployeeController(
	employeeService application.IEmployeeService,
) *employeeController {
	return &employeeController{
		employeeService: employeeService,
	}
}

func (a employeeController) CreateEmployeeController(c *gin.Context) {
	r := dto.NewEmployee()
	if err := c.ShouldBindJSON(&r); err != nil {
		res.HandleError(
			c,
			exception.NewAppError(exception.ErrCodeBadRequest, err.Error()),
			"",
		)
		return
	}
	err := a.employeeService.CreateEmployeeService(r)
	if err != nil {
		res.HandleError(c, err, "object")
		return
	}
	res.HandleSuccess(c, map[string]string{"employee_id": r.EmployeeID}, "created")
}
func (a employeeController) UpdateEmployeeController(c *gin.Context) {
	var r dto.Employee
	if err := c.ShouldBindJSON(&r); err != nil {
		res.HandleError(
			c,
			exception.NewAppError(exception.ErrCodeBadRequest, err.Error()),
			"",
		)
		return
	}
	err := a.employeeService.UpdateEmployeeService(r)
	if err != nil {
		res.HandleError(c, err, "object")
		return
	}
	res.HandleSuccess(c, map[string]string{"employee_id": r.EmployeeID}, "updated")
}
func (a employeeController) GetOneEmployeeController(c *gin.Context) {
	id := c.Param("employee_id")

	data, err := a.employeeService.FindOneEmployee(id)
	if err != nil {
		res.HandleError(c, err, "object")
		return
	}
	res.HandleSuccess(c, data, "")
}
func (a employeeController) GetAllEmployeeController(c *gin.Context) {
	var r dto.QueryEmployee
	if err := c.ShouldBindQuery(&r); err != nil {
		res.HandleError(
			c,
			exception.NewAppError(exception.ErrCodeBadRequest, err.Error()),
			"",
		)
		return
	}
	data, err := a.employeeService.FindAll(r)
	if err != nil {
		res.HandleError(c, err, "slice")
		return
	}
	res.HandleSuccess(c, data, "")
}
func (a employeeController) UpdateStatusController(c *gin.Context) {
	var r dto.Employee
	if err := c.ShouldBindJSON(&r); err != nil {
		res.HandleError(
			c,
			exception.NewAppError(exception.ErrCodeBadRequest, err.Error()),
			"",
		)
		return
	}
	err := a.employeeService.UpdateStatus(r.EmployeeID, r.Status)
	if err != nil {
		res.HandleError(c, err, "object")
		return
	}
	res.HandleSuccess(c, map[string]string{}, "updated")
}
func (a employeeController) DeleteController(c *gin.Context) {
	id := c.Param("employee_id")

	err := a.employeeService.UpdateStatusEmployee(id, dto.Disable)
	if err != nil {
		res.HandleError(c, err, "object")
		return
	}
	res.HandleSuccess(c, map[string]string{}, "updated")
}
func (a employeeController) GetCountController(c *gin.Context) {
	var r dto.QueryEmployee
	if err := c.ShouldBindQuery(&r); err != nil {
		res.HandleError(
			c,
			exception.NewAppError(exception.ErrCodeBadRequest, err.Error()),
			"",
		)
		return
	}
	data, err := a.employeeService.FindCount(r)
	if err != nil {
		res.HandleError(c, err, "object")
		return
	}
	res.HandleSuccess(c, dto.Count{Count: data}, "")
}
