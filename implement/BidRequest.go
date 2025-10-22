package implement

import (
	"github.com/cxb116/sspEngine/interfaces"
	"sync"
)

// 媒体请求ssp携带的json
var BidRequestPool = new(sync.Pool)

func init() {
	BidRequestPool.New = func() interface{} {
		return &BidRequest{}
	}
}

type BidRequest struct {
	RequestId string `json:"request_id"`  // 本次请求生成的请求id用于日志排查
	SspSlotId int32  `json:"ssp_slot_id"` // 广告位Id,创建广告位,将广告位id发给媒体,媒体会携带着去请求ssp服务器
	AppId     string `json:"app_id"`

	RequestTime int64 `json:"request_time"` // 请求打到ssp的时候的时间戳,用于统计响应时间
}

func NewBidRequest() *BidRequest {
	return &BidRequest{}
}

// 获取BidRequest 池对象
func GetBidRequest() *BidRequest {
	request := BidRequestPool.Get().(*BidRequest)
	request.Reset()
	//request.RequestTime = time.Now().UnixMilli() // 获取当前毫秒
	return request
}

// 归还BidRequest对象
func PutBidRequest(request interfaces.IBidRequest) {
	if request != nil {
		BidRequestPool.Put(request)
	}
}

// 将每一个数据置空
func (this *BidRequest) Reset() {
	this.RequestId = ""
	this.AppId = ""
}
