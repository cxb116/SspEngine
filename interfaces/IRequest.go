package interfaces

import "fmt"

type IRequest interface {
	GetRequest() string
}
type BaseRequest struct {
}

func (this *BaseRequest) GetRequest() string {
	return fmt.Sprintf("BaseRequest:%s", this)
}

type Request struct {
}

func (this *Request) GetRequest() string {
	return fmt.Sprintf("Request:%s", this)
}

func HandleRequest(request IRequest) {
	fmt.Println(request.GetRequest())
}

//func main() {
//	baseRequest := &BaseRequest{}
//	request := &Request{}
//	HandleRequest(baseRequest)
//	HandleRequest(request)
//}
