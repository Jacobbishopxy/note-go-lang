package c6

import (
	"fmt"
)

func f1() {
	fmt.Printf("In function1 at the top\n")
	defer f2()
	fmt.Printf("In function1 at the bottom!\n")
}

func f2() {
	fmt.Printf("function2: Deferred until the end of the calling function!")
}

func f3() {
	for i := 0; i < 5; i++ {
		defer fmt.Printf("%d \n", i)
	}
}

// defer 可以用作于：
// 1. 关闭文件流
// 2. 解锁一个加锁的资源
// 3. 打印最终报告
// 4. 关闭数据库链接
// ...
func C6_4() {

	f1()

	f3()
}
