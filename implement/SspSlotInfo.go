package implement

import "sync"

var SspSlotInfoMaps = make(map[int32]*SspSlotInfo)

// 管理端配置的预算位置信息
type SspSlotInfo struct {
	SspSlotId                int32        `json:"ssp_slot_id,omitempty"`                  // 广告位id
	RtaExt                   string       `json:"rta_ext,omitempty"`                      //
	AdType                   int32        `json:"ad_type,omitempty"`                      // 广告属性
	OsType                   int32        `json:"os_type,omitempty"`                      // 系统属性 1=android 2=ios
	PayType                  int8         `json:"pay_type,omitempty"`                     // 结算方式 1=固价 2=分成 3=RTB
	DealRatio                int32        `json:"deal_ratio,omitempty"`                   // 成交价系数 单位%
	AppKey                   string       `json:"app_key,omitempty"`                      // 预算分派的key
	AppSecret                string       `json:"app_secret,omitempty"`                   // 预算分派的密钥
	EncryptionKey            string       `json:"encryption_key,omitempty"`               // 价格加密Key
	EncryptionSecret         string       `json:"encryption_secret,omitempty"`            // 价格加密密钥
	AppId                    string       `json:"app_id,omitempty"`                       // appid
	Status                   int8         `json:"status,omitempty"`                       // 状态 1=可用 2=不可用
	PkgName                  string       `json:"pkg_name,omitempty"`                     // 预算位app包名
	AppName                  string       `json:"app_name,omitempty"`                     // 预算位关联产品名称
	AppVersionCodes          string       `json:"app_version_codes,omitempty"`            // 预算位应用版本号,英文逗号分割
	AppVersionCodeArray      []string     `json:"app_version_code_array,omitempty"`       // 预算位应用版本号,转化成数组
	AppStoreVersionCodes     string       `json:"app_store_version_codes,omitempty"`      // 预算位应用商店版本号,英文逗号分割
	AppStoreVersionCodeArray []string     `json:"app_store_version_code_array,omitempty"` // 预算位应用商店版本号
	AppStoreLinks            string       `json:"app_store_links,omitempty"`              // 预算位应用商店地址，用逗号分割
	AppStoreLinksArray       []string     `json:"app_store_links_array,omitempty"`        // 预算位应用商店地址
	Remark                   string       `json:"remark,omitempty"`                       // 预算位配置页面备注
	DspCompany               DspCompany   `json:"dsp_company,omitempty"`                  // 预算请求地址和方式
	rwLock                   sync.RWMutex `json:"rw_lock,omitempty"`

	// TODO 还有 deeplink 是否支持等规则 ...

}

// 问题1 数据读出后如何和请求的数据匹配
// 问题2 在对接预算文档里面拿到的这个数据如何是后台管理端配置的那个
