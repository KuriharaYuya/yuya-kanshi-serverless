package linepkg

import (
	"encoding/json"

	utils "github.com/KuriharaYuya/yuya-kanshi-serverless/util"
)

func LineGateway(req utils.Request) {

	var webhookData WebhookBody
	err := json.Unmarshal([]byte(req.Body), &webhookData)
	if err != nil {
		// handle the error
	}

	if len(webhookData.Events) > 0 && webhookData.Events[0].Message.Type == "text" {
		textMessage := webhookData.Events[0].Message.Text
		ReplyToUser(textMessage)
	}

}
