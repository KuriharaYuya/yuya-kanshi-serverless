package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/line/line-bot-sdk-go/v7/linebot"
)

func ReplyToUser() {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("読み込み出来ませんでした: %v", err)
	}

	secret := os.Getenv("LINE_CHANNEL_SECRET")
	channelToken := os.Getenv("LINE_CHANNEL_ACCESS_TOKEN")

	bot, err := linebot.New(secret, channelToken)
	if err != nil {
		fmt.Println(err)
		return

	}
	fmt.Println("Reply")
	bot.BroadcastMessage(linebot.NewTextMessage("Hello, world...")).Do()
}
