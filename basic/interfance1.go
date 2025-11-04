package basic

import (
	"fmt"
	"math"
)

// 1. 定义Shape接口
type Shape interface {
	Area() float64      // 计算面积
	Perimeter() float64 // 计算周长
}

// 2. 定义Rectangle结构体
type Rectangle struct {
	Width  float64
	Height float64
}

// Rectangle实现Shape接口的Area方法
func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Rectangle实现Shape接口的Perimeter方法
func (r Rectangle) Perimeter() float64 {
	return 2 * (r.Width + r.Height)
}

// 3. 定义Circle结构体
type Circle struct {
	Radius float64
}

// Circle实现Shape接口的Area方法
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Circle实现Shape接口的Perimeter方法
func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

// 4. 工具函数：打印形状信息
func printShapeInfo(s Shape) {
	fmt.Printf("面积: %.2f, 周长: %.2f\n", s.Area(), s.Perimeter())
}

func Interfance1() {
	// 创建Rectangle实例
	rect := Rectangle{Width: 5, Height: 3}

	// 创建Circle实例
	circle := Circle{Radius: 4}

	fmt.Println("=== 矩形信息 ===")
	fmt.Printf("矩形尺寸: %.1f x %.1f\n", rect.Width, rect.Height)
	printShapeInfo(rect)

	fmt.Println("\n=== 圆形信息 ===")
	fmt.Printf("圆形半径: %.1f\n", circle.Radius)
	printShapeInfo(circle)

	// 使用接口类型的切片
	fmt.Println("\n=== 使用接口切片 ===")
	shapes := []Shape{rect, circle}

	for i, shape := range shapes {
		fmt.Printf("形状 %d - ", i+1)
		printShapeInfo(shape)
	}
}
