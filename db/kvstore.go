package db

import (
	"github.com/jz0ojiang/go-maimai-player-counter/conf"
	"github.com/syndtr/goleveldb/leveldb"
)

var LevelDB *leveldb.DB

func init() {
	db, err := leveldb.OpenFile(conf.GetConfig().GetDatabase("leveldb"), nil)
	if err != nil {
		panic(err)
	}
	LevelDB = db
}
