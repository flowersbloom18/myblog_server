package attachment_service

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"mime/multipart"
	"myblog_server/global"
	"myblog_server/models"
	"myblog_server/models/model_type"
	"myblog_server/utils"
	"myblog_server/utils/md5"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// WhiteAttachmentList 附件白名单
var WhiteAttachmentList = []string{
	// 图片类型
	"jpg",
	"png",
	"jpeg",
	"ico",
	"tiff",
	"gif",
	"svg",
	"webp",
	"heic",
	// 音乐类型
	"mp3",
	// 视频类型
	"mp4",
}

type FileUploadResponse struct {
	IsSuccess bool   `json:"is_success"` // 是否成功
	Msg       string `json:"msg"`        // 消息-成功或失败
}

// AttachmentLocalUploadService 本地附件上传服务
// 1、获取基本信息
// 2、判断上传附件是否在白名单内
// 3、判断上传附件大小是否符合需求
// 4、生成Hash值并判断数据库是否存在该文件
// 5、判断上传方式（A.本地）
// 6、改名
// 7、写入数据库
// 【很重要】优化: 第6步，把uploads/file/反方向的钟.mp3，改为uploads/file/hash.mp3，防止文件名称过长导致意外错误而无法访问【错综复杂的名称会出现，资源找不到】
// 还有可优化的空间，但是有些掌握的还不到位，只能暂且使用重命名的方式。
func (AttachmentService) AttachmentLocalUploadService(file *multipart.FileHeader, c *gin.Context) (res FileUploadResponse) {
	// 1、获取【文件名称、生成本地url路径】
	fileName := file.Filename                      // 附件名称
	basePath := global.Config.Upload.Path          // 本地上传路径
	contentType := file.Header.Get("Content-Type") // 附件类型
	url := path.Join(basePath, file.Filename)      // Url路径；

	// 2、附件白名单判断
	nameList := strings.Split(fileName, ".")
	suffix := strings.ToLower(nameList[len(nameList)-1]) // 后缀名
	if !utils.InList(suffix, WhiteAttachmentList) {
		res.Msg = "非法附件"
		return
	}

	// 3、判断附件大小
	size := float64(file.Size) / float64(1024*1024)
	if size >= float64(global.Config.Upload.Size) {
		res.Msg = fmt.Sprintf("附件大小超过设定大小，当前大小为:%.2fMB， 设定大小为：%dMB ", size, global.Config.Upload.Size)
		return
	}
	fileSize := fmt.Sprintf("%.2fMB", size)

	// 4、生成hash值并判断数据库是否存在该附件
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

	// 5、上传
	err = c.SaveUploadedFile(file, url)
	if err != nil {
		global.Log.Error(err)
		res.Msg = err.Error()
		res.IsSuccess = false
		return
	}

	// 6、修改附件名
	ext := filepath.Ext(file.Filename)                      // 扩展名
	newFileName := fmt.Sprintf("%s%s", attachmentHash, ext) // hash值+扩展名
	oldPath := path.Join(basePath, file.Filename)           // 旧的路径
	newPath := path.Join(basePath, newFileName)             // 新的路径
	url = newPath
	err = os.Rename(oldPath, newPath)
	if err != nil {
		global.Log.Error(err)
		res.Msg = err.Error()
		res.IsSuccess = false
		return
	}

	location := model_type.Local // 附件存储位置

	// 7、附件入库
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
