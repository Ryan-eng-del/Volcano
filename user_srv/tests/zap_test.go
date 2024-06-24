package tests

import (
	"testing"

	"go.uber.org/zap"
	"volcano.user_srv/config"
	"volcano.user_srv/lib"
)



func TestZap(t *testing.T) {
	zapLib := lib.NewZap(config.ZapConf{
		MaxSize: 10,
		MaxAge: 24,
		MaxBackups: 3,
		DebugFileName: "/Users/max/Documents/coding/Backend/Golang/Personal/volcano/user_srv/log/volcano.debug.log",
		ErrorFileName: "/Users/max/Documents/coding/Backend/Golang/Personal/volcano/user_srv/log/volcano.error.log",
		InfoFileName: "/Users/max/Documents/coding/Backend/Golang/Personal/volcano/user_srv/log/volcano.info.log",
	})
	
	zapLib.Init()
	zap.S().Info("testing")
	zap.S().Error("testing")
	zap.S().Warn("testing")
	zap.S().Debug("testing")
}