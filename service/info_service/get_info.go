package info_service

import (
	"fmt"
	"myblog_server/global"
	"myblog_server/models"
)

func (InfoService) GetInfoService(typeID int) (string, error) {
	db := global.DB

	// 查询数据，如果有，则返回。否则或许数据创建，并返回。
	var info1 models.Info
	count := db.Take(&info1, "type_id=?", typeID).RowsAffected

	if count != 0 {
		return info1.Content, nil // 返回后解析
	} else {
		return "", fmt.Errorf("数据丢失ing")
	}
}
