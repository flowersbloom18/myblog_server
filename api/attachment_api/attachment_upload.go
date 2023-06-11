package attachment_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/fs"
	"myblog_server/global"
	"myblog_server/models/response"
	"myblog_server/service"
	"myblog_server/service/attachment_service"
	"os"
)

// AttachmentUploadView 上传单个附件，返回附件的url
func (AttachmentApi) AttachmentUploadView(c *gin.Context) {
	// 上传多个附件
	form, err := c.MultipartForm()
	if err != nil {
		response.FailWithMessage(err.Error(), c)
		return
	}
	// 获取上传的附件列表
	//fileList, ok := form.File["images"]
	fileList, ok := form.File["attachments"]
	if !ok {
		response.FailWithMessage("不存在的附件", c)
		return
	}

	// 判断路径是否存在
	basePath := global.Config.Upload.Path
	_, err = os.ReadDir(basePath)
	// 不存在就创建
	if err != nil {
		// 递归创建
		err = os.MkdirAll(basePath, fs.ModePerm)
		if err != nil {
			global.Log.Error(err)
		}
	}
	var lists []attachment_service.FileUploadResponse

	// 判断上传方式，然后循环列表并上传
	if !global.Config.QiNiu.Enable {
		for _, file := range fileList {
			result := service.ServiceApp.AttachmentService.AttachmentLocalUploadService(file, c) // 本地上传附件
			if !result.IsSuccess {
				lists = append(lists, result) //追加失败原因
				continue
			}
			lists = append(lists, result)
		}
	} else {
		for _, file := range fileList {
			result := service.ServiceApp.AttachmentService.AttachmentQiNiuYunUploadService(file) // 七牛云上传附件
			if !result.IsSuccess {
				lists = append(lists, result) //追加失败原因
				continue
			}
			lists = append(lists, result)
		}
	}

	// 对结果进行封装
	if len(lists) == 1 { // 只上传了一张照片的情况
		response.OkWithData(lists, c)
	} else { // 上传多张照片的情况
		success := 0
		fail := 0
		for _, v := range lists {
			if v.IsSuccess {
				success++
			} else {
				fail++
			}
		}
		result := fmt.Sprintf("本次共上传了%d个附件，成功%d个，失败%d个。", len(lists), success, fail)
		response.OkWithMessage(result, c)
	}
	//response.OkWithData(lists, c)
}
