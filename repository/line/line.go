package linepkg

import (
	"fmt"
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func ReplyToUser(cnt string) {
	if os.Getenv("ENVIRONMENT") == "development" {
		// ローカル開発環境用の処理
		err := godotenv.Load()
		if err != nil {
			fmt.Printf("読み込み出来ませんでした_line.go: %v", err)
		}
	} else {
		// 本番環境用の処理
		// 例: AWS Secrets Managerから秘密情報を取得する
	}

	secret := os.Getenv("LINE_CHANNEL_SECRET")
	channelToken := os.Getenv("LINE_CHANNEL_ACCESS_TOKEN")

	bot, err := linebot.New(secret, channelToken)
	if err != nil {
		fmt.Println(err)
		return

	}
	log.Println("Reply")
	bot.BroadcastMessage(linebot.NewTextMessage(cnt)).Do()
	return
}
