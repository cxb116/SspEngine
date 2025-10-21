package implement

import "sync"

var SspSlotInfoMaps = make(map[int32]*SspSlotInfo)

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
}

// 问题1 数据读出后如何和请求的数据匹配
// 问题2 在对接预算文档里面拿到的这个数据如何是后台管理端配置的那个
