package lib

import (
	"fmt"

	"github.com/hashicorp/consul/api"
	uuid "github.com/satori/go.uuid"

	"go.uber.org/zap"
	"volcano.user_srv/cmd"
	"volcano.user_srv/config"
)


type Consul struct {
	serviceId string
	client *api.Client
}

func NewConsul() *Consul {
	return &Consul{}
}

func (c *Consul) SetServiceId(serviceId string) {
	c.serviceId = serviceId
}

func (c *Consul) SetClient(client *api.Client) {
	c.client = client
}

func (c *Consul) Register() error {
	consulClientConf := api.DefaultConfig()
  serverConf := config.BaseMapConfInstance.Base
	consulConf := config.ServerConfInstance.ConsulInfo
	consulClientConf.Address = fmt.Sprintf("%s:%d", consulConf.Host, consulConf.Port)

	client, err := api.NewClient(consulClientConf)
	if err != nil {
		zap.S().Errorf("lib.consul.Register.NewClient error: %s", err.Error())
		return err
	}
	
	check := &api.AgentServiceCheck{
		GRPC:                           fmt.Sprintf("%s:%d", serverConf.Addr,cmd.GetCmdPort()),
		Timeout:                        "5s",
		Interval:                       "5s",
		DeregisterCriticalServiceAfter: "15s",
	}

	registration := new(api.AgentServiceRegistration)
	registration.Name = config.ServerConfInstance.Name
	serviceID := uuid.NewV4().String()
	c.client = client
	registration.ID = serviceID
	registration.Port = cmd.GetCmdPort()
	registration.Tags = consulConf.Tags
	registration.Address = serverConf.Addr
	registration.Check = check
	err = client.Agent().ServiceRegister(registration)
	if err != nil {
		zap.S().Errorf("lib.consul.Register.NewClient error: %s", err.Error())
		return err
	}

	c.SetClient(client)
	c.SetServiceId(serviceID)

	return nil
} 


func (c *Consul) DeRegister()  {
	if err := c.client.Agent().ServiceDeregister(c.serviceId); err != nil{
		zap.S().Errorf("Consul DeRegister Failed: %v", err)
	} else {
		zap.S().Info("Consul DeRegister Successfully~")
	}
}