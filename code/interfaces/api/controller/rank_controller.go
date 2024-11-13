package controller

import (
	"admin/microservice/infra/dto"
	"admin/microservice/interfaces/api/controller/res"
	"admin/microservice/interfaces/api/exceptions"
	"admin/microservice/internal/application"
	"github.com/gin-gonic/gin"
	"net/http"
)

type rankController struct {
	rankService application.IRankBarService
}

func NewRankController(
	rankService application.IRankBarService,
) *rankController {
	return &rankController{
		rankService: rankService,
	}
}

func (n rankController) CreateRankController(c *gin.Context) {
	r := dto.NewRank()
	if err := c.ShouldBindJSON(&r); err != nil {
		statusCode := http.StatusBadRequest
		c.AbortWithStatusJSON(statusCode, exceptions.ThrowNewException(statusCode, err.Error()))
		return
	}
	err := n.rankService.CreateRankService(r)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"data": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": "สร้างสำเร็จ"})

}
func (r rankController) GetAllController(c *gin.Context) {
	data, err := r.rankService.FindAllRankBarService()
	if err != nil {
		res.HandleError(c, err, "slice")
		return
	}
	res.HandleSuccess(c, data, "")
}
