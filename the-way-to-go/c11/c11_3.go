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

// 类型分类函数
// 在处理来自于外部的、类型未知的数据时，比如解析诸如 JSON 或 XML 编码的数据，类型测试和转换会非常有用。
func classifier(items ...interface{}) {
	for i, x := range items {
		switch x.(type) {
		case bool:
			fmt.Printf("Param #%d is a bool\n", i)
		case float64:
			fmt.Printf("Param #%d is a float64\n", i)
		case int, int64:
			fmt.Printf("Param #%d is a int\n", i)
		case nil:
			fmt.Printf("Param #%d is a nil\n", i)
		case string:
			fmt.Printf("Param #%d is a string\n", i)
		default:
			fmt.Printf("Param #%d is unknown\n", i)
		}
	}
}

// 测试一个值是否实现了某个接口
type Stringer interface {
	String() string
}

type String struct {
	v string
}

func (v *String) String() string {
	return ">" + v.v + "<"
}

func C11_3() {

	// 类型断言
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

	// 类型判断：type-switch

	// 变量 t 得到了 `areaIntf` 的值和类型
	switch t := areaIntf.(type) {
	case *Square:
		fmt.Printf("Type Square %T with value %v\n", t, t)
	case *Circle:
		fmt.Printf("Type Circle %T with value %v\n", t, t)
	case nil:
		fmt.Printf("nil value: nothing to check?\n")
	default:
		fmt.Printf("Unexpected type %T\n", t)
	}

	// 类型分类函数使用案例
	classifier(13, -14.3, "BELGIUM", complex(1, 2), nil, false)

	// 测试 v 是否实现了 Stringer 接口
	//（用于运行时的判断，注释是因为静态下v已经实现了 Stringer，ide亮出警告）
	// var v Stringer = &String{"233"}
	// if sv, ok := v.(Stringer); ok {
	// 	fmt.Printf("v implements String(): %s\n", sv.String()) // note: sv, not v
	// }

}
