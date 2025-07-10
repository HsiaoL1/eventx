package eventx

import (
	"bytes"
	"context"
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
	"time"
)

func PostAssignEvent() {
	GlobalEmitter.On("post.assign", func(data any) {
		// 安全的类型断言
		eventData, ok := data.(map[string]any)
		if !ok {
			slog.Error("Invalid event data type", "data", data)
			return
		}

		// 安全的字段提取
		merchantID, ok := eventData["merchant_id"].(int64)
		if !ok {
			slog.Error("Invalid merchant_id type", "merchant_id", eventData["merchant_id"])
			return
		}

		conversationID, ok := eventData["conversation_id"].(int64)
		if !ok {
			slog.Error("Invalid conversation_id type", "conversation_id", eventData["conversation_id"])
			return
		}

		conType, ok := eventData["con_type"].(int8)
		if !ok {
			slog.Error("Invalid con_type type", "con_type", eventData["con_type"])
			return
		}

		customerID, ok := eventData["customer_id"].(int64)
		if !ok {
			slog.Error("Invalid customer_id type", "customer_id", eventData["customer_id"])
			return
		}

		// 构建POST请求数据
		postData := map[string]any{
			"merchant_id":     merchantID,
			"conversation_id": conversationID,
			"con_type":        conType,
			"customer_id":     customerID,
		}

		// 序列化请求体
		reqBody, err := json.Marshal(postData)
		if err != nil {
			slog.Error("Error marshalling request body", "error", err, "data", postData)
			return
		}

		// 创建带超时的HTTP客户端
		ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
		defer cancel()

		httpClient := &http.Client{
			Timeout: 30 * time.Second,
		}

		// 创建请求
		req, err := http.NewRequestWithContext(ctx, "POST", "http://127.0.0.1:8090/api/post/assign", bytes.NewBuffer(reqBody))
		if err != nil {
			slog.Error("Error creating request", "error", err)
			return
		}

		// 设置请求头
		req.Header.Set("Content-Type", "application/json")

		// 发送请求
		resp, err := httpClient.Do(req)
		if err != nil {
			slog.Error("Error sending request", "error", err)
			return
		}
		defer resp.Body.Close()

		// 检查响应状态码
		if resp.StatusCode != http.StatusOK {
			// 读取错误响应体
			errorBody, _ := io.ReadAll(resp.Body)
			slog.Error("Error response from server",
				"status", resp.Status,
				"status_code", resp.StatusCode,
				"error_body", string(errorBody))
			return
		}

		// 处理成功响应
		slog.Info("发送请求成功",
			"merchant_id", merchantID,
			"conversation_id", conversationID,
			"con_type", conType,
			"customer_id", customerID,
			"status_code", resp.StatusCode)
	})
}
