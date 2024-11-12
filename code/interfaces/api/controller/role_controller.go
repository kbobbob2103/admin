package controller

import (
	"admin/microservice/infra/dto"
	"admin/microservice/interfaces/api/exceptions"
	"admin/microservice/internal/application"
	"github.com/gin-gonic/gin"
	"net/http"
)

type roleController struct {
	roleService application.IRoleService
}

func NewRoleController(
	roleService application.IRoleService,
) *roleController {
	return &roleController{
		roleService: roleService,
	}
}

func (a roleController) CreateRoleController(c *gin.Context) {
	r := dto.NewRole()
	if err := c.ShouldBindJSON(&r); err != nil {
		statusCode := http.StatusBadRequest
		c.AbortWithStatusJSON(statusCode, exceptions.ThrowNewException(statusCode, err.Error()))
		return
	}
	err := a.roleService.CreateRoleService(r)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"data": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "สร้างสำเร็จ"})

}
