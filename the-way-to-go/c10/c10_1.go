package c10

import (
	"fmt"
	"strings"
)

type Person struct {
	firstName string
	lastName  string
}

func upPerson(p *Person) {
	p.firstName = strings.ToUpper(p.firstName)
	p.lastName = strings.ToUpper(p.lastName)
}

func personTest() {
	// 值类型
	var p1 Person
	p1.firstName = "Jacob"
	p1.lastName = "Bishop"
	upPerson(&p1)
	fmt.Printf("The name of the person is %s %s\n", p1.firstName, p1.lastName)

	// 指针
	p2 := new(Person)
	p2.firstName = "Jacob"
	p2.lastName = "Bishop"
	(*p2).lastName = "Bishop" // 合法
	upPerson(p2)
	fmt.Printf("The name of the person is %s %s\n", p2.firstName, p2.lastName)

	// 字面值
	p3 := &Person{"Jacob", "Bishop"}
	upPerson(p3)
	fmt.Printf("The name of the person is %s %s\n", p3.firstName, p3.lastName)
}

func C10_1() {

	personTest()

}
