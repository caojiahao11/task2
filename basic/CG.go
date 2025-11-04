package basic

import (
	"fmt"
	"sync"
	"time"
)

type Task func()

// 任务执行结果
type TaskResult struct {
	TaskName string
	Duration time.Duration
	Error    error
}

// 任务调度器
type TaskScheduler struct {
	tasks     []Task
	taskNames []string
	results   chan TaskResult
	wg        sync.WaitGroup
}

func NewTaskScheduler() *TaskScheduler {
	return &TaskScheduler{
		results: make(chan TaskResult),
	}
}

// 添加任务
func (ts *TaskScheduler) AddTask(name string, task Task) {
	ts.tasks = append(ts.tasks, task)
	ts.taskNames = append(ts.taskNames, name)
}

// 执行单个任务
func (ts *TaskScheduler) executeTask(taskName string, task Task) {
	defer ts.wg.Done()

	start := time.Now()

	// 执行任务
	task()

	duration := time.Since(start)

	// 发送结果到Channel
	ts.results <- TaskResult{
		TaskName: taskName,
		Duration: duration,
		Error:    nil,
	}
}

// 启动调度器
func (ts *TaskScheduler) Start() {
	// 启动结果收集器
	go ts.collectResults()

	// 并发执行所有任务
	for i, task := range ts.tasks {
		ts.wg.Add(1)
		go ts.executeTask(ts.taskNames[i], task)
	}

	// 等待所有任务完成
	ts.wg.Wait()
	close(ts.results) // 关闭结果Channel，通知收集器结束
}

// 收集任务执行结果
func (ts *TaskScheduler) collectResults() {
	totalStart := time.Now()
	taskCount := 0

	fmt.Println("🚀 任务调度器启动...")
	fmt.Println("==================================")

	// 从Channel接收结果
	for result := range ts.results {
		taskCount++
		fmt.Printf("✅ 任务完成: %s\n", result.TaskName)
		fmt.Printf("   执行时间: %v\n", result.Duration)

		if result.Error != nil {
			fmt.Printf("   ❌ 错误: %v\n", result.Error)
		}
		fmt.Println("----------------------------------")
	}

	totalDuration := time.Since(totalStart)
	fmt.Printf("🎉 所有任务执行完成!\n")
	fmt.Printf("   总任务数: %d\n", taskCount)
	fmt.Printf("   总执行时间: %v\n", totalDuration)
}

func CG() {
	// 创建调度器
	scheduler := NewTaskScheduler()

	// 添加各种任务
	scheduler.AddTask("快速计算任务", func() {
		time.Sleep(100 * time.Millisecond)
		sum := 0
		for i := 0; i < 1000000; i++ {
			sum += i
		}
		fmt.Println(sum)
	})

	scheduler.AddTask("网络请求模拟", func() {
		time.Sleep(500 * time.Millisecond)
		fmt.Println("   网络请求完成")
	})

	scheduler.AddTask("文件处理任务", func() {
		time.Sleep(300 * time.Millisecond)
		fmt.Println("   文件处理完成")
	})

	// 启动调度器
	scheduler.Start()
}
