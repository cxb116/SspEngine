package implement

import (
	"github.com/yourusername/ssp_grpc/interfaces"
	"sync"
)

type RequestHandler struct {
	SspSlotInfoMap    map[int32]*SspSlotInfo // key=SspSlotId value=*SspSlotInfo
	Request           interfaces.IRequest    // 请求数据
	WorkerPoolSize    int32                  // 业务工作Worker池的数量
	FreeWorkers       map[int32]struct{}     // 空闲worker集合
	FreeWorkerMutex   sync.Mutex
	TaskQueue         []chan interfaces.IRequest //Worker负责取任务的消息队列
	extraFreeWorkers  map[uint32]struct{}        //// 空闲worker集合
	extraFreeWorkerMu sync.Mutex
}
