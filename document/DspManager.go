package document

import (
	"fmt"
	"github.com/cxb116/sspEngine/interfaces"
)

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// ##################################################如何高效的将流量分派到各个预算 #############################################//
//
// 思路：2 备用方案 通过管理端自定义预算名称绑定 (频繁在DspMaps中添加删除的修改感觉会有问题，所以在初始化的时候就将所有预算全部添加进去)
// 目标：RequestHandler通过BidRequest,在DspManager去调度预算，dspManager可以像插件一样实现插拔。
// 大体思路：
//		1 RequestHandler 传入BidRequest,根据sspSlotId来匹配绑定的预算 (在RequestHandler中)
//      2 首先在初始化时获取管理端配置数据，存入Dsp中SspSlotInfoMaps
//      3 每个预算都会加一个dspCode用来和管理端配置的dspCode去匹配，然后可以通过继承Dsp子类就可以获取到所有和dspCode有关的管理端配置数据
//
///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

func DspDispatchManager(req interfaces.IBidRequest, slot interfaces.ISspSlotInfo) (interfaces.IBidResponse, error) {
	registryHandler := GetDspDoc(slot.GetDspCode(), req, slot)
	if registryHandler == nil {
		return nil, fmt.Errorf("未找到 DSP: %s", slot.GetDspCode())
	}
	return registryHandler.RequestBid()
}

type DspHandlerCreator func(request interfaces.IBidRequest, slot interfaces.ISspSlotInfo) DspHandler

type DspHandler interface {
	RequestBid() (interfaces.IBidResponse, error)
	GetDspCode() string
}

var dspRegistryMaps = make(map[string]DspHandlerCreator)

func DspRegister(dspCode string, creator DspHandlerCreator) {
	if _, exist := dspRegistryMaps[dspCode]; !exist {
		dspRegistryMaps[dspCode] = creator
	}
}

// 通过dspCode去获取 BidRequest和SspSlotInfo
// 获取预算对接文档
func GetDspDoc(dspCode string, request interfaces.IBidRequest, slot interfaces.ISspSlotInfo) DspHandler {
	if creator, exist := dspRegistryMaps[dspCode]; exist {
		return creator(request, slot)
	}
	return nil
}

//func initDSP() {
//	RegisterDsp("Baidu", func(req interfaces.IBidRequest, slot *SspSlotInfo) interfaces.IDspHandlerManager {
//		return dsp.NewDspBaiduDsp(req, slot)
//	})
//}
//
//var dspRegistry = make(map[string]func(req interfaces.IBidRequest, slot *SspSlotInfo) interfaces.IDspHandlerManager)
//
//// 通过 dspCode 将handler 注册到 dspRegistry 中
//func RegisterDsp(dspCode string, register func(req interfaces.IBidRequest, slot *SspSlotInfo) interfaces.IDspHandlerManager) {
//	dspRegistry[dspCode] = register
//}
//
//// 通过dspCode获取 对接文档的信息
//func GetDspHandler(dspCode string, request interfaces.IBidRequest, slot *SspSlotInfo) interfaces.IDspHandlerManager {
//	if dspRegistry[dspCode] != nil {
//		return dspRegistry[dspCode](request, slot)
//	}
//	return nil
//}
//
//type DspHandlerManager struct {
//	DspCode     string
//	BidRequest  interfaces.IBidRequest
//	SspSlotInfo *SspSlotInfo
//}
//
//// 请求响应在这一个方法里面转换
//func (this *DspHandlerManager) RequestBid() (interfaces.IBidResponse, error) {
//	return &BidRequest{}, errors.New("先退出吧")
//}
//
//func (this *DspHandlerManager) GetDspCode() string {
//
//	return this.DspCode
//}

//// 先写一个dsp管理者////////////////////////////////////////////////////////////////
//type DspManager struct {
//	DspMaps     map[int32]interfaces.IDsp
//	DspMgrMutex sync.RWMutex
//}
//
//var (
//	DspMgr     *DspManager
//	DspMgrOnce sync.Once
//)
//
//// 单例 NewDspManger
//func NewDspManager() *DspManager {
//	DspMgrOnce.Do(func() {
//		DspMgr = &DspManager{
//			DspMaps: make(map[int32]interfaces.IDsp),
//		}
//	})
//
//	return DspMgr
//}
//
//// 注册DSP插件
//func (this *DspManager) RegisterDsp(dsp interfaces.IDsp) {
//	this.DspMgrMutex.Lock()
//	defer this.DspMgrMutex.Unlock()
//	//sspSlotId := request.(BidRequest).SspSlotId
//	//sspSlotInfo := dsp.(Dsp).SspSlotInfoMaps[sspSlotId]
//	//if sspSlotInfo != nil {
//	//	this.DspMaps[sspSlotId] = dsp
//	//}
//}
//
//// 根据request 匹配合适的DSP并发起竞价
//func (this *DspManager) DispatchDsp(request interfaces.IBidRequest) []interfaces.IBidResponse {
//	this.DspMgrMutex.RLock()
//	defer this.DspMgrMutex.RUnlock()
//
//	var bidResponses []interfaces.IBidResponse
//	//for _, dsp := range this.DspMaps {
//	//	if dsp.Match(request) { // 匹配
//	//		resp, err := dsp.Bid(request) // 发起dsp的请求 响应等
//	//		if err != nil && resp != nil {
//	//			bidResponses = append(bidResponses, resp) // TODO 这里可以创建channel进行处理
//	//		}
//	//	}
//	//}
//	return bidResponses
//}
