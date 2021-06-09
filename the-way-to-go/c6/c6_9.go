package c6

import (
	"fmt"
	"strings"
)

func fib(n int) []int {
	res := []int{}

	f := t()

	for i := 0; i < n; i++ {
		v := f()
		fmt.Println(v)
		res = append(res, v)
	}

	return res
}

// 闭包的每次调用都会改变自身的内部状态（n1, n2）
func t() func() int {
	n1, n2 := 0, 1
	return func() int {
		tmp := n1
		n1, n2 = n2, (n1 + n2)
		return tmp
	}
}

// 工厂函数，用于返回另一个函数
func makeAddSuffix(suffix string) func(string) string {
	return func(name string) string {
		if !strings.HasSuffix(name, suffix) {
			return name + suffix
		}
		return name
	}
}

func C6_9() {

	fmt.Println("Fib: ", fib(4))

	addBmp := makeAddSuffix(".bmp")
	addJpeg := makeAddSuffix(".jpeg")

	fmt.Println(addBmp("file"))
	fmt.Println(addJpeg("file"))
}
