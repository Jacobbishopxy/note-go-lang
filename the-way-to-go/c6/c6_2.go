package c6

import (
	"errors"
	"fmt"
	"math"
)

func multReturnval1(x, y int) (int, int, int) {
	return x + y, x * y, x - y
}

func multReturnval2(x, y int) (sum, prod, sub int) {
	sum = x + y
	prod = x * y
	sub = x - y
	return
}

func mySqrt1(x float64) (float64, error) {
	res, err := math.Sqrt(x), errors.New("negative input is not allowed")
	if x < 0 {
		return 0, err
	}
	return math.Sqrt(res), nil
}

func mySqrt2(x float64) (res float64, err error) {
	res, err = math.Sqrt(x), errors.New("negative input is not allowed")
	if x < 0 {
		return 0, err
	}
	return math.Sqrt(res), nil
}

// 接收指针，修改 reply 变量
func multiply(a, b int, reply *int) {
	*reply = a * b
}

func C6_2() {

	x, y := 2, 3

	a1, a2, a3 := multReturnval1(x, y)
	fmt.Println(a1, a2, a3)

	b1, b2, b3 := multReturnval2(x, y)
	fmt.Println(b1, b2, b3)

	if res, err := mySqrt1(-10); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(res)
	}

	if res, err := mySqrt2(10); err != nil {
		fmt.Println(err.Error())
	} else {
		fmt.Println(res)
	}

	n := 0
	reply := &n
	multiply(10, 5, reply)
	fmt.Println(*reply)
}
