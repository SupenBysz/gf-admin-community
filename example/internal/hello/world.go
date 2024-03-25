package hello

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
)

var MYHelloWorld = myHelloWorld{}

type myHelloWorld struct{}

type MyHelloWorldReq struct {
	g.Meta `path:"/myHelloWorld" method:"get" summary:"HELLOWORLD" tags:"HelloWorld"`
	Id     int64 `json:"id"             dc:"ID"`
}

type MyHelloWorldRes struct {
	Say string `json:"say" dc:""`
}

// SayHelloWorld 测试接口
func (c *myHelloWorld) SayHelloWorld(ctx context.Context, req *MyHelloWorldReq) (*MyHelloWorldRes, error) {
	ret := &MyHelloWorldRes{}
	ret.Say = "Hello World"

	// ....

	return ret, nil
}
