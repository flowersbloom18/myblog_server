package info

import (
	"fmt"
	"time"
)

// GetMonthDay 获取月日，用于历史上的今天查询
func GetMonthDay() string {
	currentTime := time.Now()
	month := currentTime.Month()
	day := currentTime.Day()

	// 总是以两位数显示，一位则补0
	monthStr := fmt.Sprintf("%02d", month)
	dayStr := fmt.Sprintf("%02d", day)
	return fmt.Sprintf("%s%s", monthStr, dayStr)
}
