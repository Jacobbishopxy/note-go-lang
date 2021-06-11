package c11

import "fmt"

type List []int

func (l List) Len() int {
	return len(l)
}

func (l *List) Append(val int) {
	*l = append(*l, val)
}

type Appender interface {
	Append(int)
}

func CountInfo(a Appender, start, end int) {
	for i := start; i <= end; i++ {
		a.Append(i)
	}
}

type Lener interface {
	Len() int
}

func LongEnough(l Lener) bool {
	return l.Len()*10 > 42
}

func C11_6() {
	// 一个裸值
	var lst List
	// 编译错误，因为 List 类型没有实现 Appender 接口
	// CountInfo(lst, 1, 10)
	if LongEnough(lst) { // 有效：相同的接收类型
		fmt.Println("- lst is long enough")
	}

	// 一个指针值
	plst := new(List)
	// 通过编译，因为 *List 类型实现了 Append而 接口
	// 有效：相同的接收类型
	CountInfo(plst, 1, 10)
	if LongEnough(plst) {
		// 有效： *List 可被自动解引用
		fmt.Println("- plst is long enough")
	}

	// 总结
	// 在接口上调用方法时，必须有和方法定义时相同的接收者类型或者是可以从具体类型 `P` 直接可以辨识的：
	// - 指针方法可以通过指针调用
	// - 值方法可以通过值调用
	// - 接收者是值的方法可以通过指针调用，因为指针会先被解引用
	// - 接收者是指针的方法不可以通过值调用，因为储存在接口中的值没有地址

	// Go 语言规范定义了接口方法集的调用规则：
	// 类型 *T 的可调用方法集包含接受者为 *T 或 T 的所有方法集
	// 类型 T 的可调用方法集包含接受者为 T 的所有方法集
	// 类型 T 的可调用方法集不包含接受者为 *T 的所有方法集
}
