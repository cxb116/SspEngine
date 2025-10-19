package implement

import (
	"encoding/json"
	"github.com/rs/zerolog/log"
	"github.com/yourusername/ssp_grpc/interfaces"
	"github.com/yourusername/ssp_grpc/internal/config"
	"net/http"
	"time"
)

type Engine struct {
	EngineHttpClient *http.Client
	sspRequest       interfaces.IRequest // ssp请求对象池化
	ExitChan         chan struct{}       // 异步捕获链接关闭状态
	// 心跳检测

}

func newEngineHttpClient() *http.Client {
	return &http.Client{
		Timeout: time.Millisecond * 800,
	}
}

func newEngineWithConfig(config *config.Config) *Engine {
	return &Engine{
		SspSlotInfo:      new(SspSlotInfo), //TODO 补充这个对象是多个
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
		log.Info().Msgf("SSP method not allowed")
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	response := map[string]interface{}{
		"success": true,
		"data":    "Engine instance",
	}
	encoder := json.NewEncoder(w)
	encoder.Encode(response)
	w.WriteHeader(http.StatusOK)
}
