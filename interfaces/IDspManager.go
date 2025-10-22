package interfaces

type IDspManager interface {
	GetDspId() int32
	Match(bidRequest IBidRequest) bool
	Bid(bidRequest IBidRequest) (IBidResponse, error)
}
