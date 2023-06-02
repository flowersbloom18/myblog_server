package common

import (
	"fmt"
	"gorm.io/gorm"
	"myblog_server/global"
	"myblog_server/models"
)

type Option struct {
	models.PageInfo
	Debug   bool
	Likes   []string // 模糊匹配的字段
	Preload []string // 预加载的列表
}

func ComList[T any](model T, option Option) (list []T, count int64, err error) {

	DB := global.DB
	// 调试
	if option.Debug {
		DB = global.DB.Session(&gorm.Session{Logger: global.MysqlLog})
	}
	// 排序
	if option.Sort == "" {
		option.Sort = "created_at desc" // 默认按照时间往前排
	}
	DB = DB.Where(model)
	for index, column := range option.Likes {
		if index == 0 {
			DB.Where(fmt.Sprintf("%s like ?", column), fmt.Sprintf("%%%s%%", option.Key))
			continue
		}
		DB.Or(fmt.Sprintf("%s like ?", column), fmt.Sprintf("%%%s%%", option.Key))
	}

	count = DB.Where(model).Find(&list).RowsAffected
	// 这里的query会受上面查询的影响，需要手动复位
	query := DB.Where(model)
	for _, preload := range option.Preload {
		query = query.Preload(preload)
	}
	offset := (option.Page - 1) * option.Limit
	if offset < 0 {
		offset = 0
	}
	err = query.Limit(option.Limit).Offset(offset).Order(option.Sort).Find(&list).Error

	return list, count, err
}
