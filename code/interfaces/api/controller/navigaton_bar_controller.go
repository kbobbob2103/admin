package controller

import (
	"admin/microservice/infra/dto"
	"admin/microservice/interfaces/api/exceptions"
	"admin/microservice/internal/application"
	"github.com/gin-gonic/gin"
	"net/http"
)

type navigationController struct {
	navigationService application.INavigationBarService
}

func NewNavigationController(
	navigationService application.INavigationBarService,
) *navigationController {
	return &navigationController{
		navigationService: navigationService,
	}
}

func (n navigationController) CreateNavigationController(c *gin.Context) {
	r := dto.NewNavigationBar()
	if err := c.ShouldBindJSON(&r); err != nil {
		statusCode := http.StatusBadRequest
		c.AbortWithStatusJSON(statusCode, exceptions.ThrowNewException(statusCode, err.Error()))
		return
	}
	err := n.navigationService.CreateNavigationService(r)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"data": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "สร้างสำเร็จ"})

}
