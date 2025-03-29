package routes

import (
	"backend/config"
	"backend/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	//初始化登录服务
	controllers.InitLoginService(config.DB)
	// 用户身份管理
	r.POST("/register", controllers.RegisterVoter)
	r.POST("/login", controllers.LoginWithAddress)
	// 投票
}
