package dsp

import (
	"github.com/cxb116/sspEngine/implement"
	"github.com/cxb116/sspEngine/interfaces"
)

type BaiduDsp struct {
	implement.Dsp
	//interfaces.IDsp
}

/**
  1 首先dsp 收集前置数据，供应dsp请求,过滤 等...
	1) BidRequest
    2) 管理端的配置数据，这些数据控制dsp的决策，例如dsp,请求地址，请求方式，响应时间，是否支持deeplink 等等..
*/

func (this *BaiduDsp) CreateReqMsg() string {
	return "CreateReqMsg"
}

func (dsp *BaiduDsp) CreateRes() { // 构建预算返回的物料，转化成我方返回媒体的信息

}

func (dsp *BaiduDsp) SendBidMsg() bool { // 请求dsp服务器
	return false
}

func (dsp *BaiduDsp) GetDspCost() (int64, int64) { // 获取预算请求耗时，毫秒
	return 0, 0
}

func (dsp *BaiduDsp) GetRes() interfaces.IBidResponse { // 将物料转化成我方文档数据，可通过getRes获取
	return nil
}

func (dsp *BaiduDsp) GetBidId() string { // 获取bidRequestId
	return ""
}
