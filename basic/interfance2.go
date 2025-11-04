package basic

import "fmt"

// 1. 定义Person结构体
type Person struct {
	Name string
	Age  int
}

// Person的方法
func (p Person) Introduce() string {
	return fmt.Sprintf("大家好，我是%s，今年%d岁", p.Name, p.Age)
}

// 2. 定义Employee结构体，组合Person
type Employee struct {
	Person     // 匿名嵌入，实现组合
	EmployeeID string
	Department string
	Position   string
}

// Employee的PrintInfo方法
func (e Employee) PrintInfo() {
	fmt.Println("=== 员工信息 ===")
	fmt.Printf("员工ID: %s\n", e.EmployeeID)
	fmt.Printf("姓名: %s\n", e.Name) // 直接访问Person的字段
	fmt.Printf("年龄: %d\n", e.Age)  // 直接访问Person的字段
	fmt.Printf("部门: %s\n", e.Department)
	fmt.Printf("职位: %s\n", e.Position)
	fmt.Printf("个人介绍: %s\n", e.Introduce()) // 直接调用Person的方法
	fmt.Println("=================")
}

// Employee的专用方法
func (e Employee) Work() {
	fmt.Printf("%s正在%s部门工作...\n", e.Name, e.Department)
}

// 3. 创建员工信息的工具函数
func CreateEmployee(name string, age int, id, dept, position string) Employee {
	return Employee{
		Person: Person{
			Name: name,
			Age:  age,
		},
		EmployeeID: id,
		Department: dept,
		Position:   position,
	}
}

func Interfance2() {
	// 方法1：直接创建Employee
	emp1 := Employee{
		Person: Person{
			Name: "张三",
			Age:  28,
		},
		EmployeeID: "E1001",
		Department: "技术部",
		Position:   "高级工程师",
	}

	// 方法2：使用构造函数
	emp2 := CreateEmployee("李四", 32, "E1002", "市场部", "市场经理")

	// 方法3：分步创建
	var emp3 Employee
	emp3.Name = "王五" // 直接访问嵌入结构的字段
	emp3.Age = 25
	emp3.EmployeeID = "E1003"
	emp3.Department = "人力资源部"
	emp3.Position = "招聘专员"

	fmt.Println("员工信息展示:")
	fmt.Println()

	// 调用PrintInfo方法
	emp1.PrintInfo()
	emp2.PrintInfo()
	emp3.PrintInfo()

	fmt.Println("\n工作状态:")
	emp1.Work()
	emp2.Work()
	emp3.Work()

	// 演示多态性
	fmt.Println("\n=== 多态性演示 ===")
	employees := []Employee{emp1, emp2, emp3}

	for i, emp := range employees {
		fmt.Printf("\n员工%d:\n", i+1)
		emp.PrintInfo()
	}

	// 访问嵌入结构体的方法
	fmt.Println("\n=== 个人介绍 ===")
	fmt.Println(emp1.Introduce())
	fmt.Println(emp2.Introduce())
	fmt.Println(emp3.Introduce())
}
