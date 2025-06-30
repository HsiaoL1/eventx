package eventx

import (
	"sync"
)

// GlobalEmitter 全局事件发射器实例
var GlobalEmitter *EventEmitter

func init() {
	GlobalEmitter = NewEventEmitter()
}

// EventHandler 定义事件处理函数的类型
type EventHandler func(data any)

// EventEmitter 事件发射器
type EventEmitter struct {
	listeners map[string][]EventHandler
	mu        sync.RWMutex
}

// NewEventEmitter 创建一个新的事件发射器
func NewEventEmitter() *EventEmitter {
	return &EventEmitter{
		listeners: make(map[string][]EventHandler),
	}
}

// On 注册事件监听器
func (e *EventEmitter) On(event string, handler EventHandler) {
	e.mu.Lock()
	defer e.mu.Unlock()

	if _, exists := e.listeners[event]; !exists {
		e.listeners[event] = make([]EventHandler, 0)
	}
	e.listeners[event] = append(e.listeners[event], handler)
}

// Off 移除事件监听器
func (e *EventEmitter) Off(event string, handler EventHandler) {
	e.mu.Lock()
	defer e.mu.Unlock()

	if handlers, exists := e.listeners[event]; exists {
		for i, h := range handlers {
			if &h == &handler {
				e.listeners[event] = append(handlers[:i], handlers[i+1:]...)
				break
			}
		}
	}
}

// Emit 触发事件
func (e *EventEmitter) Emit(event string, data any) {
	e.mu.RLock()
	defer e.mu.RUnlock()

	if handlers, exists := e.listeners[event]; exists {
		for _, handler := range handlers {
			go handler(data)
		}
	}
}

// Once 注册一次性事件监听器
func (e *EventEmitter) Once(event string, handler EventHandler) {
	var onceHandler EventHandler
	onceHandler = func(data any) {
		handler(data)
		e.Off(event, onceHandler)
	}
	e.On(event, onceHandler)
}

// RemoveAllListeners 移除所有监听器
func (e *EventEmitter) RemoveAllListeners(event string) {
	e.mu.Lock()
	defer e.mu.Unlock()

	if event == "" {
		e.listeners = make(map[string][]EventHandler)
	} else {
		delete(e.listeners, event)
	}
}

// ListenerCount 获取监听器数量
func (e *EventEmitter) ListenerCount(event string) int {
	e.mu.RLock()
	defer e.mu.RUnlock()

	if handlers, exists := e.listeners[event]; exists {
		return len(handlers)
	}
	return 0
}
