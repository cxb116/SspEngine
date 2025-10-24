package dsp

import (
	"errors"
	"fmt"
	"github.com/cxb116/sspEngine/implement"
	"github.com/cxb116/sspEngine/implement/registry"
	"github.com/cxb116/sspEngine/interfaces"
)

type BaiduDsp struct {
	DspCode    string
	BidRequest interfaces.IBidRequest
	Slot       interfaces.ISspSlotInfo
}

func (dsp *BaiduDsp) RequestBid() (interfaces.IBidResponse, error) {
	fmt.Printf("调用百度 DSP 请求逻辑...")
	return &implement.BidResponse{}, errors.New("百度调用异常")
}
func (dsp *BaiduDsp) GetDspCode() string {
	return dsp.DspCode
}

func init() {
	registry.Register("baidu", func(request interfaces.IBidRequest, slot interfaces.ISspSlotInfo) registry.DspHandler {
		return &BaiduDsp{
			DspCode:    "baidu",
			BidRequest: request,
			Slot:       slot,
		}
	})
}
