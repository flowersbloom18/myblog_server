package main

import (
	"fmt"
	"strings"
)

func main() {
	// 获取附件名称【去除后缀名后的】
	fileName := "a/b.cd.mp.mp5"
	nameList := strings.Split(fileName, ".")  // 按照英文点分开
	suffix := "." + nameList[len(nameList)-1] // 后缀名
	fileName = strings.Split(fileName, suffix)[0]

	fmt.Println("附件名称为:", fileName)
	fmt.Println("后缀名称为:", suffix)

}
