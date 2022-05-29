package main

import (
	"GinLearning/pkg/setting"
	"GinLearning/routers"
	"fmt"
	"net/http"
)

func main() {
	// 路由
	router := routers.InitRouter()
	s := &http.Server{
		Addr:           fmt.Sprintf(":%d", setting.HTTPPort), // 监听的端口号
		Handler:        router,                               // http 句柄 实质为 ServerHTTP，用于处理程序响应HTTP请求
		ReadTimeout:    setting.ReadTimeout,                  // 允许读取的最大时间
		WriteTimeout:   setting.WriteTimeout,                 // 允许写入的最大时间
		MaxHeaderBytes: 1 << 20,                              // 请求头的最大字节数
	}
	// 监听服务服务 监听TCP网路地址、ADDR 和调用应用程序处理连接上的请求
	s.ListenAndServe()
}
