package basic

import (
	"fmt"
	"sync"
	"time"
)

func RunTasks(tasks map[string]func()) {
	var wg sync.WaitGroup
	results := make(chan string, len(tasks)) // 使用Channel收集结果

	for name, task := range tasks {
		wg.Add(1)

		go func(taskName string, t func()) {
			defer wg.Done()

			start := time.Now()
			t()
			cost := time.Since(start)

			results <- fmt.Sprintf("%s: %v", taskName, cost)
		}(name, task)
	}

	// 等待所有任务完成
	wg.Wait()
	close(results)
	// 输出结果
	fmt.Println("执行结果:")
	for result := range results {
		fmt.Println(" ", result)
	}
}

func SimpleCg() {
	tasks := map[string]func(){
		"任务A": func() { time.Sleep(200 * time.Millisecond) },
		"任务B": func() { time.Sleep(100 * time.Millisecond) },
		"任务C": func() {
			time.Sleep(300 * time.Millisecond)
			fmt.Println("   任务C完成了一些工作")
		},
	}

	RunTasks(tasks)
}
