package c5

import (
	"fmt"
	"strings"
)

// 打印一个宽 20，高 10 的矩形 （中空）
func rectangle_stars() {
	mid := make([]string, 20)
	for i := range mid {
		if i == 0 || i == 19 {
			mid[i] = "*"
		} else {
			mid[i] = " "
		}
	}

	for i := 0; i < 10; i++ {
		if i == 0 || i == 9 {
			fmt.Println(strings.Repeat("*", 20))
		} else {
			fmt.Println(strings.Join(mid, ""))
		}
	}
}

func C5_4() {
	rectangle_stars()
}
