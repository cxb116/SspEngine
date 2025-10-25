package dsp

import (
	"errors"
	"fmt"
	"github.com/cxb116/sspEngine/document"
	"github.com/cxb116/sspEngine/implement"
	"github.com/cxb116/sspEngine/interfaces"
	"github.com/rs/zerolog/log"
)

func init() {
	log.Printf("BaiduDsp 被执行了")
	document.DspRegister("baidu", func(request interfaces.IBidRequest, slot interfaces.ISspSlotInfo) document.DspHandler {
		return &BaiduDsp{
			DspCode:    "baidu",
			BidRequest: request,
			Slot:       slot,
		}
	})
}

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
