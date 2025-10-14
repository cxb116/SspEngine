package implementation

import "time"

type Request struct {
	// 毫秒级时间戳  time.Now().UnixMilli()
	//媒体请求到SSP服务器的秒级别时间戳
	StartTime time.Time
}
