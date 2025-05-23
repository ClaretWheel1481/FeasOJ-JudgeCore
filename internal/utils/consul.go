package utils

import (
	"fmt"
	"log"
	"main/internal/config"

	"github.com/hashicorp/consul/api"
)

// 服务注册示例
func RegService(client *api.Client) error {
	agent := client.Agent()

	var protocol string
	if config.EnableHTTPS {
		protocol = "https"
	} else {
		protocol = "http"
	}

	registration := &api.AgentServiceRegistration{
		ID:   config.ConsulServiceID,   // 服务唯一ID
		Name: config.ConsulServiceName, // 服务名称
		Port: config.ServicePort,       // 服务端口
		Tags: []string{"gin", "judge"},
		Check: &api.AgentServiceCheck{
			HTTP:     fmt.Sprintf("%s://%s:%d/api/v1/judgecore/health", protocol, config.ServiceAddress, config.ServicePort), // 健康检查地址
			Interval: "60s",
			Timeout:  "6s",
		},
	}

	err := agent.ServiceRegister(registration)
	if err != nil {
		log.Println("[FeasOJ] JudgeCore service registration failed:", err)
		return err
	}
	log.Println("[FeasOJ] JudgeCore service registered successfully")
	return nil
}
