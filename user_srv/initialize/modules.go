package initialize

import (
	"errors"
	"fmt"

	"go.uber.org/zap"
	"volcano.user_srv/config"
	"volcano.user_srv/lib"
)

// const InitModuleErrorMsg =
const ServerUnmarshalCaption = "Server 基础配置"
const NacosUnmarshalCaption = "Nacos 注册中心配置"
const ZapUnmarshalCaption = "Zap 日志配置"

const ServerUnmarshalFile = "base.toml"
const NacosUnmarshalFile = "nacos.toml"
const ZapUnmarshalFile = "zap.toml"
var ErrorInitModule error = errors.New("Initialize.InitModules failed")


func InitModules(mode string, confPath string) error {
	zap.L().Info("------------------------------")
	zap.S().Infof("[INFO]  config=%s\n", confPath)


	if confPath == "" {
		zap.S().Errorf("initialize.InitModules: Please specify a config path like %s", "./conf/env")
		return ErrorInitModule
	}


	return initModules(mode, confPath,)
}

func initModules(mode string, confPath string) error {
  viperLib := lib.NewViper()
	viperLib.Init()

	// unmarshal server config
	baseConfPath := fmt.Sprintf("%s/%s/%s", confPath, mode, ServerUnmarshalFile)
	if err := viperLib.Unmarshal(baseConfPath, config.BaseMapConfInstance, ServerUnmarshalCaption); err != nil {
		return err
	}

	// unmarshal nacos config
	nacosConfPath := fmt.Sprintf("%s/%s/%s", confPath, mode, NacosUnmarshalFile)
	if err := viperLib.Unmarshal(nacosConfPath, config.NacoMapConfInstance, NacosUnmarshalCaption); err != nil {
		return err
	}

	// Init Nacos
	nacosLib := lib.NewNacos(config.NacoMapConfInstance.Base)
	if err := nacosLib.Init(); err != nil {
		return err
	}

	// Init Zap
	zapLib := lib.NewZap(config.ServerConfInstance.ZapInfo)
	if err := zapLib.Init(); err != nil {
		return err
	}

	// Init Gorm
	gormLib := lib.NewMysql(config.ServerConfInstance.MysqlInfo)
	
	if err := gormLib.Init(); err != nil {
		return err
	}
	
	return nil
}