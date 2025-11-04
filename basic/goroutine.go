package basic

import (
	"fmt"
	"sync"
	"time"
)

func printOdd() {
	for i := 1; i < 10; i += 2 {
		fmt.Print(i, ", ")
	}
	fmt.Println()
}

func printEven() {
	for i := 0; i < 10; i += 2 {
		fmt.Print(i, ", ")
	}
	fmt.Println()
}

func goPrint() {
	go printOdd()
	go printEven()
	time.Sleep(2 * time.Second)
}

func Goroutine() {

	fmt.Println("=== 奇偶数打印协程演示 ===")

	// 使用 WaitGroup 等待两个协程完成
	var wg sync.WaitGroup
	wg.Add(2) // 需要等待2个协程

	// 启动打印奇数的协程
	go func() {
		defer wg.Done() // 协程结束时通知WaitGroup
		fmt.Println("奇数协程启动...")
		for i := 1; i <= 10; i += 2 {
			fmt.Printf("奇数: %d\n", i)
			time.Sleep(100 * time.Millisecond) // 模拟工作负载，让输出更清晰
		}
		fmt.Println("奇数协程结束")
	}()

	// 启动打印偶数的协程
	go func() {
		defer wg.Done() // 协程结束时通知WaitGroup
		fmt.Println("偶数协程启动...")
		for i := 2; i <= 10; i += 2 {
			fmt.Printf("偶数: %d\n", i)
			time.Sleep(150 * time.Millisecond) // 不同的延迟，展示并发特性
		}
		fmt.Println("偶数协程结束")
	}()

	fmt.Println("主协程: 已启动两个子协程，等待它们完成...")

}
