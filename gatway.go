package main

import (
	"bytes"
	"strings"
	"sync"

	linepkg "github.com/KuriharaYuya/yuya-kanshi-serverless/repository/line"
)

var buf bytes.Buffer

const userAgent = "user-agent"
const lineBotWebhook = "LineBotWebhook"

func Gateway(req Request) *Response {
	wg := sync.WaitGroup{}
	wg.Add(1)
	ua := req.Headers[userAgent]

	go func() {
		defer wg.Done()
		// line-bot-request
		if strings.Contains(ua, lineBotWebhook) {
			// usecase.CheckDailyLog()
			linepkg.ReplyToUser("run in gateway")
		}
	}()

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
