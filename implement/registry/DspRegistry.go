package registry

import (
	"github.com/cxb116/sspEngine/interfaces"
)

// 工厂函数类型
type HandlerCreator func(request interfaces.IBidRequest, slot interfaces.ISspSlotInfo) DspHandler

type DspHandler interface {
	RequestBid() (interfaces.IBidResponse, error)
	GetDspCode() string
}

var dspRegistry = make(map[string]HandlerCreator)

func Register(dspCode string, creator HandlerCreator) {
	if _, exist := dspRegistry[dspCode]; !exist {
		dspRegistry[dspCode] = creator
	}
}

// 通过dspCode去获取 BidRequest  和 SspSlotInfo
func Get(dspCode string, request interfaces.IBidRequest, slot interfaces.ISspSlotInfo) DspHandler {
	if creator, exist := dspRegistry[dspCode]; exist {
		return creator(request, slot)
	}
	return nil
}
