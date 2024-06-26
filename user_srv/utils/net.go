package utils

import (
	"log"
	"net"

	"volcano.user_srv/config"
)

func getFreePort() (int, error) {
	addr, err := net.ResolveTCPAddr("tcp", "localhost:0")
	if err != nil {
		return 0, err
	}

	l, err := net.ListenTCP("tcp", addr)
	if err != nil {
		return 0, err
	}
	defer l.Close()
	return l.Addr().(*net.TCPAddr).Port,  nil
}


func GetFreePort() int {
	if config.BaseMapConfInstance.Base.Port == 0 {
		 p, err := getFreePort()
		 if err != nil {
			 log.Println("[ERROR] GetCmdPort failed: ", err)
			 panic("0 is not a valid port number")
		 }
		 config.BaseMapConfInstance.Base.Port = p	
		 return p
	} 
	return config.BaseMapConfInstance.Base.Port
}