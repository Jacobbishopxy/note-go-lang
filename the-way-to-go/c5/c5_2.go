package c5

import (
	"fmt"
	"strconv"
)

func C5_2() {
	var orig string = "ABC"
	var newS string

	fmt.Printf("The size of ints is: %d\n", strconv.IntSize)
	// 字符串转为整数
	an, err := strconv.Atoi(orig)
	// 惯用法，如果有错误则退出代码
	if err != nil {
		fmt.Printf("orig %s is not an integer - exiting with error\n", orig)
		return
	}
	fmt.Printf("The integer is %d\n", an)
	an = an + 5
	newS = strconv.Itoa(an)
	fmt.Printf("The new string is: %s\n", newS)
}
