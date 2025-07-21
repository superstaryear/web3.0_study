package main

import (
	"fmt"
	"sync"
	"time"
)

func receive(desc string, ch <-chan int) {
	for i := range ch {
		fmt.Printf("%s接收到值为%d\n", desc, i)
	}
}

func send(maxNum int, desc string, ch chan<- int) {
	for i := 1; i <= maxNum; i++ {
		ch <- i
		fmt.Printf("%s发送值%d\n", desc, i)
	}
	close(ch)
}
func main() {
	//创建一个无缓冲区的整数通道
	var ch10 = make(chan int)
	//创建waitGroup用于等待协程完成
	var wg sync.WaitGroup
	wg.Add(2)
	go func() {
		defer wg.Done()
		send(10, "10个整数任务：", ch10)
	}()
	go func() {
		defer wg.Done()
		receive("10个整数任务：", ch10)
	}()
	//等待所有协程完成
	wg.Wait()
	fmt.Println("\n------分割线-------\n")
	//创建一个有缓冲区的通道
	var ch100 = make(chan int, 50)
	go send(100, "100个整数任务:", ch100)
	go receive("100个整数任务", ch100)
	time.Sleep(2 * time.Second)
}
