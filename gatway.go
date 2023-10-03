package main

import (
	"bytes"
	"strings"

	usecase "github.com/KuriharaYuya/yuya-kanshi-serverless/usecase"
)

var buf bytes.Buffer

const userAgent = "user-agent"
const lineBotWebhook = "LineBotWebhook"

func Gateway(req Request) *Response {
	ua := req.Headers[userAgent]

	// line-bot-request
	if strings.Contains(ua, lineBotWebhook) {
		usecase.CheckDailyLog()
	}

	// LineBotWebhookが含まれているか

	resp := Response{
		StatusCode:      200,
		IsBase64Encoded: false,
		Body:            buf.String(),
		Headers: map[string]string{
			"Content-Type":           "application/json",
			"X-MyCompany-Func-Reply": "hello-handler",
		},
	}
	return &resp
}
