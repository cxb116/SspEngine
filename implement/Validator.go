package implement

import (
	"errors"
	"fmt"
	"github.com/cxb116/sspEngine/interfaces"
	"github.com/rs/zerolog/log"
	"sync"
)

/**
 	媒体请求过滤策略
	策略 + 单例
*/

var (
	validatorRegistry = make(map[string]interfaces.IValidator)
	registryMutex     sync.RWMutex
	onceRegistry      sync.Once
)

func init() {
	onceRegistry.Do(registerBuiltins)
	RegisterChain("request_id", Chain{"request 长度不小于36"})
	fmt.Println("llllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllllll")
}

// 注册时内置validator 单例
func registerBuiltins() {
	registryMutex.Lock()
	defer registryMutex.Unlock()
	validatorRegistry["request_id"] = &RequestIdValidator{}
	validatorRegistry["app_id"] = &AppIdValidator{}
}

// 返回注册的验证器实例(单例)
func GetValidatorByName(name string) (interfaces.IValidator, bool) {
	onceRegistry.Do(registerBuiltins) // 延迟注册内置
	registryMutex.RLock()
	defer registryMutex.RUnlock()
	vali, ok := validatorRegistry[name]
	return vali, ok
}

// 允许动态注册覆盖同名validator
func RegisterValidator(name string, validator interfaces.IValidator) {
	onceRegistry.Do(registerBuiltins)
	registryMutex.Lock()
	defer registryMutex.Unlock()
	validatorRegistry[name] = validator
}

// requestId 过滤策略
type RequestIdValidator struct{}

func (validator RequestIdValidator) Validate(request interfaces.IBidRequest) (int, error) {
	req, ok := request.(*BidRequest)
	if !ok {
		return -1, errors.New("断言失败,不是 *BidRequest")
	}
	if len(req.RequestId) > 0 {
		log.Printf("request_id 验证通过: %s", request.(*BidRequest).RequestId)
		return 0, nil
	}

	return -1, errors.New("request_id 不合法")
}

type AppIdValidator struct{}

func (validator AppIdValidator) Validate(request interfaces.IBidRequest) (int, error) {
	req, ok := request.(*BidRequest)
	if !ok {
		return -1, errors.New("断言失败,不是 *BidRequest")
	}
	if len(req.AppId) > 0 {
		log.Printf("app_id 验证通过: %s", request.(*BidRequest).AppId)
		return 0, nil
	}
	return -1, errors.New("AppId 不合法")
}

func ValidateRequest(request interfaces.IBidRequest) (int, error) {
	validators := []interfaces.IValidator{
		&RequestIdValidator{},
		&AppIdValidator{},
		//TODO 添加各个请求字段的过滤逻辑
	}
	for _, validator := range validators {
		if code, err := validator.Validate(request); err == nil { // 发现不合法直接返回
			return code, nil
		} else {
			return code, err
		}
	}
	return -1, nil
}

// 指定验证,验证ua和model,看是否匹配
//	status, err := ValidateWithChain([]string{"request_id"}, request)
//	if err != nil {
//		return status, err
//	} else {
//		return status, err
//	}
// TODO 加入校验链,例如先判断model是否合理，然后判断ua,还有caid和idfa二选一情况
func ValidateWithChain(chain Chain, request interfaces.IBidRequest) (int, error) {
	onceRegistry.Do(registerBuiltins)
	for _, chainName := range chain {
		validator, ok := GetValidatorByName(chainName)
		if !ok {
			// 找不到注册的validator 配置错误
			return -1, fmt.Errorf("找不到: %s", chainName)
		}
		if code, err := validator.Validate(request); err == nil {
			return code, nil
		}
	}
	return -1, nil
}
