package logics

import (
	"context"
	"github.com/Gouplook/rpcxinterfaxe/interface/demoA"
)

type DemoALogic struct {
}

func (d *DemoALogic) Add(ctx context.Context, args *demoA.ArgsAdd, reply *demoA.ReplyAdd) error {
	// 逻辑代码，数据库读写
	//......
	reply.Sum = args.Num_1 + args.Num_2
	return nil
}
