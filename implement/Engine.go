package implement

import (
	"encoding/json"
	"fmt"
	"github.com/cxb116/sspEngine/internal/config"
	"github.com/cxb116/sspEngine/internal/readerbyte"
	"net/http"
	"time"
)

type Engine struct {
	//RequestHandler   interfaces.IRequestHandler
	EngineHttpClient *http.Client
	//SspRequest       interfaces.IRequest // ssp请求对象池化
	ExitChan chan struct{} // 异步捕获链接关闭状态
	// 心跳检测

}

func newEngineHttpClient() *http.Client {
	return &http.Client{
		Timeout: time.Millisecond * 850, // 超时时间850毫秒
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
	if err := json.Unmarshal(reqBody, &bidRequest); err != nil {
		http.Error(w, "invalid json: "+err.Error(), http.StatusBadRequest)
		return
	}

	// 将请求数据放入SspRequest对象池中
	// TODO 这里过于需要改一下
	requestPool := GetBidRequest()
	requestPool = &bidRequest
	requestPool.RequestTime = time.Now().UnixMilli() // 获取当前毫秒

	fmt.Printf("SSP Request Body: %s\n", requestPool)
	// 将sspRequest 放入
	GSspRequestHandler.SendRequestToTaskQueue(requestPool)
	//PutSspRequest(request)   同一个对象归还多次，会在池中出现多个相同的对象指针，下次get获取池对象会发生数据竞争

	// dsp 返回物料，通过公共管道来将物料返回给这个请求
	//select {
	//case <-time.After(850 * time.Millisecond):
	//	http.Error(w, "timeout", http.StatusServiceUnavailable)
	//}

	// 5. 返回响应
	response := map[string]interface{}{
		"status": "success",
		"data":   bidRequest,
	}
	jsonData, _ := json.Marshal(response)
	w.WriteHeader(http.StatusOK)
	w.Write(jsonData)
}
