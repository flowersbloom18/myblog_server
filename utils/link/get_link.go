package link

import (
	"fmt"
	"time"
)

// GetLink 生成博客链接
func GetLink(title string) string {
	// 获取当前时间
	currentTime := time.Now()
	// 转换为年月日的表示方法
	year := currentTime.Year()
	month := currentTime.Month()
	day := currentTime.Day()
	// 格式化月数和日，如果是个位数则在前面加上0
	monthStr := fmt.Sprintf("%02d", month)
	dayStr := fmt.Sprintf("%02d", day)
	date := fmt.Sprintf("%d/%s/%s", year, monthStr, dayStr)
	link := date + "/" + title
	return link
}
