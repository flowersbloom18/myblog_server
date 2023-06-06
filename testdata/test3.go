package main

import (
	"fmt"
	"gorm.io/gorm"
	"myblog_server/core"
	"myblog_server/global"
	"myblog_server/models"
)

func createBlog(db *gorm.DB, title string, abstract string, content string, category string, tags []string) error {
	// 检查分类是否存在，如果不存在则创建
	var existingCategory models.Category
	err := db.FirstOrCreate(&existingCategory, models.Category{Name: category}).Error
	if err != nil {
		return fmt.Errorf("failed to create category: %s", err.Error())
	}

	// 检查标签是否存在，如果不存在则创建
	var existingTags []models.Tag
	for _, tagName := range tags {
		var tag models.Tag
		err := db.FirstOrCreate(&tag, models.Tag{Name: tagName}).Error
		if err != nil {
			return fmt.Errorf("failed to create tag: %s", err.Error())
		}
		existingTags = append(existingTags, tag)
	}

	// 创建博客，并关联分类和标签
	blog := models.Blog{
		Title:    title,
		Abstract: abstract,
		Content:  content,
		Tags:     existingTags,
	}

	err = db.Create(&blog).Error
	if err != nil {
		return fmt.Errorf("failed to create blog: %s", err.Error())
	}

	return nil
}

func main() {

	core.InitConf()
	global.Log = core.InitLogger()
	global.DB = core.InitGorm()

	db := global.DB

	// 示例：创建一个博客
	title := "Sample Blog"
	abstract := "This is a sample blog"
	content := "Lorem ipsum dolor sit amet, consectetur adipiscing elit."
	category := "Frontend"
	tags := []string{"Vue", "Vite"}

	err := createBlog(db, title, abstract, content, category, tags)
	if err != nil {
		panic(err)
	}

	fmt.Println("Blog created successfully!")
}
