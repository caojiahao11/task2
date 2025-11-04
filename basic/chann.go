package basic

import (
	"fmt"
	"sync"
)

func PrintWithchannel() {
	// 创建两个Channel用于控制顺序
	oddChan := make(chan bool)  // 控制奇数协程
	evenChan := make(chan bool) // 控制偶数协程
	var wg sync.WaitGroup
	wg.Add(2)

	// 启动奇数协程
	go func() {
		defer wg.Done() // 修复：添加 Done()
		for i := 1; i <= 10; i += 2 {
			<-oddChan // 等待打印许可
			fmt.Printf("奇数: %d\n", i)
			if i < 10 { //
				evenChan <- true // 通知偶数协程打印
			}
		}
	}()

	// 启动偶数协程
	go func() {
		defer wg.Done() // 修复：添加 Done()
		for i := 2; i <= 10; i += 2 {
			<-evenChan // 等待打印许可
			fmt.Printf("偶数: %d\n", i)
			if i < 10 {
				oddChan <- true // 通知奇数协程打印
			}
		}
	}()

	// 启动打印序列
	oddChan <- true

	// 等待打印完成
	wg.Wait() // 等待两个协程完成
	fmt.Println("所有打印完成！")
}
