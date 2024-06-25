package main

import (
	"os"
	"os/signal"
	"syscall"

	"volcano.user_srv/cmd"
	"volcano.user_srv/initialize"
	"volcano.user_srv/lib"
)


func main() {
	// Get command line arguments
	if err := cmd.CmdExecute(); err != nil {
		panic(err)
	}

	// Init modules such as gorm | viper | nacos | zap
	if err := initialize.InitModules(cmd.GetCmdMode(), cmd.GetCmdConf()); err != nil {
    panic(err)
	}

	// Using consul for service registration
	consulLib := lib.NewConsul()
	if err := consulLib.Register(); err != nil {
		panic(err)
	}

	if err := initialize.RegisterGrpc(); err != nil {
		panic(err)
	}
	
	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	consulLib.DeRegister()
}