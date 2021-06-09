package c7

import "fmt"

// 字符串生成字节切片
func stringToByteSlice() {
	s := "\u00ff\u754c"
	for i, c := range s {
		fmt.Printf("%d:%c \n", i, c)
	}
}

// 修改字符串中的某个字符
func changeStringChar(s string, ix int, c rune) string {
	sc := []byte(s)
	sc[ix] = byte(c)
	return string(sc)
}

// 字节数组对比
func compareStringByteSlice(a, b []byte) int {
	for i := 0; i < len(a) && i < len(b); i++ {
		switch {
		case a[i] > b[i]:
			return 1
		case a[i] < b[i]:
			return -1
		}
	}

	switch {
	case len(a) < len(b):
		return -1
	case len(a) > len(b):
		return 1
	}
	return 0
}

func appendFuncCommonUsage() {

	a := []int{1, 2, 3}
	b := []int{4, 5, 6, 7}

	// 1. 将切片 b 的元素追加到切片 a 之后
	fmt.Println(1, append(a, b...))

	// 2. 复制切片 a 的元素到新的切片 b 上
	bb := make([]int, len(a))
	copy(bb, a)
	fmt.Println(2, bb)

	// 3. 删除位于索引 i 的元素
	fmt.Println(3, append(b[:1], b[2:]...))

	// 4. 切除切片中从索引 i 至 j 位置的元素
	fmt.Println(4, append(b[:1], b[3:]...))

	// 5. 为切片扩展 j 个元素长度
	fmt.Println(5, append(a, make([]int, 10)...))

	// 6. 在索引 i 位置插入元素 x
	fmt.Println(6, append(a[:1], append([]int{1}, a[1:]...)...))

	// 7. 在索引 i 位置插入长度为 j 的新切片
	fmt.Println(7, append(a[:1], append(make([]int, 10), a[1:]...)...))

	// 8. 在索引 i 位置插入切片 b 的所有元素
	fmt.Println(8, append(a[:1], append(b, a[1:]...)...))

	// 9. 取出位于切片 a 最末尾的元素 x
	x, a := a[len(a)-1:], a[:len(a)-1]
	fmt.Println(9, x)

	// 10. 将元素 x 追加到切片 a
	fmt.Println(10, append(a, x...))

}

func C7_6() {

	stringToByteSlice()

	fmt.Println(changeStringChar("Halo", 3, 'u'))

	a := []byte("ha")
	b := []byte("hoo")
	fmt.Println(compareStringByteSlice(a, b))

	appendFuncCommonUsage()
}
