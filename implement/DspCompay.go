package implement

type DspCompany struct {
	DspId       string // 预算公司Id
	RequestUrl  string // 预算请求url
	RequestType int8   // 请求方式 1=post 2=get
	DspResTime  int64  // 请求预算响应时间,毫秒时间戳
	Status      int8   // 状态 1=可用 2=不可用
}

func (dspCompany *DspCompany) GetDspId() string {
	return dspCompany.DspId
}

func (dspCompany *DspCompany) GetRequestType() string {
	return dspCompany.RequestUrl
}

func (dspCompany *DspCompany) GetDspResTime() int64 {
	return dspCompany.DspResTime
}

func (dspCompany *DspCompany) GetDspStatus() int8 {
	return dspCompany.Status
}
