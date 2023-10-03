package notionpkg

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/jomei/notionapi"
)

type Database *notionapi.Database

func GetLog() (Database, error) {
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("読み込み出来ませんでした: %v", err)

	}
	integration_token := os.Getenv("NOTION_API_KEY")

	client := notionapi.NewClient(notionapi.Token(integration_token))
	db, err := client.Database.Get(context.Background(), notionapi.DatabaseID("8af74dfac9a0482bab353741bb355971"))
	if err != nil {
		return nil, err
	}
	return db, nil
}
