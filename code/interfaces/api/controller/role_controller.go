package controller

import (
	"admin/microservice/exception"
	"admin/microservice/infra/dto"
	"admin/microservice/interfaces/api/controller/res"
	"admin/microservice/internal/application"
	"github.com/gin-gonic/gin"
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
		res.HandleError(
			c,
			exception.NewAppError(exception.ErrCodeBadRequest, err.Error()),
			"",
		)
		return
	}
	err := a.roleService.CreateRoleService(r)
	if err != nil {
		res.HandleError(c, err, "object")
		return
	}
	res.HandleSuccess(c, map[string]string{"role_id": r.RoleID}, "created")

}
func (a roleController) UpdateRoleController(c *gin.Context) {
	var r dto.Role
	if err := c.ShouldBindJSON(&r); err != nil {
		res.HandleError(
			c,
			exception.NewAppError(exception.ErrCodeBadRequest, err.Error()),
			"",
		)
		return
	}
	err := a.roleService.UpdateRoleService(r)
	if err != nil {
		res.HandleError(c, err, "object")
		return
	}
	res.HandleSuccess(c, map[string]string{"role_id": r.RoleID}, "updated")

}

func (a roleController) GetOneRoleController(c *gin.Context) {
	id := c.Param("role_id")

	data, err := a.roleService.FindOneRoleService(id)
	if err != nil {
		res.HandleError(c, err, "object")
		return
	}
	res.HandleSuccess(c, data, "")
}
func (a roleController) GetAllRoleController(c *gin.Context) {
	var r dto.QueryEmployee
	if err := c.ShouldBindQuery(&r); err != nil {
		res.HandleError(
			c,
			exception.NewAppError(exception.ErrCodeBadRequest, err.Error()),
			"",
		)
		return
	}
	data, err := a.roleService.FindAllRoleService(r)
	if err != nil {
		res.HandleError(c, err, "slice")
		return
	}
	res.HandleSuccess(c, data, "")
}
