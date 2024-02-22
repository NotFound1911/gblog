package main

import (
	"github.com/NotFound1911/mserver"
	"net/http"
)

type user struct {
}

func (u *user) transfer(ctx *mserver.Context) error {
	// 服务发现
	// todo grpc 调用user服务
	ctx.SetStatus(http.StatusOK).Json("login")
	return nil
}

// 网关 调用转发http请求
func main() {
	core := mserver.NewCore()
	u := new(user)
	userGroup := core.Group("/user")
	userGroup.Any("/login", u.transfer)
	core.Start(":8888") // 监听8888端口
}
