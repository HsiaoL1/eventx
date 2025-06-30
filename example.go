package eventx

import (
	"fmt"
	"time"
)

// 示例：如何使用事件系统
func Example() {
	// 创建事件发射器
	emitter := NewEventEmitter()

	// 定义事件处理函数
	handler1 := func(data interface{}) {
		fmt.Printf("Handler 1 received: %v\n", data)
	}

	handler2 := func(data interface{}) {
		fmt.Printf("Handler 2 received: %v\n", data)
	}

	// 注册事件监听器
	emitter.On("user.created", handler1)
	emitter.On("user.created", handler2)

	// 注册一次性事件监听器
	emitter.Once("user.deleted", func(data interface{}) {
		fmt.Printf("One-time handler received: %v\n", data)
	})

	// 触发事件
	emitter.Emit("user.created", map[string]interface{}{
		"id":   1,
		"name": "John Doe",
	})

	// 触发一次性事件
	emitter.Emit("user.deleted", map[string]interface{}{
		"id": 1,
	})

	// 再次触发一次性事件（不会触发处理函数）
	emitter.Emit("user.deleted", map[string]interface{}{
		"id": 2,
	})

	// 移除特定事件的所有监听器
	emitter.RemoveAllListeners("user.created")

	// 获取特定事件的监听器数量
	count := emitter.ListenerCount("user.created")
	fmt.Printf("Number of listeners for user.created: %d\n", count)

	// 等待所有goroutine完成
	time.Sleep(time.Second)
}

// 实际使用示例
func ExampleUsage() {
	// 注册事件监听器
	GlobalEmitter.On("order.created", func(data interface{}) {
		order := data.(map[string]interface{})
		fmt.Printf("New order created: %v\n", order)
		// 处理订单创建后的逻辑
	})

	GlobalEmitter.On("order.paid", func(data interface{}) {
		order := data.(map[string]interface{})
		fmt.Printf("Order paid: %v\n", order)
		// 处理订单支付后的逻辑
	})

	// 在业务代码中触发事件
	GlobalEmitter.Emit("order.created", map[string]interface{}{
		"id":     1,
		"amount": 100.00,
		"status": "created",
	})
}
