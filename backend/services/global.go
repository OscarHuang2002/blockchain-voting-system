package services

import (
	"log"
)

// GlobalContractService 是全局共享的合约服务实例
var GlobalContractService *ContractService

// InitGlobalContractService 初始化全局合约服务实例
func InitGlobalContractService(service *ContractService) {
    GlobalContractService = service
    log.Println("全局合约服务已初始化")
}

// GetGlobalContractService 获取全局合约服务实例
func GetGlobalContractService() *ContractService {
    return GlobalContractService
}