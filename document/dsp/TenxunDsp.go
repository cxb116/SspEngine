package dsp

import (
	"github.com/cxb116/sspEngine/implement"
	"github.com/cxb116/sspEngine/interfaces"
	"github.com/rs/zerolog/log"
)

type TenxunDsp struct {
	implement.DspManager
}

func (b *TenxunDsp) GetDspId() int32 {
	log.Print("获取预算TenxunDsp的DspId")
	return 0
}
func (b *TenxunDsp) Match(bidRequest interfaces.IBidRequest) bool {
	maps := b.DspMaps // 拿到了注册的预算map
	log.Print("获取TenxunDsp 匹配信息")
	return true
}

//
func (b *TenxunDsp) Bid(bidRequest interfaces.IBidRequest) (interfaces.IBidResponse, error) {
	log.Print("TenxunDsp 请求转化完成，返回IBidResponse 对象")
	BidResponse := &implement.BidResponse{}
	return BidResponse, nil
}
