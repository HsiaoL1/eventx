package eventx

import "log/slog"

// PostEmit 触发 post.assign 事件
// 该函数用于发射一个 post.assign 事件，包含商户ID、会话ID、会话类型和客户ID等信息。
// 该事件可以被其他监听器捕获并处理，通常用于在系统中进行后续操作或通知。
// // 参数:
// - merchant_id: 商户ID
// - conversation_id: 会话ID
// - con_type: 会话类型（例如，1表示新访客会话，2表示老访客会话等）
// - customer_id: 客户ID
// // 使用示例:
// PostEmit(12345, 67890, 1, 54321)
// // 这将触发一个 post.assign 事件，包含商户ID 12345、会话ID 67890、会话类型 1 和客户ID 54321。
// 其他监听器可以通过 GlobalEmitter.On("post.assign", handler) 来注册处理
func PostEmit(merchant_id int64, conversation_id int64, con_type int8, customer_id int64) {
	GlobalEmitter.Emit("post.assign", map[string]any{
		"merchant_id":     merchant_id,
		"conversation_id": conversation_id,
		"con_type":        con_type,
		"customer_id":     customer_id,
	})
	slog.Info("post.emit", "merchant_id", merchant_id, "conversation_id", conversation_id, "con_type", con_type, "customer_id", customer_id)
}
