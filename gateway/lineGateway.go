package gateway

import (
	"encoding/json"
	"strings"

	linepkg "github.com/KuriharaYuya/yuya-kanshi-serverless/repository/line"
	"github.com/KuriharaYuya/yuya-kanshi-serverless/usecase"
	utils "github.com/KuriharaYuya/yuya-kanshi-serverless/util"
)

func LineGateway(req utils.Request) {
	// get message
	msg, err := extractMsg(req)
	if err != nil {
		return
	}

	// handling usecase with message

	if strings.Contains(msg, "debug") {
		usecase.DebugNotionAPI(msg)
	}

}

func extractMsg(req utils.Request) (msg string, err error) {
	var webhookData linepkg.WebhookBody
	err = json.Unmarshal([]byte(req.Body), &webhookData)
	if err != nil {
		return "", err
	}

	if len(webhookData.Events) > 0 && webhookData.Events[0].Message.Type == "text" {
		textMessage := webhookData.Events[0].Message.Text
		return textMessage, nil
	}
	return "", nil
}
