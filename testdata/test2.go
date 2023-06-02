package main

import "fmt"

func main() {

	idList := [2]uint{1, 3}
	logContent := fmt.Sprintf("用户删除，删除ID列表%v", idList)
	fmt.Println(logContent)
}
