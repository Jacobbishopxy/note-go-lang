package c7

import "fmt"

// 增加切片容量
// 创建一个更大的切片再拷贝原分片内容
func copyAppendSlice() {
	sl_from := []int{1, 2, 3}
	sl_to := make([]int, 10)

	n := copy(sl_to, sl_from)
	fmt.Println(sl_to)
	fmt.Printf("Copied %d elements\n", n) // n == 3

	sl3 := []int{1, 2, 3}
	tmp := []int{4, 5, 6}
	sl3 = append(sl3, tmp...)
	fmt.Println(sl3)
}

func appendByte(slice []byte, data ...byte) []byte {
	m := len(slice)
	n := m + len(data)
	if n > cap(slice) {
		newSlice := make([]byte, (n+1)*2)
		copy(newSlice, slice)
		slice = newSlice
	}
	slice = slice[0:n]
	copy(slice[m:n], data)
	return slice
}

func C7_5() {
	copyAppendSlice()

	s := []byte("Here is a string.")
	fmt.Println(appendByte(s))
}
