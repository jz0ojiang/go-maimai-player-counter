package api

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/jz0ojiang/go-maimai-player-counter/service"
)

func HandleGetArcadeList(c *gin.Context) {
	arcades, err := service.GetArcadeList()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": -1, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "data": arcades})
}

func HandleGetArcadeByArcadeID(c *gin.Context) {
	arcadeID := c.Param("arcadeID")
	if arcadeID == "" {
		c.JSON(http.StatusBadRequest, gin.H{"code": -1, "message": "arcadeID is required"})
		return
	}
	arcadeIDInt, err := strconv.ParseInt(arcadeID, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": -1, "message": err.Error()})
		return
	}
	arcade, err := service.GetArcadeByArcadeID(int(arcadeIDInt))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": -1, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "data": arcade})
}

func HandleGetArcadeListByCityCode(c *gin.Context) {
	cityCode := c.Param("cityCode")
	if cityCode == "" {
		c.JSON(http.StatusBadRequest, gin.H{"code": -1, "message": "cityCode is required"})
		return
	}
	arcades, err := service.GetArcadeListByCityCode(cityCode)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": -1, "message": err.Error()})
		return
	}
	if len(arcades) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"code": -1, "data": []gin.H{}, "message": "no arcade found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "data": arcades})
}

func HandleGetArcadeListByProvinceCode(c *gin.Context) {
	provinceCode := c.Param("provinceCode")
	if provinceCode == "" {
		c.JSON(http.StatusBadRequest, gin.H{"code": -1, "message": "provinceCode is required"})
		return
	}
	arcades, err := service.GetArcadeListByProvinceCode(provinceCode)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": -1, "message": err.Error()})
		return
	}
	if len(arcades) == 0 {
		c.JSON(http.StatusNotFound, gin.H{"code": -1, "data": []gin.H{}, "message": "no arcade found"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "data": arcades})
}

type PostCreateArcadeRequest struct {
	ArcadeName   string `json:"arcade_name" binding:"required"`
	Address      string `json:"address" binding:"required"`
	MachineCount int    `json:"machine_count" binding:"required"`
	ProvinceCode int    `json:"province_code" binding:"required"`
	CityCode     int    `json:"city_code" binding:"required"`
	Token        string `json:"token" binding:"required"`
}

func HandlePostCreateCustomArcade(c *gin.Context) {
	var postCreateArcadeRequest PostCreateArcadeRequest
	if err := c.ShouldBindJSON(&postCreateArcadeRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": -1, "message": err.Error()})
		return
	}
	if !service.ValidateTotp(postCreateArcadeRequest.Token) {
		c.JSON(http.StatusBadRequest, gin.H{"code": -1, "message": "invalid token"})
		return
	}
	arcadeID, err := service.CreateCustomArcade(postCreateArcadeRequest.ArcadeName,
		postCreateArcadeRequest.MachineCount,
		postCreateArcadeRequest.Address,
		postCreateArcadeRequest.ProvinceCode,
		postCreateArcadeRequest.CityCode)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": -1, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0, "data": arcadeID})
}

type PostDeleteArcadeRequest struct {
	ArcadeID int    `json:"arcade_id" binding:"required"`
	Token    string `json:"token" binding:"required"`
}

func HandlePostDeleteArcade(c *gin.Context) {
	var postDeleteArcadeRequest PostDeleteArcadeRequest
	if err := c.ShouldBindJSON(&postDeleteArcadeRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": -1, "message": err.Error()})
		return
	}
	if !service.ValidateTotp(postDeleteArcadeRequest.Token) {
		c.JSON(http.StatusBadRequest, gin.H{"code": -1, "message": "invalid token"})
		return
	}
	err := service.DeleteArcadeByArcadeID(postDeleteArcadeRequest.ArcadeID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": -1, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0})
}

type PostUpdateAllArcadeRequest struct {
	Token string `json:"token" binding:"required"`
}

func HandlePostUpdateAllArcade(c *gin.Context) {
	var postUpdateAllArcadeRequest PostUpdateAllArcadeRequest
	if err := c.ShouldBindJSON(&postUpdateAllArcadeRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": -1, "message": err.Error()})
		return
	}
	if !service.ValidateTotp(postUpdateAllArcadeRequest.Token) {
		c.JSON(http.StatusBadRequest, gin.H{"code": -1, "message": "invalid token"})
		return
	}
	err := service.UpdateArcadeListWithArcadeMap()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": -1, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0})
}
