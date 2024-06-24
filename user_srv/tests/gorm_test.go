package tests

import (
	"log"
	"testing"

	"volcano.user_srv/config"
	"volcano.user_srv/lib"
)

func TestGorm(t *testing.T) {
	gormLib := lib.NewMysql(config.MysqlConfig{
		Host: "127.0.0.1",
		Port: 3307,
		Name: "volcano_user_srv",
		User: "root",
		Password: "123456",
		MaxOpenConn: 10,
		MaxIdleConn: 20,
		MaxCoonLifeTime: 100,
		TimeLocation: "Asia%2FShanghai",
	})
	
	if err := gormLib.Init(); err != nil {
		log.Fatal(err)
	}
}