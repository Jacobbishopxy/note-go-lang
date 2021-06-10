package c11

import (
	"fmt"
	"math"
)

type Square struct {
	side float32
}

type Circle struct {
	radius float32
}

type Shaper interface {
	Area() float32
}

func (sq *Square) Area() float32 {
	return sq.side * sq.side
}

func (ci *Circle) Area() float32 {
	return ci.radius * ci.radius * math.Pi
}

// 类型断言
func C11_3() {
	// 声明变量 areaIntf 实现 Shaper 接口
	var areaIntf Shaper

	// 变量 sql 为 Square 实例的指针
	sql := new(Square)
	// 赋值
	sql.side = 5
	// 再赋值
	areaIntf = sql

	// 确认 areaIntf 为 Square 类型
	if t, ok := areaIntf.(*Square); ok {
		fmt.Printf("The type of areaIntf is: %T\n", t)
	}

	// areaIntf 并不为 Circle 类型
	if u, ok := areaIntf.(*Circle); ok {
		fmt.Printf("The type of areaIntf is: %T\n", u)
	} else {
		fmt.Println("areaIntf does not contain a variable of type Circle")
	}

}
