package main

import (
	"fmt"
	"time"

	"github.com/Shopify/sarama"
)

func main() {
	ticker := time.Tick(time.Minute * 10) //定义一个1秒间隔的定时器
	var i int
	i = 0
	fmt.Println("任务开始")
	for _ = range ticker {
		kafka_producer(i)
		i ++
	}
}

func kafka_producer(a int){
	config := sarama.NewConfig()
	config.Producer.RequiredAcks = sarama.WaitForAll          // 发送完数据需要leader和follow都确认
	config.Producer.Partitioner = sarama.NewRandomPartitioner // 新选出一个partition
	config.Producer.Return.Successes = true                   // 成功交付的消息将在success channel返回

	// 构造一个消息
	msg := &sarama.ProducerMessage{}
	msg.Topic = "test"
	msg.Value = sarama.StringEncoder("this is a test log")
	// 连接kafka
	client, err := sarama.NewSyncProducer([]string{"127.0.0.1:9092"}, config)
	fmt.Println("连接kafka")
	if err != nil {
		fmt.Println("producer closed, err:", err)
		return
	}
	defer client.Close()
	fmt.Println("发送消息")
	// 发送消息
	for i := 10000 * a; i < 10000 * (a +1); i++ {
		msg.Value = sarama.StringEncoder(fmt.Sprintf("this is a test log%d", i))
		pid, offset, err := client.SendMessage(msg)
		if err != nil {
			fmt.Println("send msg failed, err:", err)
			return
		}
		fmt.Printf("pid:%v offset:%v\n", pid, offset)
	}
}
