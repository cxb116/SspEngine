package dsp

import (
	"errors"
	"github.com/cxb116/sspEngine/implement"
	"github.com/cxb116/sspEngine/interfaces"
)

type BaiduDsp struct {
	*implement.DspHandlerManager
}

func NewDspBaiduDsp() *BaiduDsp {
	return &BaiduDsp{
		DspHandlerManager: &implement.DspHandlerManager{
			DspCode: "Baidu",
		},
	}
}

func (handler *BaiduDsp) RequestBid(request interfaces.IBidRequest) (interfaces.IBidResponse, error) {
	return nil, errors.New("BaiduDsp失败")
}

func (this *BaiduDsp) GetDspCode() string {
	return this.DspCode
}
