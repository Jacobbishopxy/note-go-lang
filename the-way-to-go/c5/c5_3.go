package c5

import (
	"errors"
	"fmt"
)

func case1() {
	var num1 int = 7

	switch {
	case num1 < 0:
		fmt.Println("Number is negative")
	case num1 > 0 && num1 < 10:
		fmt.Println("Number is between 0 and 10")
	default:
		fmt.Println("Number is 10 or greater")
	}
}

func case2() {
	k := 6
	switch k {
	case 4:
		fmt.Println("was <= 4")
		fallthrough
	case 5:
		fmt.Println("was <= 5")
		fallthrough
	case 6:
		fmt.Println("was <= 6")
		fallthrough
	case 7:
		fmt.Println("was <= 7")
		fallthrough
	case 8:
		fmt.Println("was <= 8")
		fallthrough
	default:
		fmt.Println("default case")
	}
}

// 输入整数返回月份
func season(s int) (string, error) {
	res, err := "", errors.New("input for season: 1 - 12")
	switch s {
	case 1:
		res = "Jan"
	case 2:
		res = "Feb"
	case 3:
		res = "Mar"
	case 4:
		res = "Apr"
	case 5:
		res = "May"
	case 6:
		res = "Jun"
	case 7:
		res = "Jul"
	case 8:
		res = "Aug"
	case 9:
		res = "Sep"
	case 10:
		res = "Oct"
	case 11:
		res = "Nov"
	case 12:
		res = "Dec"
	default:
		return res, err
	}
	return res, nil
}

func C5_3() {
	case1()

	case2()

	var s int = 12

	if res, err := season(s); err != nil {
		println("Error occurred", err.Error())
	} else {
		println("Result: ", res)
	}

}
