package interfaces

type IBidRequest interface {
	GetRequestId() string
	SetRequestId(string)
	GetSspSlotId() int32
	SetSspSlotId(int32)
	GetAppId() string
	SetAppId(string)
	GetRequestTime() int64
	SetRequestTime(int64)
}
