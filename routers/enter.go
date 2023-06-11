package routers

import (
	"github.com/gin-gonic/gin"
	"myblog_server/global"
	"net/http"
)

type RouterGroup struct {
	*gin.RouterGroup
}

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	router := gin.Default()

	// 在gin框架中注册一个静态文件服务，将指定目录下的文件暴露给客户端访问。
	router.StaticFS("uploads", http.Dir("uploads"))
	// 全局路由前缀api
	apiRouterGroup := router.Group("api")

	// 虚拟的路由群包裹真正的路由群构建了对象
	routerGroupApp := RouterGroup{apiRouterGroup}
	// 挂载
	routerGroupApp.User()
	routerGroupApp.Log()
	routerGroupApp.Category()
	routerGroupApp.Tag()
	routerGroupApp.Blog()
	routerGroupApp.Info()
	routerGroupApp.About()
	routerGroupApp.FriendLink()
	routerGroupApp.Music()
	routerGroupApp.Collect()
	routerGroupApp.Comment()
	routerGroupApp.Announcement()
	routerGroupApp.Attachment()

	return router
}
