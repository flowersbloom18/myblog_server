package info_service

import (
	"myblog_server/global"
	"myblog_server/models"
	info2 "myblog_server/utils/info"
)

func (InfoService) GetInfoService(url string, typeID int) (string, error) {
	db := global.DB

	// 查询数据，如果有，则返回。否则或许数据创建，并返回。
	var info1 models.Info
	count := db.Take(&info1, "type_id=?", typeID).RowsAffected

	if count != 0 {
		return info1.Content, nil // 返回后解析
	} else if count == 0 {
		// 获取响应结果
		responseData, err := info2.GetHttpResponse(url)
		if err != nil {
			global.Log.Error("Info创建出错，err=", err)
			return "", err
		}
		// 创建数据
		err = db.Create(&models.Info{
			Content: responseData,
			TypeId:  typeID,
		}).Error
		if err != nil {
			global.Log.Error("Info创建出错，err=", err)
			return "", err
		}
	}

	// 查找并返回
	var info3 models.Info
	err := db.Take(&info3, "type_id=?", typeID).Error
	if err != nil {
		global.Log.Error("数据不存在，err=", err)
		return "", err
	}

	return info3.Content, nil
}
