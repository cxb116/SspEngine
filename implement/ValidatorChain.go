package implement

import "sync"

// 拦截策略
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
