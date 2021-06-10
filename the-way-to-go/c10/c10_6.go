package c10

import (
	"fmt"
	"math"
)

// 结构体类型加载方法

type TwoInts struct {
	a int
	b int
}

func (tn *TwoInts) AddThem() int {
	return tn.a + tn.b
}

func (tn *TwoInts) AddToParam(param int) int {
	return tn.a + tn.b + param
}

func structMethod() {
	two1 := new(TwoInts)
	two1.a = 12
	two1.b = 10

	fmt.Printf("The sum is: %d\n", two1.AddThem())
	fmt.Printf("Add them to the param: %d\n", two1.AddToParam(20))

	two2 := TwoInts{3, 4}
	fmt.Printf("The sum is: %d\n", two2.AddThem())
}

// 非结构体类型加载方法

type IntVector []int

func (v IntVector) Sum() (s int) {
	for _, x := range v {
		s += x
	}
	return
}

func nonStructMethod() {
	fmt.Println(IntVector{1, 2, 3}.Sum())
}

// 指针或值作为接收者
// 鉴于性能原因，`recv` 最常见的是一个指向 receiver_type 的指针
// 因为我们不想要一个实例的拷贝，特别是 receiver 类型是结构体时

type Bp struct {
	thing int
}

func (b *Bp) change() { b.thing = 1 }

func (b Bp) write() string { return fmt.Sprint(b) }

func pointerValue() {
	var b1 Bp // b1 是值
	b1.change()
	fmt.Println(b1.write())

	b2 := new(Bp) // b2是指针
	b2.change()
	fmt.Println(b2.write())
}

// 在值和指针上调用方法：

type List []int

func (l List) Len() int        { return len(l) }
func (l *List) Append(val int) { *l = append(*l, val) }

func methodSet1() {
	// 值
	var lst List
	lst.Append(1)
	fmt.Printf("%v (len: %d)", lst, lst.Len()) // [1] (len: 1)

	// 指针
	plst := new(List)
	plst.Append(2)
	fmt.Printf("%v (len: %d)", plst, plst.Len()) // &[2] (len: 1)
}

// 方法和未导出字段
// （假设以下结构体位于别的包中）编写 getter setter 方法

type Person2 struct {
	firstName string
	// lastName  string
}

func (p *Person2) FirstName() string {
	return p.firstName
}

func (p *Person2) SetFirstName(newName string) {
	p.firstName = newName
}

func usePerson2() {
	p := new(Person2)

	p.SetFirstName("Jacob")
	fmt.Println(p.FirstName())
}

// 内嵌类型的方法和继承

type Engine interface {
	Start()
	Stop()
}

type Car struct {
	Engine
}

func (c *Car) GoToWorkIn() {
	// get in car
	c.Start()
	// drive to work
	c.Stop()
	// get out of car
}

// 完整例子：展示内嵌结构体上的方法可以直接在外层类型的实例上调用

// 原始结构体
type Point struct {
	x, y float64
}

// 其方法
func (p *Point) Abs() float64 {
	return math.Sqrt(p.x*p.x + p.y*p.y)
}

// 外壳
type NamedPoint struct {
	Point
	name string
}

func method3() {
	n := &NamedPoint{Point{3, 4}, "Pythagoras"}

	// 直接在外层类型实例上调用内嵌结构体的 Abs 方法
	fmt.Println(n.Abs())
}

// 如何在类型中嵌入功能
// a. 聚合（或组合）：包含一个所需功能类型的具名字段
// b. 内嵌：如上面的 Point 和 NamedPoint
//
// 内嵌类型不需要指针，`Customer` 也不需要 `Add` 方法，
// `Customer` 有自己的 `String` 方法，并且在其中调用了 `Log` 的 `String` 方法

// 方法 a 聚合

type Log1 struct {
	msg string
}

type Customer1 struct {
	Name string
	log  *Log1
}

func (l *Log1) Add(s string) {
	l.msg += "\n" + s
}

func (l *Log1) String() string {
	return l.msg
}

func (c *Customer1) Log() *Log1 {
	return c.log
}

func methodA() {
	c := new(Customer1)
	c.Name = "Jacob Bishop"
	c.log = new(Log1)
	c.log.msg = "1 - Study Golang"

	// shorter
	c = &Customer1{"Jacob Bishop", &Log1{"1 - Study Golang"}}
	c.Log().Add("2 - Keep tracking Rust")
	fmt.Println(c.Log())
}

// 方法 b 内嵌

type Log2 struct {
	msg string
}

type Customer2 struct {
	Name string
	Log2
}

func (l *Log2) Add(s string) {
	l.msg += "\n" + s
}

func (l *Log2) String() string {
	return l.msg
}

func (c *Customer2) String() string {
	return c.Name + "\nLog:" + fmt.Sprintln(c.Log2)
}

func methodB() {
	c := &Customer2{"Jacob Bishop", Log2{"1 - Study Golang"}}
	c.Add("2 - Keep tracking Rust")
	fmt.Println(c)
}

// 多重继承
// 多重继承指的是类型获得多个父类型行为的能力。

type Camera struct{}

func (c *Camera) TakeAPicture() string {
	return "Click"
}

type Phone struct{}

func (p *Phone) Call() string {
	return "Ring Ring"
}

type CameraPhone struct {
	Camera
	Phone
}

func multiInheritance() {
	cp := new(CameraPhone)
	fmt.Println("Our new CameraPhone exhibits multiple behaviors...")
	fmt.Println("It exhibits behavior of a Camera: ", cp.TakeAPicture())
	fmt.Println("It works like a Phone too: ", cp.Call())
}

func C10_6() {

	structMethod()

	nonStructMethod()

	pointerValue()

	methodSet1()

	usePerson2()

	method3()

	methodA()

	methodB()

	multiInheritance()
}
