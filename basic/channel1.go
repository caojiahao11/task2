package basic

import (
	"fmt"
	"sync"
)

// 使用带缓冲通道的版本
func Channel1() {
	fmt.Println("=== 带缓冲通道版本 ===\n")

	// 创建缓冲大小为3的通道
	ch := make(chan int, 3)
	var wg sync.WaitGroup

	// 生产者协程
	wg.Add(1)
	go func() {
		defer wg.Done()
		defer close(ch)

		fmt.Println("生产者开始工作...")
		for i := 1; i <= 10; i++ {
			fmt.Printf("生产者发送: %d (通道长度: %d/%d)\n", i, len(ch), cap(ch))
			ch <- i
		}
		fmt.Println("生产者完成")
	}()

	// 消费者协程
	wg.Add(1)
	go func() {
		defer wg.Done()

		fmt.Println("消费者开始工作...")
		for num := range ch {
			fmt.Printf("消费者接收: %d (通道长度: %d/%d)\n", num, len(ch), cap(ch))
		}
		fmt.Println("消费者完成")
	}()

	wg.Wait()
	fmt.Println("\n=== 程序执行完毕 ===")
}
