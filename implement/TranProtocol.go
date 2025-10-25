package implement

import (
	"bytes"
	"encoding/json"
	"github.com/cxb116/sspEngine/interfaces"
	"io"
	"net/http"
)

//TODO 逻辑有问题
func initC() {
	RegisterProtocol("json", func(protocol interfaces.ITranProtocol) interfaces.ITranProtocol {
		return &JsonProtocol{}
	})
}

var protocolRegistry = make(map[string]func(protocol interfaces.ITranProtocol) interfaces.ITranProtocol)

func RegisterProtocol(name string, protocol func(protocol interfaces.ITranProtocol) interfaces.ITranProtocol) {
	protocolRegistry[name] = protocol
}
func GetProtocol(name string, proto interfaces.ITranProtocol) interfaces.ITranProtocol {
	if protocol, ok := protocolRegistry[name]; ok {
		return protocol(proto)
	}
	return nil
}

type TranProtocol struct {
	Url          string
	Method       string
	ResponseTime int64 // 响应时间
}

type JsonProtocol struct {
	TranProtocol
}

func (proto *JsonProtocol) DoRequest(payload any) ([]byte, error) {
	body, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}
	req, err := http.NewRequest(proto.Method, proto.Url, bytes.NewReader(body))
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	return io.ReadAll(resp.Body)
}
