package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jz0ojiang/go-maimai-player-counter/service"
)

func HandleGetCityList(c *gin.Context) {
	province := c.Param("provinceCode")
	if province == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    -1,
			"message": "province is required",
		})
		return
	}
	cities, err := service.GetCityListByProvinceCode(province)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    -1,
			"message": err.Error(),
		})
		return
	}
	if len(cities) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    -1,
			"data":    []string{},
			"message": "no city found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "success",
		"data":    cities,
	})
}

func HandleGetCityByCityCode(c *gin.Context) {
	cityCode := c.Param("cityCode")
	if cityCode == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    -1,
			"message": "cityCode is required",
		})
	}
	city, err := service.GetCityByCode(cityCode)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    -1,
			"message": err.Error(),
		})
		return
	}
	if city.Code == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    -1,
			"data":    gin.H{},
			"message": "no city found",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "success",
		"data":    city,
	})
}
