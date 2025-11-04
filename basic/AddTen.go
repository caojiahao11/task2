package basic

import (
	"fmt"
	"github.com/caojiahao11/task2/util"
)

func AddTen() {
	//指针问题 +10
	fmt.Println("---指针问题+10---")
	num := 5
	fmt.Printf("调用前 - num值: %d, 地址: %p\n", num, &num)

	util.AddTenImpl((&num)) // 传递num的地址

	fmt.Printf("调用后 - num值: %d\n", num)

}
