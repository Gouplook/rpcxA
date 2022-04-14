package main

import (
	"fmt"
	"github.com/smallnest/rpcx/server"
	"rpcxA/routers"
)

func main() {

	//启动服务
	rpcServer := server.NewServer()
	routers.InitRpcRouters(rpcServer)

	//启动链路追踪

	address := "0.0.0.0:9001"
	if err := rpcServer.Serve("tcp", address); err != nil {
		// 打印日志
		fmt.Println("rpcServer fail")
	}

}
