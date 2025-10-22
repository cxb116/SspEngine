package implement

type DspCompany struct {
	DspId       string // 预算公司Id
	DspCode     string // 自定义预算位Code
	RequestUrl  string // 预算请求url
	RequestType int8   // 请求方式 1=post 2=get
	DspResTime  int64  // 请求预算响应时间,毫秒时间戳
	Status      int8   // 状态 1=可用 2=不可用
}

func (dspCompany *DspCompany) GetDspId() string {
	return dspCompany.DspId
}

func (dspCompany *DspCompany) SetDspId(dspId string) {
	dspCompany.DspId = dspId
}

func (dspCompany *DspCompany) getDspCode() string {
	return dspCompany.DspCode
}

func (dspCompany *DspCompany) SetDspCode(dspCode string) {
	dspCompany.DspCode = dspCode
}

func (dspCompany *DspCompany) GetRequestType() string {
	return dspCompany.RequestUrl
}

func (dspCompany *DspCompany) SetRequestType(requestType int8) {
	dspCompany.RequestType = requestType
}

func (dspCompany *DspCompany) GetDspResTime() int64 {
	return dspCompany.DspResTime
}

func (dspCompany *DspCompany) SetDspResTime(dspResTime int64) {
	dspCompany.DspResTime = dspResTime
}

func (dspCompany *DspCompany) GetDspStatus() int8 {
	return dspCompany.Status
}

func (dspCompany *DspCompany) SetDspStatus(dspStatus int8) {
	dspCompany.Status = dspStatus
}
