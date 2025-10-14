package interfaces

// 切面接口
// 接口里的IRequest 媒体请求json数据
type IHandler interface {
	// 控频策略,配合redis缓存完成
	PreHandle(request IRequest)

	Handle(request IRequest)

	AfterHandle(request IRequest)
}
