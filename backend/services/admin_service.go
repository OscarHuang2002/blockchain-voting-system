package services

import (
	"context"
	"fmt"
	"log"
	"math/big"
)

// 添加候选人
func (cs *ContractService) AddCandidate(projectID uint, name string, imageUrl string, description string) error {
    if cs == nil || cs.contract == nil {
        return ErrContractNotInitialized
    }

    tx, err := cs.contract.AddCandidate(
        cs.auth,
        big.NewInt(int64(projectID)),
        name,
        imageUrl,
        description,
    )
    if err != nil {
        log.Printf("添加候选人失败: %v", err)
        return err
    }
    
    // 等待交易完成
    _, err = cs.client.TransactionReceipt(context.Background(), tx.Hash())
    if err != nil {
        log.Printf("等待交易完成失败: %v", err)
        return err
    }

    return nil
}

// 修改候选人
func (cs *ContractService) UpdateCandidate(projectID uint, candidateID uint, name string, imageUrl string, description string) error {
    if cs == nil || cs.contract == nil {
        return ErrContractNotInitialized
    }

    // 检查候选人状态 - 使用下划线忽略未使用的返回值
    _, _, _, _, _, candidateActive, err := cs.contract.GetCandidate(nil, big.NewInt(int64(projectID)), big.NewInt(int64(candidateID)))
    if err != nil {
        log.Printf("获取候选人信息失败: %v", err)
        return err
    }

    // 检查候选人是否活跃
    if !candidateActive { 
        log.Printf("候选人状态为非活跃，无法修改: 项目ID=%d, 候选人ID=%d", projectID, candidateID)
        return fmt.Errorf("候选人状态为非活跃，无法修改")
    }

    // 调用智能合约的 updateCandidate 方法
    tx, err := cs.contract.UpdateCandidate(
        cs.auth,
        big.NewInt(int64(projectID)),
        big.NewInt(int64(candidateID)),
        name,
        imageUrl,
        description,
    )
    if err != nil {
        log.Printf("修改候选人失败: %v", err)
        return err
    }

    // 等待交易完成
    _, err = cs.client.TransactionReceipt(context.Background(), tx.Hash())
    if err != nil {
        log.Printf("等待交易完成失败: %v", err)
        return err
    }

    return nil
}

// 删除候选人
func (cs *ContractService) DeleteCandidate(projectID uint, candidateID uint) error {
    if cs == nil || cs.contract == nil {
        return ErrContractNotInitialized
    }

    // 调用智能合约的 deleteCandidate 方法
    tx, err := cs.contract.DeleteCandidate(
        cs.auth, 
        big.NewInt(int64(projectID)), 
        big.NewInt(int64(candidateID)),
    )
    if err != nil {
        log.Printf("删除候选人失败: %v", err)
        return err
    }

    // 等待交易完成
    _, err = cs.client.TransactionReceipt(context.Background(), tx.Hash())
    if err != nil {
        log.Printf("等待交易完成失败: %v", err)
        return err
    }

    return nil
}

// 开始投票
func (cs *ContractService) StartVoting(projectID uint) error {
    if cs == nil || cs.contract == nil {
        return ErrContractNotInitialized
    }

    // 调用智能合约的 startVotingForProject 方法
    tx, err := cs.contract.StartVotingForProject(cs.auth, big.NewInt(int64(projectID)))
    if err != nil {
        log.Printf("开始投票失败: %v", err)
        return err
    }

    // 等待交易完成
    _, err = cs.client.TransactionReceipt(context.Background(), tx.Hash())
    if err != nil {
        log.Printf("等待交易完成失败: %v", err)
        return err
    }

    return nil
}

// 结束投票
func (cs *ContractService) EndVoting(projectID uint) error {
    if cs == nil || cs.contract == nil {
        return ErrContractNotInitialized
    }

    // 调用智能合约的 endVotingForProject 方法
    tx, err := cs.contract.EndVotingForProject(cs.auth, big.NewInt(int64(projectID)))
    if err != nil {
        log.Printf("结束投票失败: %v", err)
        return err
    }

    // 等待交易完成
    _, err = cs.client.TransactionReceipt(context.Background(), tx.Hash())
    if err != nil {
        log.Printf("等待交易完成失败: %v", err)
        return err
    }

    return nil
}

// 创建投票项目 (新增功能)
func (cs *ContractService) CreateVotingProject(description string) error {
    if cs == nil || cs.contract == nil {
        return ErrContractNotInitialized
    }

    tx, err := cs.contract.CreateVotingProject(cs.auth, description)
    if err != nil {
        log.Printf("创建投票项目失败: %v", err)
        return err
    }

    // 等待交易完成
    _, err = cs.client.TransactionReceipt(context.Background(), tx.Hash())
    if err != nil {
        log.Printf("等待交易完成失败: %v", err)
        return err
    }

    return nil
}

// 获取项目总数 (新增功能)
func (cs *ContractService) GetProjectCount() (uint, error) {
    if cs == nil || cs.contract == nil {
        return 0, ErrContractNotInitialized
    }

    count, err := cs.contract.ProjectCount(nil)
    if err != nil {
        log.Printf("获取项目总数失败: %v", err)
        return 0, err
    }

    return uint(count.Uint64()), nil
}

// 获取项目信息 (新增功能)
func (cs *ContractService) GetProjectInfo(projectID uint) (map[string]interface{}, error) {
    if cs == nil || cs.contract == nil {
        return nil, ErrContractNotInitialized
    }

    // 修改为接收5个返回值
    id, description, isActive, isDeleted, err := cs.contract.GetProjectInfo(nil, big.NewInt(int64(projectID)))
    if err != nil {
        log.Printf("获取项目信息失败: %v", err)
        return nil, err
    }

    return map[string]interface{}{
        "id":          id.Uint64(),
        "description": description,
        "isActive":    isActive,
        "isDeleted":   isDeleted,
    }, nil
}

// 删除项目 (新增功能)
func (cs *ContractService) DeleteProject(projectID uint) error {
    if cs == nil || cs.contract == nil {
        return ErrContractNotInitialized
    }

    tx, err := cs.contract.DeleteProject(cs.auth, big.NewInt(int64(projectID)))
    if err != nil {
        log.Printf("删除项目失败: %v", err)
        return err
    }

    // 等待交易完成
    _, err = cs.client.TransactionReceipt(context.Background(), tx.Hash())
    if err != nil {
        log.Printf("等待交易完成失败: %v", err)
        return err
    }

    return nil
}