package main

import (
	"fmt"
	"github.com/Gouplook/dzgin"
	"github.com/smallnest/rpcx/server"
	"rpcxA/routers"
)

// <your_token>：刚刚生成的token
// <REPO>：要访问的仓库名称，当前仓库的名称
// git remote set-url origin  https://<your_token>@github.com/<USERNAME>/<REPO>.git
// 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

func main() {

	// 添加日志系统

	fmt.Println("service .....==")
	//启动服务
	rpcServer := server.NewServer()
	routers.InitRpcRouters(rpcServer)

	//启动链路追踪
	//address := "0.0.0.0:9001"
	address := fmt.Sprintf("%v:%v", dzgin.AppConfig.String("rpchost"), dzgin.AppConfig.String("rpcport"))
	if err := rpcServer.Serve("tcp", address); err != nil {
		// 打印日志
		fmt.Println("rpcServer fail")
	}

}
