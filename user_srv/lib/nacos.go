package lib

import (
	"encoding/json"
	"errors"
	"log"

	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"gopkg.in/natefinch/lumberjack.v2"
	"volcano.user_srv/config"
)

type Nacos struct {
	conf config.NacosConf
}

func NewNacos(conf config.NacosConf) *Nacos {
	return &Nacos{conf}
}

var ErrorNacosInit = errors.New("nacos error initializing")

func (n *Nacos) Init() error {
	c := n.conf
	if c.Host == "" || c.Port == 0 {
		log.Printf("[ERROR] NacoMapConfInstance host and port must be specified")
		return ErrorNacosInit
	}

	serverConf := []constant.ServerConfig{
		{
			IpAddr: c.Host,
			Port:  c.Port,
		},
	}

	clientConf := constant.ClientConfig {
		NamespaceId:        c.Namespace, 
		TimeoutMs:           c.Timeout,
		NotLoadCacheAtStart: c.NotLoadCacheAtStart,
		LogDir:             c.LogDir,
		CacheDir:            c.CacheDir,
		LogLevel:            c.LogLevel,
		LogRollingConfig: &lumberjack.Logger{
		  MaxAge: c.MaxAge,
			MaxBackups: c.MaxBackUp,
		},
	}

	configClient, err := clients.CreateConfigClient(map[string]interface{}{
		"serverConfigs": serverConf,
		"clientConfig":  clientConf,
	})

	if err != nil {
		log.Printf("[ERROR] lib.nacos.Init.CreateConfigClient: %s", err.Error())
		return err
	}

	content, err := configClient.GetConfig(vo.ConfigParam{
		DataId: c.DataId,
		Group:  c.Group})

		if err != nil {
			log.Printf("[ERROR] lib.nacos.Init.GetConfig: %s", err.Error())
			return err
		}

	if err := json.Unmarshal([]byte(content), &config.ServerConfInstance); err != nil {
		log.Printf("[ERROR] lib.nacos.Init.Unmarshal: %s", err.Error())
	}

	log.Printf("[INFO] Server config from nacos %+v", config.ServerConfInstance)
	return nil
}
