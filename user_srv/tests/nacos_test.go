package tests

import (
	"testing"

	"volcano.user_srv/config"
	"volcano.user_srv/lib"
)


func TestNacos(t *testing.T) {
	nacosLib := lib.NewNacos(config.NacosConf{
		Host:  "0.0.0.0",
		Port: 8848,
		Namespace: "5a9dd777-c218-4532-a308-a3eddbec1c5a",
		User: "nacos",
		Password: "nacos",
		DataId: "user_srv_dev.json",
		Group: "DEFAULT_GROUP",
		LogDir: "/Users/max/Documents/coding/Backend/Golang/Personal/volcano/user_srv/nacos_dir/nacos/log",
		CacheDir: "/Users/max/Documents/coding/Backend/Golang/Personal/volcano/user_srv/nacos_dir/nacos/cache",
		MaxAge: 24,
		LogLevel: "debug",
		RotateTime: "1h",
		Timeout: 5000,
		NotLoadCacheAtStart: true,
		MaxBackUp: 3,
	})
	nacosLib.Init()
	// log.Printf("server config instance %+v", config.ServerConfInstance)
}