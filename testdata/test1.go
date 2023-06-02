package main

import (
	"fmt"
	"strings"
)

// 使用go判断字符串在去除空字符串之后是否在结尾包含.com

func main() {
	str := "  abc.com  "
	str = strings.TrimSpace(str)
	if strings.HasSuffix(str, ".com") {
		fmt.Println("包含.com\t登录方式为：邮箱")
	} else {
		fmt.Println("不包含.com\t登录方式为：用户名")
	}

	// qq直接登录的话算什么？就定位qq登录。
}
