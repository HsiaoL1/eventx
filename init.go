package eventx

// InitPostEvent 初始化 post 事件
// 该函数用于初始化 post 事件，包括注册 post.assign 事件的监听器。
// 当 post.assign 事件被触发时，会调用 PostAssignEvent 函数来处理该事件。
// 该函数通常在应用程序启动时被调用，以确保 post 事件能够正常处理。
// 使用示例:
// InitPostEvent()
// 这将初始化 post 事件，并注册 post.assign 事件的监听器。
// 其他监听器可以通过 GlobalEmitter.On("post.assign", handler) 来注册处理
// 注册在main.go中 main 函数中
func InitPostEvent() {
	PostAssignEvent()
}

func init() {
	InitPostEvent()
}
