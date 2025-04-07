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
    
    // 用户投票界面
    r.POST("/vote", controllers.CastVote)
    r.GET("/userVoteDetail", controllers.GetUserVoteDetail)
    r.GET("/userVoteHistory", controllers.GetUserVoteDetails) // 添加新的路由
    r.GET("/candidates", controllers.GetAllCandidates) 
	// 在现有路由的基础上添加
	r.GET("/verifyVote", controllers.VerifyVote)
    
    // 项目相关
    r.GET("/projects", controllers.GetProjects)

    r.POST("/upload/image", controllers.UploadImage)
    // 设置静态文件服务，用于访问上传的图片
    r.Static("/images", "./uploads/images")
    
    
    // 管理员操作
    admin := r.Group("/AdminPanel")
    {
        admin.POST("/createProject", controllers.CreateVotingProject)
        admin.POST("/addCandidate", controllers.AddCandidate)
        admin.PUT("/updateCandidate", controllers.UpdateCandidate)
        admin.DELETE("/deleteCandidate/:id", controllers.DeleteCandidate)
        admin.POST("/startVoting", controllers.StartVoting)
        admin.POST("/endVoting", controllers.EndVoting)
        admin.DELETE("/deleteProject/:id", controllers.DeleteProject)
        // admin.POST("/upload/image", controllers.UploadImage)
    }
}