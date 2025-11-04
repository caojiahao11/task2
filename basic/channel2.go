package basic

import (
	"fmt"
	"sync"
	"time"
)

// 生产者协程：向缓冲通道发送100个整数
func producer(ch chan<- int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("🚀 生产者开始工作...")

	for i := 1; i <= 100; i++ {
		// 发送数据到缓冲通道
		ch <- i
		fmt.Printf("📤 生产者发送: %d (通道状态: %d/%d)\n", i, len(ch), cap(ch))

		// 模拟生产耗时
		time.Sleep(50 * time.Millisecond)
	}

	close(ch) // 关闭通道，表示没有更多数据
	fmt.Println("✅ 生产者完成，已关闭通道")
}

// 消费者协程：从缓冲通道接收整数并打印
func consumer(ch <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Println("🎯 消费者开始工作...")

	count := 0
	for num := range ch {
		count++
		fmt.Printf("📥 消费者接收: %d (通道状态: %d/%d) - 已处理: %d/100\n",
			num, len(ch), cap(ch), count)

		// 模拟消费耗时
		time.Sleep(80 * time.Millisecond)
	}

	fmt.Printf("✅ 消费者完成，总共处理了 %d 个数字\n", count)
}

func Channel2() {
	fmt.Println("=== 🎪 带缓冲通道演示程序 ===\n")

	// 创建缓冲大小为10的通道
	bufferSize := 10
	ch := make(chan int, bufferSize)

	fmt.Printf("📊 创建了缓冲大小为 %d 的通道\n\n", bufferSize)

	var wg sync.WaitGroup

	// 启动生产者协程
	wg.Add(1)
	go producer(ch, &wg)

	// 启动消费者协程
	wg.Add(1)
	go consumer(ch, &wg)

	// 等待所有协程完成
	wg.Wait()

	fmt.Println("\n=== 🎉 所有协程执行完毕 ===")
}
