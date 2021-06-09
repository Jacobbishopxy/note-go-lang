package c8

import "fmt"

func C8_2() {

	var value int
	var isPresent bool

	// map 初始化
	map1 := make(map[string]int)
	map1["Tokyo"] = 50
	map1["Shanghai"] = 20
	map1["Seattle"] = 25

	// 取值
	value, isPresent = map1["Shanghai"]
	if isPresent {
		fmt.Printf("The value of \"Beijing\" in map1 is: %d\n", value)
	} else {
		fmt.Printf("map1 does not contain Beijing")
	}

	// if 赋值式取值
	if v, ok := map1["Paris"]; ok {
		fmt.Printf("Value is: %d\n", v)
	} else {
		fmt.Printf("Is \"Paris\" in map1 ?: %t\n", ok)
	}

	// 删键值
	delete(map1, "Tokyo")
	value, isPresent = map1["Washington"]
	if isPresent {
		fmt.Printf("The value of \"Washington\" in map1 is: %d\n", value)
	} else {
		fmt.Println("map1 does not contain Washington")
	}
}
