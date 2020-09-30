package main

import (
	// "fmt"
	"github.com/gin-gonic/gin"
)

// 代理客户端类
type ServiceProxyClient struct {
	Pid				string
	Type			string
	ServicesConfig	interface{}
	ServicesIDs		[]int64
	Url				string
}

// 获取服务相关配置
func (client *ServiceProxyClient) GetServicesConfigMap() map[string]interface{} {
	result := make(map[string]interface{})
	return result
}

// 获取请求参数
func request_data(c *gin.Context) map[string]string {
	result := make(map[string]string)
	return result
}


// 新建一个代理客户端
func CreateClient(pid string) *ServiceProxyClient {
	client := ServiceProxyClient{
		Pid: pid,
	}
	return &client
}
