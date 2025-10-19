package interfaces

// 请求处理器
type IRequestHandler interface {
	AddRequest(request IRequest)
}
