package implement

import (
	"errors"
	"fmt"
	"github.com/rs/zerolog/log"
	"github.com/yourusername/ssp_grpc/interfaces"
	"sync"
)

const (
	WORK_POOL_SIZE = 10 // 池大小
	MAX_CONN       = 10
)

var GSspRequestHandler *RequestHandler

func init() {
	GSspRequestHandler = NewRequestHandler()
	GSspRequestHandler.StartWorkerPool()
}

type RequestHandler struct {
	// TODO 对象SspSlotInfoMaps 初始化完成但是没有数据
	SspSlotInfoMaps  map[int32]*SspSlotInfo // key=SspSlotId value=*SspSlotInfo
	SspRequest       interfaces.IBidRequest // 请求数据
	WorkerPoolSize   int32                  // 业务工作Worker池的数量
	MaxWorkerTaskLen int32                  // 业务工作Worker对应负责的任务队列最大任务存储数量
	FreeWorkers      map[int32]struct{}     // 空闲worker集合
	FreeWorkerMutex  sync.Mutex
	TaskQueue        []chan interfaces.IBidRequest //Worker负责取任务的消息队列
	//ExtraFreeWorkers  map[int32]struct{}         // 空闲worker集合
	//ExtraFreeWorkerMu sync.Mutex
}

func NewRequestHandler() *RequestHandler {
	var freeWorkers map[int32]struct{}
	//var extraFreeWorkers map[int32]struct{}

	freeWorkers = make(map[int32]struct{}, WORK_POOL_SIZE)
	for i := int32(0); i < WORK_POOL_SIZE; i++ {
		freeWorkers[i] = struct{}{}
	}

	//extraFreeWorkers = make(map[int32]struct{}, MAX_CONN-WORK_POOL_SIZE)
	//for i := WORK_POOL_SIZE; i < MAX_CONN; i++ {
	//	extraFreeWorkers[int32(i)] = struct{}{}
	//}
	TaskQueueLen := MAX_CONN
	return &RequestHandler{
		SspSlotInfoMaps:  make(map[int32]*SspSlotInfo, 100),
		SspRequest:       NewBidRequest(),
		WorkerPoolSize:   WORK_POOL_SIZE,
		MaxWorkerTaskLen: MAX_CONN,
		FreeWorkers:      freeWorkers,
		//ExtraFreeWorkers: extraFreeWorkers,
		TaskQueue: make([]chan interfaces.IBidRequest, TaskQueueLen),
	}
}

func (rh *RequestHandler) StartWorkerPool() {
	for i := int32(0); i < rh.WorkerPoolSize; i++ {
		rh.TaskQueue[i] = make(chan interfaces.IBidRequest, rh.MaxWorkerTaskLen)
		go rh.StartOnWorker(i, rh.TaskQueue[i])
	}
}

// 将收到的消息推入TaskQueue
func (requestHandler *RequestHandler) SendRequestToTaskQueue(bidRequest interfaces.IBidRequest) {
	workerId := requestHandler.useWorker()
	log.Info().Msgf("SendRequestToTaskQueue workerId:%d\n", workerId)
	requestHandler.TaskQueue[workerId] <- bidRequest
	log.Info().Msgf("SendRequestToTaskQueue 将数据放到 TaskQueue中 workerId: %d\n", workerId)
	//PutSspRequest(sspRequest) //
}

func (requestHandler *RequestHandler) useWorker() int32 {
	var workerId int32
	requestHandler.FreeWorkerMutex.Lock()
	// 尝试从工作线程中取出人一个空闲的 workerId
	for workerId = range requestHandler.FreeWorkers {
		delete(requestHandler.FreeWorkers, workerId)
		fmt.Println("freeWorker: ", requestHandler.FreeWorkers)
		requestHandler.FreeWorkerMutex.Unlock()
		return workerId
	}
	requestHandler.FreeWorkerMutex.Unlock()

	return -1 // 没有空闲的workerId
}

func (requestHandler *RequestHandler) StartOnWorker(workerId int32, taskQueue chan interfaces.IBidRequest) {
	// 不断地等待队列中的消息
	for {
		select {
		// 有消息则取出队列的request,并执行绑定业务方法
		case request, ok := <-taskQueue:
			if !ok {
				//临时创建的worker, 是通过关闭taskQueue 来销毁当前worker
				return
			}
			switch req := request.(type) {
			case interfaces.IBidRequest: // Client message request
				requestHandler.doRequestDispatcher(req, workerId)
			}
		}
	}
}

func (requestHandler *RequestHandler) doRequestDispatcher(request interfaces.IBidRequest, workerId int32) (int, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Printf("workerId %s doRequestDispence panic err:%v", workerId, err)
		}
	}()

	log.Info().Msgf("doRequestDispence workerId:%d", workerId, "Req: ", request)
	// TODO 处理请求request
	validateStatus, err := ValidateRequest(request)
	if err != nil {
		fmt.Println("validateRequest err:", err)
		return -1, err
	}
	if validateStatus == -1 {
		fmt.Println("validateRequest err: invalid request")
		return -1, errors.New("invalid request")
	}
	// TODO 1 从 SspSlotInfoMaps 中获取管理端配置信息
	// TODO 2 和请求id数据匹配
	return validateStatus, err
}
