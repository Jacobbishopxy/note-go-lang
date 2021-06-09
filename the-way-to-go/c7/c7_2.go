package c7

import (
	"bytes"
	"fmt"
)

func f1() {
	var arr1 [6]int
	var slice1 []int = arr1[2:5] // 索引 5 没有被包含！

	// 加载整数进数组：0,1,2,3,4,5
	for i := 0; i < len(arr1); i++ {
		arr1[i] = i
	}

	// 打印切片
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}

	fmt.Printf("The length of arr1 is %d\n", len(arr1))
	fmt.Printf("The length of slice1 is %d\n", len(slice1))
	fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))

	// 扩充切片
	slice1 = slice1[0:4]
	for i := 0; i < len(slice1); i++ {
		fmt.Printf("Slice at %d is %d\n", i, slice1[i])
	}
	fmt.Printf("The length of slice1 is %d\n", len(slice1))
	fmt.Printf("The capacity of slice1 is %d\n", cap(slice1))
}

// fake func, only for demonstration
func getNextString() (string, bool) {
	return "", true
}

// 这种实现方式比使用 `+=` 要更节省内存和 CPU
func f2() {
	var buffer bytes.Buffer
	for {
		if s, ok := getNextString(); ok {
			buffer.WriteString(s)
		} else {
			break
		}
	}
	fmt.Print(buffer.String(), "\n")
}

func sum(arrF []float32) (res float32) {
	res = 0
	for _, v := range arrF {
		res += v
	}
	return
}

func C7_2() {
	f1()

	f2()

	var arrF = []float32{1, 2, 3, 4}
	fmt.Printf("sumations is: %v\n", sum(arrF))
	fmt.Printf("sumations is: %v\n", sum(arrF[:]))
}
