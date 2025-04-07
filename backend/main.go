package main

import (
	"backend/config"
	"backend/controllers"
	"backend/routes"
	"backend/services"
	"log"

	"github.com/gin-contrib/cors" // 引入 CORS 中间件
	"github.com/gin-gonic/gin"
)

func main() {
	// 初始化数据库
	config.InitDB()

	// 初始化登录服务
	controllers.InitLoginService(config.DB)

	// 初始化 ContractService
	clientURL := "http://127.0.0.1:8545"                                             // 替换为你的以太坊节点 URL
	privateKey := "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80" // 替换为你的钱包私钥
	contractAddress := "0x5FbDB2315678afecb367f032d93F642f64180aa3"                  // 替换为你的智能合约地址

	contractService, err := services.NewContractService(clientURL, privateKey, contractAddress)
	if err != nil {
    	log.Fatalf("Failed to initialize ContractService: %v", err)
	}

	// 使用全局初始化函数
	services.InitGlobalContractService(contractService)
	services.InitContractService(contractService)
	controllers.InitDependencies(contractService, config.DB)

	// 设置 Gin 路由
	r := gin.Default()

	// 启用 CORS 中间件
	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},                   // 允许来自前端的请求
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}, // 允许的 HTTP 方法
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization", "X-Requested-With"}, // 允许的请求头
		ExposeHeaders:    []string{"Content-Length"},                          // 允许暴露的响应头
		AllowCredentials: true,                                                // 允许跨域携带 Cookie
	}))

	// 设置路由并运行服务
	routes.SetupRouter(r)
	r.Run(":8080") // 在 8080 端口启动服务
}
