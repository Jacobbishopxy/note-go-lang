package c11

import "fmt"

// 空接口或者最小接口
// 不包含任何方法，对实现不做任何要求

var str = "ABC"

type Person struct {
	name string
	age  int
}

type Any interface{}

func emptyInterface1() {
	var val Any

	val = 5
	fmt.Printf("val has the value: %v\n", val)

	val = str
	fmt.Printf("val has the value: %v\n", val)

	pers1 := new(Person)
	pers1.name = "Rob Pike"
	pers1.age = 55
	val = pers1
	fmt.Printf("val has the value: %v\n", val)

	switch t := val.(type) {
	case int:
		fmt.Printf("Type int %T\n", t)
	case string:
		fmt.Printf("Type string %T\n", t)
	case bool:
		fmt.Printf("Type boolean %T\n", t)
	case *Person:
		fmt.Printf("Type pointer to Person %T\n", t)
	default:
		fmt.Printf("Unexpected type %T", t)
	}
}

// 空接口在 `type-switch` 中联合 `lambda` 函数：

type specialString string

var whatIsThis specialString = "hello"

func typeSwitch() {
	testFunc := func(any interface{}) {
		switch v := any.(type) {
		case bool:
			fmt.Printf("any %v is a bool type", v)
		case int:
			fmt.Printf("any %v is an int type", v)
		case float32:
			fmt.Printf("any %v is a float32 type", v)
		case string:
			fmt.Printf("any %v is a string type", v)
		case specialString:
			fmt.Printf("any %v is a special String!", v)
		default:
			fmt.Println("unknown type!")

		}
	}

	testFunc(whatIsThis)
}

// 通用类型的节点数据结构

type Node struct {
	le   *Node
	data interface{}
	ri   *Node
}

func NewNode(left, right *Node) *Node {
	return &Node{left, nil, right}
}

func (n *Node) SetData(data interface{}) {
	n.data = data
}

func nodeStructures() {
	root := NewNode(nil, nil)
	root.SetData("root node")
	a := NewNode(nil, nil)
	a.SetData("left node")
	b := NewNode(nil, nil)
	b.SetData("right node")
	root.le = a
	root.ri = b

	fmt.Printf("%v\n", root)
}

// 接口到接口
// 一个接口的值可以赋值给另一个接口变量，只要底层类型实现了必要的方法。
// 这个转换是在运行时进行检查的，转换失败会导致一个运行时错误

func C11_9() {

	emptyInterface1()

	typeSwitch()

	nodeStructures()
}
