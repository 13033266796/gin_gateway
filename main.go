package main

import (
	"fmt"
	_ "io/ioutil"
	"github.com/gin-gonic/gin"

	_ "github.com/jinzhu/gorm/dialects/mysql"

	"gin_gateway/pkg/model"
	"gin_gateway/pkg/conf"
)

var welcome string

func init() {
	welcome = "hello, welcome my gin"
	conf.SetupConf()
	model.SetupDB()
}

func main() {
	engine := gin.Default()
	// 需要转发的地址
	proxy := engine.Group("/robot/v1")
	proxy.POST("/:type/*dest", func(context *gin.Context) {

		// 获取url上需要的地址
		_type := context.Param("type")
		dest := context.Param("dest")
		fmt.Println(_type)
		fmt.Println(dest)

		// 获取请求参数
		buf := make([]byte, 1024)
		n, _ := context.Request.Body.Read(buf)
		// data, _ := ioutil.ReadAll(context.Request.Body)
		fmt.Println(string(buf[:n]))
		context.Writer.Write([]byte(dest))
	})

	fmt.Println(welcome)
	engine.Run(":8090")
}