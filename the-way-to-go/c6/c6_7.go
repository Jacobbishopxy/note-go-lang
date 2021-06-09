package c6

import (
	"fmt"
)

func textReplaceNonASCII(s, replacement string) string {
	res := ""

	for i, c := range s {
		if c > 255 {
			res = res + replacement
		} else {
			res = res + string(s[i])
		}
	}

	return res
}

func C6_7() {

	text := "哈hah 你好，how are you？"

	fmt.Printf(">%s<\n", textReplaceNonASCII(text, "?"))
	fmt.Printf(">%s<\n", textReplaceNonASCII(text, " "))

}
