## 这里说一下这个东西怎么用

```go
// 比如我这里是业务代码
// 假设现在是一个handler 也就是接收请求返回响应
func HandleRequest(w http.ResponseWriter, r *http.Request){
  // 这里处理业务
  // 业务处理完了，我需要异步做一件事，比如更新某个员工的状态，但这种更新的结果我不关注
  // 当然我希望它能保证完成任务，但是不想等待这件事的执行，我现在就要把函数返回
  // 那么这里使用事件来做
  ...... 业务代码
  // 这里触发我想要做的那件事
  // TODO:触发创建角色权限的事件
	eventx.GlobalEmitter.Emit("role.init", map[string]any{
        // 这里是我想要做的那件事 需要使用到的参数
		"merchant_id": staff_register_resp.MerchantID,
		"staff_id":    staff_register_resp.StaffID,
	})
}

    // 好，现在函数就已经返回了，它不会去等待这件事的执行
```

-- 使用事件之前先要找个地方先把事件要做的事情先注册了

```go
// 这里是我的事件要做的是
func NotifySomeOneDoSomeThing(){
    eventx.GlobalEmitter.On("session.assign_staff", func(data any) {
		eventData := data.(map[string]any)
		staff_id := eventData["staff_id"].(int64)
		merchant_id := eventData["merchant_id"].(int64)

    // 上面是获取事件通知传递的参数
    // 下面是我要做的事		
	})
}

```

-- 最后，在main函数注册这件事的监听器

```go
func main(){
    // 注册这个事件的监听器
    event.NotifySomeOneDoSomeThing()
}
```