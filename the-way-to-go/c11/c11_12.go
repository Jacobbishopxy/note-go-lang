package c11

import (
	"errors"
	"fmt"
	"io"
	"log"
)

// Go 的动态类型
//
// Go 没有类：数据（结构体或更一般的类型）和方法是一种松耦合的正交关系。
//
// Go 中的接口跟 Java/C# 类似：都是必须提供一个指定方法集的实现。
// 但是更加灵活通用：任何提供了接口方法实现代码的类型都隐式地实现了该接口，而不用显式地声明。
//
// 和其它语言相比，Go 是唯一结合了接口值，静态类型检查（是否该类型实现了某个接口），
// 运行时动态转换的语言，并且不需要显式地声明类型是否满足某个接口。
// 该特性允许我们在不改变已有的代码的情况下定义和使用新接口。

type IDuck interface {
	Quack()
	Walk()
}

func DuckDance(duck IDuck) {
	for i := 1; i <= 3; i++ {
		duck.Quack()
		duck.Walk()
	}
}

type Bird struct {
	// ...
}

func (b *Bird) Quack() {
	fmt.Println("I am quacking!")
}

func (b *Bird) Walk() {
	fmt.Println("I am walking!")
}

func duckDance() {
	b := new(Bird)
	DuckDance(b)
}

// 动态方法调用
//
// Go 的动态类型不是延迟绑定的（即不是在运行时进行），通常需要编译器静态检查的支持：
// 当变量被赋值给一个接口类型的变量时，编译器会检查其是否实现了该接口的所有函数。

// 定义了两个接口 `GobEncoder` `GobDecoder`
// 允许类型自己实现从流编解码的具体方式：如果没有实现就使用标准的反射方式。

type xmlWriter interface {
	WriteXML(w io.Writer) error
}

// Exported XML streaming function.
func StreamXML(v interface{}, w io.Writer) error {
	if xw, ok := v.(xmlWriter); ok {
		// It’s an  xmlWriter, use method of asserted type.
		return xw.WriteXML(w)
	}
	// No implementation, so we have to use our own function (with perhaps reflection):
	return encodeToXML(v, w)
}

// Internal XML encoding function.
func encodeToXML(v interface{}, w io.Writer) error {
	// ...
	return errors.New("fake error")
}

// 接口的提取
//
// "提取接口" 是非常有用的设计模式，可以减少需要的类型和方法数量，
// 而且不需要像传统的基于类的面向对象语言那样维护整个的类层次结构。
//
// Go 接口可以让开发者找出自己写的程序中的类型。
// 假设有一些拥有共同行为的对象，并且开发者想要抽象出这些行为，这时就可以创建一个接口来使用。

// 旧接口
type MyShaper interface {
	Area() float32
}

// （添加）新接口
type MyTopologicalGenus interface {
	Rank() int
}

type MySquare struct {
	side float32
}

// 旧接口的实现
func (sq *MySquare) Area() float32 {
	return sq.side * sq.side
}

// （添加）新接口的实现
func (sq *MySquare) Rank() int {
	return 1
}

type MyRectangle struct {
	length, width float32
}

// 旧接口的实现
func (r MyRectangle) Area() float32 {
	return r.length * r.width
}

// （添加）新接口的实现
func (r MyRectangle) Rank() int {
	return 2
}

// 不用提前设计出所有的接口；整个设计可以持续演进，而不用废弃之前的决定。
// 类型要实现某个接口，它本身不用改变，你只需要在这个类型上实现新的方法。
func multiInterfacesPoly() {

	r := MyRectangle{5, 3} // Area() of Rectangle needs a value
	q := &MySquare{5}      // Area() of Square needs a pointer

	// 包含 rectangle 和 square 的数组（旧接口）
	shapes := []MyShaper{r, q}
	fmt.Println("Looping through shapes for area ...")
	// 打印数组，并且打印数组元素的 Area 方法的返回
	for n := range shapes {
		fmt.Println("Shape details: ", shapes[n])
		fmt.Println("Area of this shape is: ", shapes[n].Area())
	}

	// 包含 rectangle 和 square 的数组（新接口）
	topgen := []MyTopologicalGenus{r, q}
	fmt.Println("Looping through topgen for rank ...")

	for n := range topgen {
		fmt.Println("Shape details: ", topgen[n])
		fmt.Println("Topological Genus of this shape is: ", topgen[n].Rank())
	}
}

// 接口的继承
//
// 当一个类型包含（内嵌）另一个类型（实现了一个或多个接口）的指针时，
// 这个类型就可以使用（另一个类型）所有的接口方法。

type Task struct {
	Command string
	*log.Logger
}

// Task 类型的工厂方法像这样
// 当 `log.Logger` 实现了 `Log()` 方法后，Task 实例就可以这样调用该方法：`task.Log()`
func NewTask(command string, logger *log.Logger) *Task {
	return &Task{command, logger}
}

func C11_12() {

	duckDance()

	multiInterfacesPoly()
}
