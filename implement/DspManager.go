package implement

import (
	"github.com/cxb116/sspEngine/interfaces"
	"sync"
)

type DspManager struct {
	dsps     []interfaces.IDSP
	dspMutex sync.RWMutex
}

var dspMgr = &DspManager{}

func GetDspManager() *DspManager {
	return dspMgr
}

// 注册DSP插件
func (this *DspManager) RegisterDsp(dsp interfaces.IDSP) {
	this.dspMutex.Lock()
	defer this.dspMutex.Unlock()
	this.dsps = append(this.dsps, dsp)
}

// 根据request 匹配合适的DSP并发起竞价
func (this *DspManager) DispatchDsp(request interfaces.IBidRequest) []interfaces.IBidResponse {
	this.dspMutex.RLock()
	defer this.dspMutex.RUnlock()

	var bidResponses []interfaces.IBidResponse
	for _, dsp := range this.dsps {
		if dsp.Match(bidResponses) { // 匹配
			resp, err := dsp.Bid(request) // 发起dsp的请求 响应等
			if err != nil && resp != nil {
				bidResponses = append(bidResponses, resp) // TODO 这里可以创建channel进行处理
			}
		}
	}
	return bidResponses
}
