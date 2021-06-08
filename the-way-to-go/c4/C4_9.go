package c4

import "fmt"

func C4_9() {
	var i1 = 5
	fmt.Printf("An integer: %d, its location in memory: %p\n", i1, &i1)

	// 指向 i1 的指针，类型为 *int
	var intP *int = &i1
	fmt.Printf("The value at memory location %p is %d\n", intP, *intP)

	// 使用指针修改值
	s := "good bye"
	var p *string = &s
	*p = "ciao"
	fmt.Printf("Here is the pointer p: %p\n", p)  // prints address
	fmt.Printf("Here is the string *p: %s\n", *p) // prints string
	fmt.Printf("Here is the string s: %s\n", s)   // prints same string

	// 不可对空指针解引用赋值
	// var t *int = nil
	// *t = 0
}
