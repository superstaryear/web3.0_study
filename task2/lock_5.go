package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

/*
使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值
*/
type safeCount struct {
	mu    sync.Mutex
	count int
}

func (s *safeCount) Inc() {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.count++
}

func (s *safeCount) getCount() int {
	s.mu.Lock()
	defer s.mu.Unlock()
	return s.count
}

/*
使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值
*/
type atomicCount struct {
	count int64
}

func (s *atomicCount) Inc() {
	atomic.AddInt64(&s.count, 1)
}

func (s *atomicCount) getCount() int64 {
	return atomic.LoadInt64(&s.count)
}

func main() {
	//编写一个程序，使用 sync.Mutex 来保护一个共享的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值。
	var safeCount safeCount = safeCount{sync.Mutex{}, 0}
	wg := sync.WaitGroup{}
	wg.Add(10)
	//启动10个协程
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			//每个协程对计数器进行1000次递增操作
			for i := 0; i < 1000; i++ {
				safeCount.Inc()
			}
		}()
	}
	wg.Wait()
	fmt.Printf("sync.Mutex 统计结果值为%d\n", safeCount.getCount())
	fmt.Println("\n------分割线-------\n")
	//使用原子操作（ sync/atomic 包）实现一个无锁的计数器。启动10个协程，每个协程对计数器进行1000次递增操作，最后输出计数器的值
	var atomicCount atomicCount = atomicCount{0}
	//启动10个协程
	for i := 0; i < 10; i++ {
		go func() {
			//每个协程对计数器进行1000次递增操作
			for i := 0; i < 1000; i++ {
				atomicCount.Inc()
			}
		}()
	}
	time.Sleep(2 * time.Second)
	fmt.Printf("atomicCount 统计结果值为%d\n", atomicCount.getCount())
}
