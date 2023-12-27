package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jz0ojiang/go-maimai-player-counter/service"
)

func HandleGetProvinceList(c *gin.Context) {
	provinces, err := service.GetProvinceList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    -1,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "success",
		"data":    provinces,
	})
}
