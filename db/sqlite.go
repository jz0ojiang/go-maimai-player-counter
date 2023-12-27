package db

import (
	"os"

	"github.com/glebarez/sqlite"
	"github.com/jz0ojiang/go-maimai-player-counter/conf"
	"gorm.io/gorm"
)

var SqliteDB *gorm.DB

func init() {
	var err error
	SqliteDB, err = gorm.Open(sqlite.Open(conf.GetConfig().GetDatabase("sqlite")), &gorm.Config{})
	if err != nil {
		cpath, _ := os.Getwd()
		panic("failed to connect database (current path:" + cpath + ")")
	}
}
