package basic

import (
	"fmt"
	"github.com/caojiahao11/task2/util"
	_ "github.com/caojiahao11/task2/util"
)

func DoubleSliceWithRange() {
	numbers := []int{1, 2, 3, 4, 5}
	fmt.Printf("修改前: %v\n", numbers)

	util.DoubleSliceWithRangeImpl(&numbers) // 传递切片的指针

	fmt.Printf("修改后: %v\n", numbers)
}
