package main

import (
	"fmt"
	"myblog_server/core"
	"myblog_server/global"
	"myblog_server/models"
)

func main() {

	core.InitConf()
	global.Log = core.InitLogger()
	global.DB = core.InitGorm()

	db := global.DB
	// 查询分类下的博客数量
	var category models.Category
	categoryID := 1 // 替换为实际的分类ID
	err := db.Preload("Blogs").First(&category, categoryID).Error
	if err != nil {
		panic("Failed to query category")
	}

	blogCount := len(category.Blogs)
	fmt.Printf("Category '%s' has %d blogs\n", category.Name, blogCount)
}
