package implement

import (
	"errors"
	"github.com/cxb116/sspEngine/interfaces"
)

//////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// ##################################################如何高效的将流量分派到各个预算 #############################################//
//思路：1将预算注册到maps中，然后通过广告位置id来绑定这个预算 (这个有点冲突感觉有点麻烦)
//      2 备用方案 通过管理端自定义预算名称绑定,
// 具体步骤：
//   	1 创建dspManager 来存储dsp预算，DspManager 不能防止频繁创建销毁这个对象，因使用单例模式
//          1) dsp如何实现: dsp文档对接需要管理端的配置数据等，还需要预算请求体，预算响应体，预算响应时间(防止超时)
//   	    2) dsp实例化，首先在main加载之前将管理端设备信息加载到 Dsp 对象的SspSlotInfoMaps map[int64]*SspSlotInfo,中
//          3) 例如要对接TenxunDsp文档,只需要继承dsp,然后重写dsp方法，能够实现文档的对接功能
//          4）对接完成后,将dsp加入DspManager中的dspMaps,通过请求id检索然后执行id 所对应的预算
//	 	2 有个问题是没有和媒体绑定的dsp预算是没有广告位id的
//     解决办法: dsp配置信息是通过管理端来配置的,我可以在管理端配置信息后通过redis的订阅发布者模式来推动引擎跟新DspManager信息
//
// 思路：2 备用方案 通过管理端自定义预算名称绑定 (频繁在DspMaps中添加删除的修改感觉会有问题，所以在初始化的时候就将所有预算全部添加进去)
// 目标：RequestHandler通过BidRequest,在DspManager去调度预算，dspManager可以像插件一样实现插拔。
// 大体思路：
//		1 RequestHandler 传入BidRequest,根据sspSlotId来匹配绑定的预算 (在RequestHandler中)
//      2 首先在初始化时获取管理端配置数据，存入Dsp中SspSlotInfoMaps
//      3 每个预算都会加一个dspCode用来和管理端配置的dspCode去匹配，然后可以通过继承Dsp子类就可以获取到所有和dspCode有关的管理端配置数据
//
///////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////

var dspRegistry = make(map[string]interfaces.IDspHandlerManager)

// 通过 dspCode 将handler 注册到 dspRegistry 中
func RegisterDsp(handler interfaces.IDspHandlerManager) {
	dspRegistry[handler.GetDspCode()] = handler
}

// 通过dspCode获取 对接文档的信息
func GetDspHandler(dspCode string) interfaces.IDspHandlerManager {
	return dspRegistry[dspCode]
}

type DspHandlerManager struct {
	DspCode string
}

// 请求响应在这一个方法里面转换吗
func (this *DspHandlerManager) RequestBid(reqeuest interfaces.IBidRequest) (interfaces.IBidResponse, error) {
	return &BidRequest{}, errors.New("先退出吧")
}

func (this *DspHandlerManager) GetDspCode() string {

	return this.DspCode
}

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
