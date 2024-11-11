package controller

import (
	"admin/microservice/infra/dto"
	"admin/microservice/interfaces/api/exceptions"
	"admin/microservice/internal/application"
	"github.com/gin-gonic/gin"
	"net/http"
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
		statusCode := http.StatusBadRequest
		c.AbortWithStatusJSON(statusCode, exceptions.ThrowNewException(statusCode, err.Error()))
		return
	}
	err := a.employeeService.CreateEmployeeService(r)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"data": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "สร้างสำเร็จ"})

}
