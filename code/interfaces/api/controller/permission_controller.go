package controller

import (
	"admin/microservice/infra/dto"
	"admin/microservice/interfaces/api/exceptions"
	"admin/microservice/internal/application"
	"github.com/gin-gonic/gin"
	"net/http"
)

type permissionController struct {
	permissionService application.IPermissionService
}

func NewPermissionController(
	permissionService application.IPermissionService,
) *permissionController {
	return &permissionController{
		permissionService: permissionService,
	}
}
func (p permissionController) Test(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"data": "สำเร็จ"})

}
func (p permissionController) CreatePermissionController(c *gin.Context) {
	r := dto.NewPermission()
	if err := c.ShouldBindJSON(&r); err != nil {
		statusCode := http.StatusBadRequest
		c.AbortWithStatusJSON(statusCode, exceptions.ThrowNewException(statusCode, err.Error()))
		return
	}
	err := p.permissionService.CreatePermissionService(r)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"data": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "สร้างสำเร็จ"})

}
