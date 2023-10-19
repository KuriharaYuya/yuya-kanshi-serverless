package linepkg

import (
	"encoding/json"

	utils "github.com/KuriharaYuya/yuya-kanshi-serverless/util"
)

func LineGateway(req utils.Request) {
	msg, err := extractMsg(req)
	if err != nil {
		return
	}
	ReplyToUser(msg)
}

func extractMsg(req utils.Request) (msg string, err error) {
	var webhookData WebhookBody
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
