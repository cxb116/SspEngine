package implement

import "sync"

//TODO 还有问题，后续要在chain中输入ua和model,然后根据model 判定ua是否一致且合法,caid和idfa,等参数联合校验等问题
// BidRequest 验证是否合格的策略链路

type Chain []string

var (
	ChainRegistry = make(map[string]Chain)
	ChainMutex    sync.RWMutex
)

// 添加注册
func RegisterChain(key string, chain Chain) {
	ChainMutex.Lock()
	defer ChainMutex.Unlock()
	ChainRegistry[key] = chain
}

// 根据key获取策略
func GetChain(key string) Chain {
	ChainMutex.RLock()
	defer ChainMutex.RUnlock()
	return ChainRegistry[key]
}
