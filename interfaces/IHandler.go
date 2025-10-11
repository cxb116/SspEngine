package interfaces

// 切面接口
// 接口里的IRequest 媒体请求json数据
type IHandler interface {
	PreHandle(request IRequest)
	Handle(request IRequest)
	AfterHandle(request IRequest)
}
