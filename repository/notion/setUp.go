package notionpkg

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/jomei/notionapi"
)

var client *notionapi.Client

func init() {
	if os.Getenv("ENVIRONMENT") == "development" {
		// ローカル開発環境用の処理
		err := godotenv.Load()
		if err != nil {
			fmt.Printf("読み込み出来ませんでした_setUp.go: %v", err)
		}
	} else {
		// 本番環境用の処理
		// 例: AWS Secrets Managerから秘密情報を取得する
	}
	client = CreateClient()
}

func CreateClient() *notionapi.Client {
	integration_token := os.Getenv("NOTION_API_KEY")
	return notionapi.NewClient(notionapi.Token(integration_token))
}

func switchLogDb() string {
	debugDbID := "b2c4752a33904be3a434f2c6542a4b75"
	prodDbID := "8af74dfac9a0482bab353741bb355971"
	debugMode := os.Getenv("DEBUG_MODE")
	if debugMode == "true" {
		return debugDbID
	}
	return prodDbID
}
func setLogDB() (db *notionapi.Database, err error) {
	dbId := switchLogDb()
	db, err = client.Database.Get(context.Background(), notionapi.DatabaseID(dbId))
	if err != nil {
		return nil, err
	}
	return db, err
}
