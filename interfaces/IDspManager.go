package interfaces

type IDSP interface {
	GetDspId() int64
	Match(bidRequest IBidRequest) bool
	Bid(bidRequest IBidRequest) (IBidRequest, error)
}
