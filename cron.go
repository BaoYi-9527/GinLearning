package main

import (
	"GinLearning/models"
	"github.com/robfig/cron"
	"log"
	"time"
)

func main() {
	log.Println("Starting...")

	c := cron.New() // 创建一个新的 Cron job runner
	// AddFunc 会向 Cron job runner 添加一个 func 以给定的时间表运行
	c.AddFunc("* * * * * *", func() {
		log.Println("Run models,CleanAllTag...")
		models.CleanAllTag()
	})
	c.AddFunc("* * * * * *", func() {
		log.Println("Run models,CleanAllArticle...")
		models.CleanAllArticle()
	})
	// 在当前执行的程序中启动 Cron 调度程序
	// 主体是 goroutine + for + select + timer
	c.Start()
	// 创建一个新的定时器
	t1 := time.NewTimer(time.Second * 10)
	// for + select 阻塞 select 等待 channel
	for {
		select {
		case <-t1.C:
			// 重置定时器
			t1.Reset(time.Second * 10)
		}
	}
}
