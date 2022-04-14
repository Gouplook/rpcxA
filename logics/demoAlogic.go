package logics

import "context"

type DemoALogic struct {
}

func (d *DemoALogic) Add(ctx context.Context, a, b int, rs *int) error {
	// 逻辑代码，数据库读写
	//......
	*rs = a + b
	return nil
}
