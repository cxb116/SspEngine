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
	RequestId   string `json:"request_id"`
	RequestTime int64  `json:"request_time"` // 请求打到ssp的时候的时间戳
	AppId       string `json:"app_id"`
}

func NewBidRequest() *BidRequest {
	return &BidRequest{}
}

// 获取BidRequest 池对象
func GetBidRequest() *BidRequest {
	request := BidRequestPool.Get().(*BidRequest)
	request.Reset()
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
