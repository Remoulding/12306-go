package main

import (
	"github.com/Remoulding/12306-go/cronJob/script"
	"github.com/Remoulding/12306-go/ticket-service/configs"
	"github.com/robfig/cron/v3"
	"log"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	c := cron.New()
	configs.InitDBInstance()
	script.InitData()
	// 每天 23:50 执行
	_, err := c.AddFunc("50 23 * * *", script.InitData)
	if err != nil {
		log.Fatal("添加定时任务失败:", err)
	}
	// 启动 cron 调度器
	c.Start()
	select {} // 阻塞主 goroutine，保持程序运行
}
