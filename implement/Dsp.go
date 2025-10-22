package implement

import "github.com/cxb116/sspEngine/interfaces"

type Dsp struct {
	SspSlotInfoMaps map[int32]*SspSlotInfo  // key=SspSlotId value=*SspSlotInfo
	DspRequest      interfaces.IDspRequest  // 请求dsp
	DspResponse     interfaces.IDspResponse // dsp响应
	BidRequest      interfaces.IDspRequest  // 请求ssp
	BidResponse     interfaces.IDspResponse // ssp响应

}

//func (dsp *Dsp) GetDspId(request interfaces.IBidRequest) int32 {
//	sspSlotId := request.(BidRequest).SspSlotId
//	if dsp.SspSlotInfoMaps != nil && dsp.SspSlotInfoMaps[sspSlotId] != nil {
//		return sspSlotId
//	}
//	return -1
//}

func (dsp *Dsp) CreateReqMsg() string { // 获取请求+ 管理端配置的信息
	return "CreateReqMsg"
}
func (dsp *Dsp) CreateRes() { // 构建预算返回的物料，转化成我方返回媒体的信息

}
func (dsp *Dsp) SendBidMsg() bool { // 请求dsp服务器

	return false
}
func (dsp *Dsp) GetDspCost() (int64, int64) { // 获取预算请求耗时，毫秒
	return 0, 0
}

func (dsp *Dsp) GetRes() interfaces.IBidResponse { // 将物料转化成我方文档数据，可通过getRes获取
	return nil
}
func (dsp *Dsp) GetBidId() string { // 获取bidRequestId
	return ""
}
