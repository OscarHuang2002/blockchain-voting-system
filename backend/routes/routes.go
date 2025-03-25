package routes

import (
	"backend/controllers"

	"github.com/gin-gonic/gin"
)

func SetupRouter(r *gin.Engine) {
	// 用户身份管理
	r.POST("/register", controllers.RegisterVoter)
	r.POST("/login", controllers.Login)
	// 投票
}
