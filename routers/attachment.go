package routers

import api2 "myblog_server/api"

func (router RouterGroup) Attachment() {
	api := api2.ApiGroupApp.AttachmentApi
	// 上传附件
	router.POST("attachment", api.AttachmentUploadView)

	// 删除附件（Hook如果数据在七牛云，则仅仅删除本地数据库信息）
	router.DELETE("attachment", api.AttachmentRemoveView)

	// 修改附件名称
	router.PUT("attachment", api.AttachmentUpdateView)

	// 查找附件（根据名称）
	router.GET("attachment", api.AttachmentListView)
}
