package api

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jz0ojiang/go-maimai-player-counter/service"
)

func HandleGetCountLogsByArcadeID(c *gin.Context) {
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
	countLogs, err := service.GetCountLogsByArcadeID(int(arcadeIDInt))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": -1, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "success",
		"data":    countLogs,
	})
}

func HandleGetCountLogsByCityCode(c *gin.Context) {
	cityCode := c.Param("cityCode")
	if cityCode == "" {
		c.JSON(http.StatusBadRequest, gin.H{"code": -1, "message": "cityCode is required"})
		return
	}
	if len(cityCode) == 12 {
		cityCode = cityCode[0:4]
	}
	cityCodeInt, err := strconv.ParseInt(cityCode, 10, 64)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": -1, "message": err.Error()})
		return
	}
	counts, err := service.GetCountsByCity(int(cityCodeInt))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": -1, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"code":    0,
		"message": "success",
		"data":    counts,
	})
}

type CountLogRequest struct {
	ArcadeID int    `json:"arcade_id" binding:"required"`
	Count    *int   `json:"count" binding:"required"`
	Type     *int   `json:"type" binding:"required"`
	Token    string `json:"token" binding:"required"`
}

func HandlePostCountLog(c *gin.Context) {
	var countLog CountLogRequest
	if err := c.ShouldBindJSON(&countLog); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"code": -1, "message": err.Error()})
		return
	}
	switch *countLog.Type {
	case 0:
		c.JSON(http.StatusTeapot, gin.H{"code": -1, "message": "I'm a teapot"})
		return
	case 1:
		if !service.VerifyCaptcha(countLog.Token) {
			c.JSON(http.StatusBadRequest, gin.H{"code": -1, "message": "invalid token"})
			return
		}
	case 2:
		if !service.VerifyToken(countLog.Token) {
			c.JSON(http.StatusBadRequest, gin.H{"code": -1, "message": "invalid token"})
			return
		}
	default:
		c.JSON(http.StatusBadRequest, gin.H{"code": -1, "message": "invalid type"})
		return
	}
	if *countLog.Count < 0 {
		c.JSON(http.StatusBadRequest, gin.H{"code": -1, "message": "invalid count"})
		return
	}
	if err := service.AddCountLog(service.CountLog{
		ArcadeId:        countLog.ArcadeID,
		Count:           *countLog.Count,
		UpdateTimestamp: time.Now().UnixNano() / 1e6,
		Type:            *countLog.Type,
	}); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"code": -1, "message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"code": 0})
}
