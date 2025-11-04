package basic

import (
	"fmt"
	"sync"
	"sync/atomic"
)

func Sync2() {
	var count int64       // 共享计数器（必须使用int64类型，原子操作函数要求）
	var wg sync.WaitGroup // 用于等待所有协程完成

	// 启动10个协程
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()

			// 每个协程执行1000次递增
			for j := 0; j < 1000; j++ {
				// 原子递增操作：将count的值加1，返回新值（此处不关心返回值）
				atomic.AddInt64(&count, 1)
			}
		}()
	}

	wg.Wait() // 等待所有协程完成
	fmt.Printf("最终计数器的值: %d\n", count)
}
