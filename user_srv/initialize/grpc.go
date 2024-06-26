package initialize

import (
	"fmt"
	"log"
	"net"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/health"
	"google.golang.org/grpc/health/grpc_health_v1"
	"volcano.user_srv/config"
	"volcano.user_srv/utils"
)

func RegisterGrpc() error {
	server := grpc.NewServer()
	serverConf := config.BaseMapConfInstance.Base
	log.Printf("serverConf %+v", serverConf)

	
	// proto.RegisterUserServer(server, &handler.UserServer{})

	lis, err := net.Listen("tcp", fmt.Sprintf("%s:%d", serverConf.Addr, utils.GetFreePort()))
	if err != nil {
		zap.S().Errorf("initialize.grpc.RegisterGrpc.Listen: %v", err)
		return err
	}

	grpc_health_v1.RegisterHealthServer(server, health.NewServer())


	go func() {
    fmt.Println(`
 _     ____  _     ____  ____  _      ____ 
/ \ |\/  _ \/ \   /   _\/  _ \/ \  /|/  _ \
| | //| / \|| |   |  /  | / \|| |\ ||| / \|
| \// | \_/|| |_/\|  \_ | |-||| | \||| \_/|
\__/  \____/\____/\____/\_/ \|\_/  \|\____/
                                           `)

		zap.S().Infof("GRPC Server listening on %s:%d", serverConf.Addr, utils.GetFreePort())
		err = server.Serve(lis)
		if err != nil {
			zap.S().Errorf("initialize.grpc.RegisterGrpc.Serve: %v", err)
			panic(err)
		}
	}()

	return nil
}