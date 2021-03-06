package routers

import (
	"api/controller/admin"
	_ "api/docs"
	"api/middleware"
	"api/pkg/setting"
	"api/pkg/upload"
	"github.com/gin-gonic/gin"
	"github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	//r.Use(cors.Default())

	// 跨域问题
	r.Use(middleware.Cors())

	r.Use(gin.Logger())

	r.Use(gin.Recovery())

	r.LoadHTMLGlob("templates/*")

	gin.SetMode(setting.RunMode)

	// gin swagger
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// 业务初始化
	r.GET("/admin/exec/ws/:id/ssh/:token", admin.SyncWsApp)

	// 业务上线
	r.GET("/admin/deploy/ws/:id/ssh/:token", admin.PutAppDeployRedo)

	// 业务回滚
	r.GET("/admin/undeploy/ws/:id/ssh/:token", admin.PutAppDeployUndo)

	// ConsoleHost 逻辑
	r.GET("/admin/host/ssh/:id", admin.ConsoleHost)
	r.GET("/admin/ws/:id/ssh/:token", admin.SshConsumer)

	//r.POST("/upload", admin.UploadImage)

	r.StaticFS("/upload/images", http.Dir(upload.GetImageFullPath()))

	//apiv1 := r.Group("/api")
	//{
	//}

	adminv1 := r.Group("/admin", middleware.TokenAuthMiddleware())
	{
		// 用户登录
		adminv1.POST("/user/login", admin.Login)

		adminv1.POST("/user/logout", admin.Logout)
		adminv1.GET("/user/perms/:id", admin.GetUserMenu)

		adminv1.GET("/user", admin.GetUsers)
		adminv1.POST("/user", admin.PostUser)
		adminv1.PUT("/user/:id", admin.PutUser)
		adminv1.DELETE("/user/:id", admin.DeleteUser)
		adminv1.GET("/perms", admin.GetPerms)
		adminv1.POST("/perms", admin.PostPerms)
		adminv1.DELETE("/perms/:id", admin.DeletePerms)
		adminv1.PUT("/perms/:id", admin.PutPerms)
		adminv1.GET("/perms/lists", admin.GetAllPerms)

		adminv1.GET("/roles",admin.GetRole)
		adminv1.POST("/roles", admin.PostRole)
		adminv1.DELETE("/roles/:id", admin.DeleteRole)
		adminv1.PUT("/roles/:id", admin.PutRole)
		adminv1.GET("/roles/:id/permissions", admin.GetRolePerms)
		adminv1.POST("/roles/:id/permissions", admin.PostRolePerms)

		adminv1.GET("/menus", admin.GetMenus)
		adminv1.POST("/menus", admin.PostMenus)
		adminv1.DELETE("/menus/:id", admin.DeleteMenus)
		adminv1.PUT("/menus/:id", admin.PutMenus)

		adminv1.GET("/submenus", admin.GetSubMenu)
		adminv1.POST("/submenus", admin.PostSubMenu)
		adminv1.PUT("/submenus/:id", admin.PutSubMenus)
		adminv1.DELETE("/submenus/:id", admin.DeleteMenus)

		adminv1.GET("/domain/info", admin.GetDomainInfo)
		adminv1.POST("/domain/info", admin.AddDomainInfo)
		adminv1.PUT("/domain/info/:id", admin.PutDomainInfo)
		adminv1.DELETE("/domain/info/:id", admin.DelDomainInfo)

		adminv1.GET("/domain/cert", admin.GetDomainCret)
		adminv1.POST("/domain/cert", admin.AddDomainCret)
		adminv1.PUT("/domain/cert/:id", admin.PutDomainCret)
		adminv1.DELETE("/domain/cert/:id", admin.DelDomainCret)

		adminv1.GET("/host/role", admin.GetHostRole)
		adminv1.POST("/host/role", admin.AddHostRole)
		adminv1.PUT("/host/role/:id", admin.PutHostRole)
		adminv1.DELETE("/host/role/:id", admin.DelHostRole)

		adminv1.GET("/host", admin.GetHost)
		adminv1.POST("/host", admin.AddHost)
		adminv1.PUT("/hosts/:id", admin.PutHost)
		adminv1.DELETE("/hosts/:id", admin.DelHost)

		adminv1.GET("/host/app", admin.GetHostApp)
		adminv1.POST("/host/app", admin.AddHostApp)
		adminv1.PUT("/host/app/:id", admin.PutHostApp)
		adminv1.DELETE("/host/app/:id", admin.DelHostApp)
		adminv1.GET("/host/appid", admin.GetHostByAppId)

		adminv1.GET("/config/env", admin.GetConfigEnv)
		adminv1.POST("/config/env", admin.AddConfigEnv)
		adminv1.PUT("/config/env/:id", admin.PutConfigEnv)
		adminv1.DELETE("/config/env/:id", admin.DelConfigEnv)

		adminv1.GET("/config/type", admin.GetAppType)
		adminv1.POST("/config/type", admin.AddAppType)
		adminv1.PUT("/config/type/:id", admin.PutAppType)
		adminv1.DELETE("/config/type/:id", admin.DelAppType)

		adminv1.GET("/config/app", admin.GetConfigApp)
		adminv1.POST("/config/app", admin.AddConfigApp)
		adminv1.PUT("/config/app/:id", admin.PutConfigApp)
		adminv1.DELETE("/config/app/:id", admin.DelConfigApp)
		adminv1.GET("/config/template", admin.GetAppTemplate)

		adminv1.GET("/config/value", admin.GetAppValue)
		adminv1.POST("/config/value", admin.AddAppValue)
		adminv1.PUT("/config/value/:id", admin.PutAppValue)
		adminv1.DELETE("/config/value/:id", admin.DelAppValue)

		adminv1.GET("/config/deploy", admin.GetDeployExtend)
		adminv1.POST("/config/deploy", admin.AddDeployExtend)
		adminv1.PUT("/config/deploy/:id", admin.PutDeployExtend)
		adminv1.DELETE("/config/deploy/:id", admin.DelDeployExtend)

		adminv1.GET("/deploy/app", admin.GetAppDeploy)
		adminv1.POST("/deploy/app", admin.AddAppDeploy)
		adminv1.PUT("/deploy/app/:id", admin.PutAppDeploy)
		adminv1.DELETE("/deploy/app/:id", admin.DelAppDeploy)
		adminv1.PUT("/deploy/app/:id/review/:status", admin.PutAppDeployStatus)
		//adminv1.PUT("/deploy/app/:id/undo/:status", admin.PutAppDeployUndo)
		adminv1.GET("/deploy/app/:id/branch", admin.GetGitBranch)
		adminv1.GET("/deploy/app/:id/commit/:branch", admin.GetGitCommit)
	}

	return r
}
