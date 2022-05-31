package main

import (
	"GinLearning/pkg/setting"
	"GinLearning/routers"
	"fmt"
	"github.com/fvbock/endless"
	"log"
	"syscall"
)

func main() {

	endless.DefaultReadTimeOut = setting.ReadTimeout
	endless.DefaultWriteTimeOut = setting.WriteTimeout
	endless.DefaultMaxHeaderBytes = 1 << 20
	endPoint := fmt.Sprintf(":%d", setting.HTTPPort)

	// endless.NewServer 返回一个初始化的 endlessServer 对象
	server := endless.NewServer(endPoint, routers.InitRouter())
	// server.BeforeBegin 时输出当前进程的 pid
	server.BeforeBegin = func(add string) {
		log.Printf("Actual pid is %d", syscall.Getpid())
	}
	// 调用 server.ListenAndServe 将实际“启动”服务
	err := server.ListenAndServe()
	if err != nil {
		log.Printf("Server err: %v", err)
	}
}
