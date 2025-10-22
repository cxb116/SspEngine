package interfaces

type IDsp interface {
	//GetDspId(request IBidRequest) int32

	CreateReqMsg() string // 获取请求+ 管理端配置的信息

	CreateRes() // 构建预算返回的物料，转化成我方返回媒体的信息

	SendBidMsg() bool // 请求dsp服务器

	GetDspCost() (int64, int64) // 获取预算请求耗时，毫秒

	GetRes() IBidResponse // 将物料转化成我方文档数据，可通过getRes获取

	GetBidId() string // 获取bidRequestId
}
