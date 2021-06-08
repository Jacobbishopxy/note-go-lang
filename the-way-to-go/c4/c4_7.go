package c4

import (
	"fmt"
	"strconv"
	"strings"
)

func stringHasPrefix() {
	var str string = "This is an example of a string"

	fmt.Printf("T/F? Does the string \"%s\" have prefix %s?", str, "Th")

	// strings.HasPrefix()
	fmt.Printf("%t\n", strings.HasPrefix(str, "Th"))
}

func indexInString(str string) {

	fmt.Printf("The position of \"Marc\" is: ")
	fmt.Printf("%d\n", strings.Index(str, "Marc"))

	// 第一次出现
	fmt.Printf("The position of the first instance of \"Hi\" is: ")
	fmt.Printf("%d\n", strings.Index(str, "Hi"))
	// 最后一次出现
	fmt.Printf("The position of the last instance of \"Hi\" is: ")
	fmt.Printf("%d\n", strings.LastIndex(str, "Hi"))

	fmt.Printf("The position of \"Burger\" is: ")
	fmt.Printf("%d\n", strings.Index(str, "Burger"))
}

func replaceString(str string) {
	// 替换是基于原本str的拷贝而执行的
	fmt.Printf("Changed: %v; original: %v\n", strings.Replace(str, "Hi", "Hello", -1), str)
}

func countString(str string) {
	// 计算字符串出现次数
	fmt.Printf("H count: %d\n", strings.Count(str, "H"))
}

func repeatString() {
	// 重复字符串
	var origS string = "Hi there! "
	var newS string = strings.Repeat(origS, 3)
	fmt.Printf("The new repeated string is: %s\n", newS)
}

func caseString() {
	// 大小写
	var orig string = "Hey, how are you George?"
	var lower string
	var upper string

	fmt.Printf("The original string is: %s\n", orig)
	lower = strings.ToLower(orig)
	fmt.Printf("The lowercase string is: %s\n", lower)
	upper = strings.ToUpper(orig)
	fmt.Printf("The uppercase string is: %s\n", upper)
}

func trimString() {
	// 修剪
	var strToTrim string = " halo, Jacob~  "
	fmt.Printf(">%s<\n", strings.TrimSpace(strToTrim))
	fmt.Printf(">%s<\n", strings.Trim(strToTrim, " "))
	fmt.Printf(">%s<\n", strings.TrimLeft(strToTrim, " "))
	fmt.Printf(">%s<\n", strings.TrimRight(strToTrim, " "))
}

func splitString() {
	// 分割
	var strToSplit string = "Halo Jacob. How are you?"
	fmt.Printf("%+q\n", strings.Fields(strToSplit))
	fmt.Printf("%+q\n", strings.Split(strToSplit, "."))
}

func joinString() {
	// 拼接
	str := strings.Join([]string{"Halo", "Jacob"}, " ")
	fmt.Printf("joinString: %s\n", str)
}

func convertString() {
	// 转换
	var orig string = "666"
	var an int
	var newS string

	fmt.Printf("The size of ints is: %d\n", strconv.IntSize)

	an, _ = strconv.Atoi(orig)
	fmt.Printf("The integer is: %d\n", an)
	an = an + 5
	newS = strconv.Itoa(an)
	fmt.Printf("The new string is: %s\n", newS)
}

func C4_7() {

	stringHasPrefix()

	var str string = "Hi, I'm Marc, Hi."

	indexInString(str)

	replaceString(str)

	countString(str)

	repeatString()

	caseString()

	trimString()

	splitString()

	splitString()

	joinString()

	convertString()
}
