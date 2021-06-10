package c10

import "fmt"

type innerS struct {
	in1 int
	in2 int
}

// 在一个结构体中对于每一种数据类型只能有一个匿名字段
type outerS struct {
	b      int
	c      float32
	int    // 匿名字段
	innerS // 匿名字段
}

func structsAnonymousFields() {
	outer := new(outerS)
	outer.b = 6
	outer.c = 7.5
	outer.int = 60
	outer.in1 = 5
	outer.in2 = 10

	fmt.Printf("outer.b is: %d\n", outer.b)
	fmt.Printf("outer.c is: %f\n", outer.c)
	fmt.Printf("outer.int is: %d\n", outer.int)
	fmt.Printf("outer.in1 is: %d\n", outer.in1)
	fmt.Printf("outer.in2 is: %d\n", outer.in2)

	// 使用结构体字面量
	outer2 := outerS{6, 7.5, 60, innerS{5, 10}}
	fmt.Println("outer2 is:", outer2)
}

type A struct {
	ax, ay int
}

type B struct {
	A
	bx, by float32
}

func embeddStruct() {
	b := B{A{1, 2}, 3.0, 4.0}
	fmt.Println(b.ax, b.ay, b.bx, b.by)
	fmt.Println(b.A)
}

func C10_5() {

	structsAnonymousFields()

	embeddStruct()

}
