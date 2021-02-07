package main

import (
	"fmt"
	"time"

	"github.com/robfig/cron/v3"
)

func CronTask(){
	fmt.Println(time.Now().Format("2006-01-02 15:04:05"))
}

func main(){
	c := cron.New()
    c.AddFunc("10 * * * * *", CronTask)  //2 * * * * *, 2 表示每分钟的第2s执行一次
    c.Start()
	defer c.Stop()
	for {
		time.After(time.Second * 10)
	}
}