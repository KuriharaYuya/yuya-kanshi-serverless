package main

import (
	"bytes"
	"strings"
	"sync"
)

var buf bytes.Buffer

const userAgent = "user-agent"
const lineBotWebhook = "LineBotWebhook"

func Gateway(req Request) *Response {
	var wg sync.WaitGroup

	wg.Add(1)
	// line-bot-request
	// "user-agent":"LineBotWebhook/2.0"
	ua := req.Headers[userAgent]

	// LineBotWebhookが含まれているか

	if strings.Contains(ua, lineBotWebhook) {
		// LineBotWebhookの場合
		go func() {
			ReplyToUser("確認しました")
			wg.Done()
		}()
		wg.Wait()
		return nil
	}

	resp := Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            buf.String(),
		Headers: map[string]string{
			"Content-Type":           "application/json",
			"X-MyCompany-Func-Reply": "hello-handler",
		},
	}
	wg.Wait()
	return &resp
}
