package main

import (
	"fmt"
	"sync"
	"time"
)

// 打印奇数
func printJiShu(wg *sync.WaitGroup) {
	defer wg.Done() // 协程结束时通知WaitGroup
	for i := 1; i <= 10; i++ {
		if i%2 != 0 {
			fmt.Println("基数打印", i)
		}
	}
}

// 打印偶数
func printOuShu(wg *sync.WaitGroup) {
	defer wg.Done() //协程结束时通知WaitGroup
	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println("偶数打印", i)
		}
	}
}

type task func()

type taskResult struct {
	index    int
	duration time.Duration
}

type taskDemo struct{}

// 任务调度器
func (td *taskDemo) schedule(tasks []task) []taskResult {
	var wg sync.WaitGroup
	results := make([]taskResult, len(tasks))
	for i, t := range tasks {
		wg.Add(1)
		go func(idx int, t task) {
			defer wg.Done()
			start := time.Now()
			t()
			elapsed := time.Since(start)
			results[idx] = taskResult{idx, elapsed}

		}(i, t)
	}
	wg.Wait()
	return results
}

func main() {
	fmt.Println("-----所有数字打印开始-----")
	var wg sync.WaitGroup
	//我们有两个协程需要等待
	wg.Add(2)
	go printJiShu(&wg)
	go printOuShu(&wg)
	wg.Wait()
	fmt.Println("-----所有数字打印完成-----")

	//设计一个任务调度器，接收一组任务（可以用函数表示），并使用协程并发执行这些任务，同时统计每个任务的执行时间。
	demo := &taskDemo{}
	tasks := []task{
		func() {
			time.Sleep(500 * time.Millisecond)
			fmt.Println("任务1完成")
		},
		func() {
			time.Sleep(100 * time.Millisecond)
			fmt.Println("任务2完成")
		},
		func() {
			time.Sleep(700 * time.Millisecond)
			fmt.Println("任务3完成")

		},
	}
	fmt.Println("\n-----任务调度开始-----")
	results := demo.schedule(tasks)
	for _, result := range results {
		fmt.Printf("任务%d耗时：%v\n", result.index+1, result.duration)
	}
	fmt.Println("-----任务调度完成-----")

}
