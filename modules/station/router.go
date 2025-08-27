package station

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/oktaharis/mrt-schedules/common/response"
)

func Initiate(router *gin.RouterGroup) {
	stationService := NewService()

	station := router.Group("/stations")
	station.GET("", func(c *gin.Context) {
		GetAllStation(c, stationService)
	})
	
	station.GET("/:id", func(c *gin.Context) {
		CheckScheduleByStation(c, stationService)
	})
}

func GetAllStation(c *gin.Context, service Service) {
	datas, err := service.GetAllStation()

	if err != nil {
		c.JSON(http.StatusBadRequest, response.ApiResponse{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		response.ApiResponse{
			Success: true,
			Message: "success",
			Data:    datas,
		},
	)
}

func CheckScheduleByStation(c *gin.Context, service Service) {
	id := c.Param("id")

	datas, err := service.CheckScheduleByStation(id)

	if err != nil {
		c.JSON(http.StatusBadRequest, response.ApiResponse{
			Success: false,
			Message: err.Error(),
			Data:    nil,
		},
		)
		return
	}

	c.JSON(
		http.StatusOK,
		response.ApiResponse{
			Success: true,
			Message: "success",
			Data:    datas,
		},
	)
}