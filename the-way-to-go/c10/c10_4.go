package c10

import (
	"fmt"
	"reflect"
)

// 临时注释掉（因为 tag 的内容不符合语法）以便屏蔽 warning
type TagType struct {
	field1 bool   //"An important answer"
	field2 string //"The name of the thing"
	field3 int    //"How much there are"
}

func refTag(tt TagType, ix int) {
	ttType := reflect.TypeOf(tt)
	ixField := ttType.Field(ix)
	fmt.Printf("%v\n", ixField.Tag)
}

func C10_4() {

	tt := TagType{true, "Jacob Bishop", 1}

	for i := 0; i < 3; i++ {
		refTag(tt, i)
	}

}
