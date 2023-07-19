package models

import (
	"gorm.io/gorm"
	"myblog_server/global"
	"myblog_server/models/model_type"
	"os"
)

type Attachment struct {
	MODEL
	Name     string                  `json:"name"`                        // 附件名称
	Url      string                  `json:"url"`                         // 附件存储url路径
	Type     string                  `json:"type"`                        // 附件类型（MP3、MP4等）
	Size     string                  `json:"size"`                        // 附件大小(MB；字符串)
	Hash     string                  `json:"hash"`                        // 附件的hash值，用于判断重复附件
	Location model_type.LocationType `gorm:"default:1" json:"image_type"` // 附件的存储位置，本地还是七牛
}

// BeforeDelete 【Hook 是在创建、查询、更新、删除等操作之前、之后调用的函数】
func (attachment *Attachment) BeforeDelete(tx *gorm.DB) (err error) {
	// 如果数据存储在本地，则删除本地存储，否则只删除数据库数据。[即远程只删除记录，不删除数据。]
	if attachment.Location == model_type.Local {
		// 本地图片，删除，还要删除本地的存储
		err = os.Remove(attachment.Url)
		if err != nil {
			global.Log.Error(err)
			return err
		}
	}
	return nil
}
