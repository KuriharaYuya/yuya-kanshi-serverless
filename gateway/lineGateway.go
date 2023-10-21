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
		return
	}

	if strings.Contains(msg, "コンソール") {
		consoleType := detectConsole(msg)
		date := ExtractDate(msg)

		if consoleType == status {
			usecase.CheckDailyLog(date)
		}
		return
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

func ExtractDate(msg string) (date string) {
	// ex msg) コンソール: ステータス2023/08/13

	// まず日付を抽出する
	// ステータスまたは投稿という文字列が含まれているか確認し、含まれていればそこからテキストの終わりまでを抽出する
	dateString := strings.Split(msg, "ステータス")[1]

	//  スラッシュをハイフンに変換する
	return strings.Replace(dateString, "/", "-", -1)
}

const (
	other  = 0
	status = 1
	post   = 2
)

func detectConsole(msg string) (date int) {

	// 含まれているかどうかを確認する
	if strings.Contains(msg, "ステータス") {
		return status
	}
	if strings.Contains(msg, "投稿") {
		return post
	}
	return 0
}
