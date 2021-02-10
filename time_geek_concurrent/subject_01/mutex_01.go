package main

import (
	"fmt"
	"sync"
)

/**
我们创建了 10 个 goroutine，同时不断地对一个变量（count）进行加 1 操作，
每个 goroutine 负责执行 10 万次的加 1 操作，
我们期望的最后计数的结果是 10 * 100000 = 1000000 (一百万)
*/

func main() {
	CounterA()
	// 输出结果: 250475
	// 每次结果都不一样

	CounterB()
	// 输出结果： 1000000

	CounterC()
	// 输出结果： 1000000
}

// CounterA name
func CounterA() {
	var count = 0 // 使用WaitGroup等待10个goroutine完成
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done() // 对变量count执行10次加1
			for j := 0; j < 100000; j++ {
				count++
			}
		}()
	}

	// 等待10个goroutine完成
	wg.Wait()
	fmt.Println(count)
}

func CounterB() {
	// 互斥锁
	var mu sync.Mutex
	// 计数值
	var count = 0

	// 辅助变量，用来确认所有的goroutine都完成
	var wg sync.WaitGroup
	wg.Add(10)
	// 启动10个gourontine
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			// 累加10万次
			for j := 0; j < 100000; j++ {
				mu.Lock()
				count++
				mu.Unlock()
			}
		}()
	}
	// 等待10个goroutine完成
	wg.Wait()
	fmt.Println(count)
}

// 查看data race 情况
// go tool compile -race -S subject_01_mutex.go | grep racefuncenter

type Counter struct {
	mu    sync.Mutex
	Count uint64
}

func CounterC() {
	var counter Counter
	var wg sync.WaitGroup
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 100000; j++ {
				counter.mu.Lock()
				counter.Count++
				counter.mu.Unlock()
			}
		}()
	}

	// 等待10个goroutine完成
	wg.Wait()
	fmt.Println(counter.Count)
}
