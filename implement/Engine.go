package implement

import (
	"encoding/json"
	"fmt"
	"github.com/cxb116/sspEngine/internal/config"
	"github.com/cxb116/sspEngine/internal/readerbyte"
	"net/http"
	"time"
)

//var tempBufPool = sync.Pool{
//	New: func() interface{} {
//		return make([]readerbyte, 32*1024)
//	},
//}

type Engine struct {
	//RequestHandler   interfaces.IRequestHandler
	EngineHttpClient *http.Client
	//SspRequest       interfaces.IRequest // ssp请求对象池化
	ExitChan chan struct{} // 异步捕获链接关闭状态
	// 心跳检测

}

func newEngineHttpClient() *http.Client {
	return &http.Client{
		Timeout: time.Millisecond * 800,
	}
}

func newEngineWithConfig(config *config.Config) *Engine {
	return &Engine{
		//SspSlotInfo:      new(SspSlotInfo), //TODO 补充这个对象是多个
		//RequestHandler:   NewRequestHandler(),
		EngineHttpClient: newEngineHttpClient(),
		ExitChan:         nil,
	}

}

func ServerEngine(config *config.Config) {
	newEngineWithConfig(config)
	http.HandleFunc("/", SSP)
	http.ListenAndServe(":80", nil)
}

func SSP(w http.ResponseWriter, req *http.Request) {
	if req.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	defer req.Body.Close()

	reqBody, err := readerbyte.ReadBodyWithFixedBuf(req.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	var bidRequest BidRequest
	currentMilli := time.Now().UnixMilli() // 获取当前毫秒
	if err := json.Unmarshal(reqBody, &bidRequest); err != nil {
		http.Error(w, "invalid json: "+err.Error(), http.StatusBadRequest)
		return
	}

	// 将请求数据放入SspRequest对象池中
	request := GetBidRequest()
	request.RequestId = bidRequest.RequestId
	request.AppId = bidRequest.AppId
	request.RequestTime = currentMilli

	fmt.Printf("SSP Request Body: %s\n", request)
	// 将sspRequest 放入
	GSspRequestHandler.SendRequestToTaskQueue(request)
	//PutSspRequest(request)   同一个对象归还多次，会在池中出现多个相同的对象指针，下次get获取池对象会发生数据竞争

	// 5. 返回响应
	response := map[string]interface{}{
		"status": "success",
		"data":   bidRequest,
	}
	jsonData, _ := json.Marshal(response)
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}
