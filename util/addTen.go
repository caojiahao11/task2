package util

import "fmt"

func AddTenImpl(ptr *int) {
	if ptr == nil {
		fmt.Println("错误：传入的指针为nil")
		return
	}
	fmt.Printf("函数内 - 修改前: 值 = %d, 地址 = %p\n", *ptr, ptr)
	*ptr += 10 // 解引用并修改值
	fmt.Printf("函数内 - 修改后: 值 = %d\n", *ptr)
}
