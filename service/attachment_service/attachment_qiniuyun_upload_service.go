package attachment_service

import (
	"fmt"
	"io"
	"mime/multipart"
	"myblog_server/global"
	"myblog_server/models"
	"myblog_server/models/model_type"
	"myblog_server/plugins/qiniu"
	"myblog_server/utils"
	"myblog_server/utils/md5"
	"path/filepath"
	"strings"
)

// AttachmentQiNiuYunUploadService 七牛云附件上传服务
// 1、获取基本信息
// 2、判断上传附件是否在白名单内
// 3、判断上传附件大小是否符合需求
// 4、生成Hash值并判断数据库是否存在该附件
// 5、判断上传方式（B.七牛云）
// 6、写入数据库
// 优化上传的附件名称，依然是哈希值
func (AttachmentService) AttachmentQiNiuYunUploadService(file *multipart.FileHeader) (res FileUploadResponse) {
	// 1、获取【附件名称、类型、url路径】
	fileName := file.Filename                      // 附件名称（带扩展名）
	contentType := file.Header.Get("Content-Type") // 附件类型
	url := ""                                      // 初始化路径

	// 2、附件白名单判断
	nameList := strings.Split(fileName, ".")
	suffix := strings.ToLower(nameList[len(nameList)-1]) // 后缀名
	if !utils.InList(suffix, WhiteAttachmentList) {
		res.Msg = "非法附件"
		return
	}

	// 3、判断附件大小
	size := float64(file.Size) / float64(1024*1024)
	if size >= global.Config.QiNiu.Size {
		res.Msg = fmt.Sprintf("附件大小超过设定大小，当前大小为:%.2fMB， 设定大小为：%.fMB ", size, global.Config.QiNiu.Size)
		return
	}
	fileSize := fmt.Sprintf("%.2fMB", size)

	// 4、生成hash值并判断数据库是否存在该附件【确保七牛云和本地数据不重复】
	fileObj, err := file.Open() // 读取附件内容 hash
	if err != nil {
		global.Log.Error(err)
	}
	byteData, err := io.ReadAll(fileObj)
	attachmentHash := md5.Md5(byteData)

	var attachment models.Attachment // 去数据库中查这个附件是否存在
	err = global.DB.Take(&attachment, "hash = ?", attachmentHash).Error
	if err == nil { // 找到了
		res.Msg = "附件" + attachment.Name + "已存在"
		return
	}

	var location model_type.LocationType
	// 5、七牛云上传
	// 附件名称优化（abc.mp3->hash.mp3)
	ext := filepath.Ext(fileName)                           // 扩展名
	newFileName := fmt.Sprintf("%s%s", attachmentHash, ext) // hash值+扩展名

	url, err = qiniu.UploadAttachment(byteData, newFileName, global.Config.QiNiu.Prefix)
	if err != nil {
		global.Log.Error(err)
		res.Msg = err.Error()
		res.IsSuccess = false
		return
	}
	//res.Msg = "上传七牛成功"
	location = model_type.QiNiu

	// 6、附件入库
	fileName = utils.GetFileName(fileName) // 修正为名称
	global.DB.Create(&models.Attachment{
		Name:     fileName,       // 附件名称
		Url:      url,            // 路径
		Type:     contentType,    // 附件类型
		Size:     fileSize,       // 附件大小
		Hash:     attachmentHash, // 哈希值
		Location: location,       // 存储位置
	})

	res.Msg = "附件" + fileName + "上传成功"
	res.IsSuccess = true
	return
}
