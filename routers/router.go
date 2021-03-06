package routers

import (
	"fmt"
	"github.com/smallnest/rpcx/server"
	"rpcxA/service"
)

//搜索注册路由
func InitRpcRouters(rpcServer *server.Server) {

	// API段访问需要注册路由
	//err := rpcServer.RegisterName("DemoA", new(service.DemoA), "")
	err := rpcServer.Register(new(service.DemoA), "")
	if err != nil {
		fmt.Println("failed to register rpcRouter: %v", err)
	}
}
