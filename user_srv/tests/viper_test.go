package tests

import (
	"log"
	"testing"

	"volcano.user_srv/config"
	"volcano.user_srv/lib"
)


func TestViper(t *testing.T) {
	nacosConf := config.NacoMapConfInstance
	viper := lib.NewViper()
	viper.Init()
	if err := viper.Unmarshal("/Users/max/Documents/coding/Backend/Golang/Personal/volcano/user_srv/config/dev/nacos.toml", &nacosConf, "nacos 配置中心"); err != nil {
		t.Fatal(err)
	}

	log.Printf("config: %+v", nacosConf)
}