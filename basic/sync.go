package basic

import (
	"fmt"
	"sync"
)

func Sync1() {
	var count int         // 共享的计数器
	var mutex sync.Mutex  // 用于保护计数器的互斥锁
	var wg sync.WaitGroup // 用于等待所有协程完成

	// 启动10个协程
	for i := 0; i < 10; i++ {
		wg.Add(1) // 每启动一个协程，WaitGroup计数+1
		go func() {
			defer wg.Done() // 协程结束时，WaitGroup计数-1

			// 每个协程对计数器进行1000次递增
			for j := 0; j < 1000; j++ {
				mutex.Lock()   // 加锁：独占访问共享资源
				count++        // 临界区操作：修改计数器
				mutex.Unlock() // 解锁：释放对共享资源的独占
			}
		}()
	}

	wg.Wait() // 等待所有10个协程完成
	fmt.Printf("最终计数器的值: %d\n", count)
}
