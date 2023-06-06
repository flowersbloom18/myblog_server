package desensitization

import (
	"myblog_server/global"
	"strings"
)

// DesensitizationEmail 邮箱脱敏处理
func DesensitizationEmail(email string) string {
	if email == "" {
		return "" // 邮箱不存在直接给空
	}
	eList := strings.Split(email, "@")
	if len(eList) != 2 {
		return "" // 邮箱不完整，直接给空
	}

	// 如果邮箱长度>=5获取邮箱前三个和后两个，否则给个*@.com
	if len(eList[0]) >= 5 {
		return eList[0][:3] + "***" + eList[0][len(eList[0])-2:] + "@" + eList[1]
	} else {
		global.Log.Info("***@" + eList[1])
		return "***@" + eList[1]
	}
}
func DesensitizationUserName(name string) string {

	// 获取第一位数
	first := ""
	if len(name) > 0 {
		first = string(name[0])
	}

	// 获取最后一位数
	last := ""
	if len(name) >= 2 {
		last = string(name[len(name)-1])
	} else {
		last = ""
	}

	return first + "***" + last
}
