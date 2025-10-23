package interfaces

//  媒体-> ssp 流量自带 sspSlotId, appId
//     -> ssp 引擎 (匹配预算)
//		  -> Dsp 预算接口(dspCode url method budgetSlotId)
//
//   SSP 流量（sspSlotId）和 DSP 预算位（budgetSlotId）绑定
//   根据 dspCode 找到对接文档
//   在请求时快速路由到正确的预算 handler
//   流量分发路由  +   策略映射 + 动态扩展

type IDspHandlerManager interface {
	RequestBid(reqeuest IBidRequest) (IBidResponse, error)

	GetDspCode() string
}

//type IDspManager interface {
//	GetDspId() int32
//	Match(bidRequest IBidRequest) bool
//	Bid(bidRequest IBidRequest) (IBidResponse, error)
//}
