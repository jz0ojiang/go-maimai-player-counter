package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jz0ojiang/go-maimai-player-counter/api"
	"github.com/jz0ojiang/go-maimai-player-counter/conf"
	"github.com/jz0ojiang/go-maimai-player-counter/middleware"
)

func main() {
	if !conf.GetConfig().GetDebug() {
		gin.SetMode(gin.ReleaseMode)
	}

	r := gin.Default()

	// Middleware
	r.Use(middleware.Cors())

	// Province
	r.GET("/getProvinceList", api.HandleGetProvinceList)

	// City
	r.GET("/getCityList/:provinceCode", api.HandleGetCityList)
	r.GET("/getCity/:cityCode", api.HandleGetCityByCityCode)

	// Arcade
	r.GET("/getArcade/:arcadeID", api.HandleGetArcadeByArcadeID)
	r.GET("/getArcadeList/", api.HandleGetArcadeList)
	r.GET("/getArcadeList/city/:cityCode", api.HandleGetArcadeListByCityCode)
	r.GET("/getArcadeList/province/:provinceCode", api.HandleGetArcadeListByProvinceCode)
	r.POST("/createCustomArcade", api.HandlePostCreateCustomArcade)
	r.POST("/deleteArcade", api.HandlePostDeleteArcade)
	r.POST("/updateArcade", api.HandlePostUpdateAllArcade)

	// Count
	r.GET("/getCount/:arcadeID", api.HandleGetCountLogsByArcadeID)
	r.POST("/logCount", api.HandlePostCountLog)

	// Token
	r.POST("/generateToken", api.HandlePostGenerateToken)

	r.Run(conf.GetConfig().GetHost())
}
