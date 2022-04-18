package service

import (
	"context"
	"github.com/Gouplook/rpcxinterfaxe/interface/demoA"
	"rpcxA/logics"
)

type DemoA struct {
}

func (d *DemoA) Add(ctx context.Context, args *demoA.ArgsAdd, reply *demoA.ReplyAdd) error {
	logic := new(logics.DemoALogic)
	err := logic.Add(ctx, args, reply)
	if err != nil {
		panic(err)
	}
	return nil
}
