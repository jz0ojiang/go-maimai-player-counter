package api

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jz0ojiang/go-maimai-player-counter/service"
)

type TokenRequest struct {
	Token  string `json:"token" binding:"required"`
	Remark string `json:"remark"`
}

func HandlePostGenerateToken(c *gin.Context) {
	var tokenRequest TokenRequest
	if err := c.ShouldBindJSON(&tokenRequest); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    -1,
			"message": err.Error(),
		})
		return
	}
	if !service.ValidateTotp(tokenRequest.Token) {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    -1,
			"message": "invalid token",
		})
		return
	}
	token, err := service.GenerateToken()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    -1,
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"token": token,
	})
}
