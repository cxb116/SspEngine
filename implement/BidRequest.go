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

type BidRequest struct {
	RequestId   string `json:"request_id"`  // 本次请求生成的请求id用于日志排查
	SspSlotId   int32  `json:"ssp_slot_id"` // 广告位Id
	AppId       string `json:"app_id"`      // AppId
	BidFloor    int32  `json:"bid_floor"`
	Geo         Geo    `json:"geo"`
	App         App    `json:"app"`
	Device      Device `json:"device"`
	Caid        []Caid `json:"caid"`
	RequestTime int64  `json:"request_time"` // 请求打到ssp的时候的时间戳,用于统计响应时间
}
type Caid struct {
	Caid    string `json:"caid"`
	CaidVer string `json:"caid_ver"`
}

type App struct {
	AppName  string `json:"app_name"`  // 引用名称
	Bundle   string `json:"bundle"`    // 应用信息或者包名
	Domain   string `json:"domain"`    // 应用域名
	StoreUrl string `json:"store_url"` // 应用商店地址
	Version  string `json:"version"`   // 应用版本号
}

type Device struct {
	Ua             string  `json:"ua"`          // 浏览器User-Agent字符串
	Ip             string  `json:"ip"`          // 最接近设备的IPv4地址
	Ipv6           string  `json:"ipv6"`        // 最接近设备的IPV6地址
	DeviceType     string  `json:"device_type"` // 1 手机 2 平板 3 联网设备 4 机顶盒
	Make           string  `json:"make"`        // 设备制造商，例如 "Apple"
	Model          string  `json:"model"`       // 设备型号，例如 "iphone"
	Os             string  `json:"os"`          // 设备操作系统， 例如 “ios","android"
	Osv            string  `json:"os_version"`  // 设备操作系统版本号， 例如 “3.1.2”
	Hwv            string  `json:"hwv"`         // 设备硬件版本， 例如 “5S”
	H              int     `json:"h"`           // 屏幕的物理高度， 以像素为单位
	W              int     `json:"w"`           // 屏幕的物理宽度，以像素为单位
	Ppi            int     `json:"ppi"`         // 以像素每英寸表示的屏幕尺寸
	Pxratio        float64 `json:"pxratio"`     // 设备物理像素与设备无关像素的比率
	Language       string  `json:"language"`    // 浏览器语言
	Carrier        string  `json:"carrier"`
	ConnectionType string  `json:"connection_type"` // 1 wifi, 2 2G, 3 3G,4 4G,5 5G
	Imei           string  `json:"imei"`            // imei信息
	AndroidId      string  `json:"android_id"`      // 安卓id
	Mac            string  `json:"mac"`             // 设备mac地址

}
type Geo struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}

func (this *BidRequest) GetRequestId() string {
	return this.RequestId
}
func (this *BidRequest) SetRequestId(requestId string) {
	this.RequestId = requestId
}
func (this *BidRequest) GetSspSlotId() int32 {
	return this.SspSlotId
}
func (this *BidRequest) SetSspSlotId(sspSlotId int32) {
	this.SspSlotId = sspSlotId
}
func (this *BidRequest) GetAppId() string {
	return this.AppId
}
func (this *BidRequest) SetAppId(appId string) {
	this.AppId = appId
}
func (this *BidRequest) GetRequestTime() int64 {
	return this.RequestTime
}
func (this *BidRequest) SetRequestTime(requestTime int64) {
	this.RequestTime = requestTime
}

//func (this *BidRequest) GetRequestId() string {
//	return this.RequestId
//}
//
//func (this *BidRequest) GetSspSlotId() int32 {
//	return this.SspSlotId
//}
//
//func (this *BidRequest) GetAppId() string {
//	return this.AppId
//}
//
//func (this *BidRequest) GetRequestTime() int64 {
//	return this.RequestTime
//}
//
//func (this *BidRequest) SetRequestId(requestId string) {
//	this.RequestId = requestId
//}
//
//func (this *BidRequest) SetSspSlotId(ssp_slot_id int32) {
//	this.SspSlotId = ssp_slot_id
//}
//
//func (this *BidRequest) SetAppId(app_id string) {
//	this.AppId = app_id
//}
//
//func (this *BidRequest) SetRequestTime(requestTime int64) {
//	this.RequestTime = requestTime
//}

func NewBidRequest() *BidRequest {
	return &BidRequest{}
}
