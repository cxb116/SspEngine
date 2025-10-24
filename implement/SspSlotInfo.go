package implement

import (
	"fmt"
	"sync"
)

//var SspSlotInfoMaps = make(map[int32]*SspSlotInfo)

var (
	GSspSlotInfoBindings *SspSlotInfoBindings
	once                 sync.Once
)

func GetSspSlotInfoBindings() *SspSlotInfoBindings {
	once.Do(func() {
		GSspSlotInfoBindings = &SspSlotInfoBindings{}
		GSspSlotInfoBindings.Init()
	})
	return GSspSlotInfoBindings
}

// 初始化绑定隐射，可重复调用刷新数据
func (this *SspSlotInfoBindings) Init() {
	this.UpdateBinding()
}

// 管理端配置的预算位置信息
type SspSlotInfo struct {
	SspSlotId                int32      // 广告位id
	RtaExt                   string     //
	AdType                   int32      // 广告属性
	OsType                   int32      // 系统属性 1=android 2=ios
	PayType                  int8       // 结算方式 1=固价 2=分成 3=RTB
	DealRatio                int32      // 成交价系数 单位%
	AppKey                   string     // 预算分派的key
	AppSecret                string     // 预算分派的密钥
	EncryptionKey            string     // 价格加密Key
	EncryptionSecret         string     // 价格加密密钥
	AppId                    string     // appid
	Status                   int8       // 状态 1=可用 2=不可用
	PkgName                  string     // 预算位app包名
	AppName                  string     // 预算位关联产品名称
	AppVersionCodes          string     // 预算位应用版本号,英文逗号分割
	AppVersionCodeArray      []string   // 预算位应用版本号,转化成数组
	AppStoreVersionCodes     string     // 预算位应用商店版本号,英文逗号分割
	AppStoreVersionCodeArray []string   // 预算位应用商店版本号
	AppStoreLinks            string     // 预算位应用商店地址，用逗号分割
	AppStoreLinksArray       []string   // 预算位应用商店地址
	Remark                   string     // 预算位配置页面备注
	DspCompany               DspCompany // 预算请求地址和方式
	rwLock                   sync.RWMutex

	// TODO 还有 deeplink 是否支持等规则 ...
}

type SspSlotInfoBindings struct {
	SlotBindingMaps sync.Map
	SspSlotInfos    []SspSlotInfo
}

// 用媒体广告位sspSlotId 和 dsp广告位 budgetSlotid绑定
// 当有新配置下发（Redis 订阅或配置文件变化）  redis: key = sspSlotId:AppId
func (this *SspSlotInfoBindings) UpdateBinding() {
	for _, biding := range this.SspSlotInfos {
		key := fmt.Sprintf("%d:%s", biding.SspSlotId, biding.AppId)
		this.SlotBindingMaps.Store(key, biding)
	}
}

// 动态跟新配置方案 订阅/发布者模式
func (this *SspSlotInfoBindings) onConfigUpdate(msg string) bool {
	return false
}

func (sspSlotInfo *SspSlotInfo) GetDspCode() string {
	return sspSlotInfo.DspCompany.DspCode
}
