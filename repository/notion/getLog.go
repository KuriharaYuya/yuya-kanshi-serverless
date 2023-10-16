package notionpkg

import (
	"context"
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/jomei/notionapi"
)

type Database *notionapi.Database

func GetLog(date string) (Log, error) {
	db, err := GetDB()
	// db.Query(context.Background(), notionapi.DatabaseQuery{
	// 	Filter: notionapi.DatabaseQueryFilter{
	// 		Property: "FilledAtr",
	// 		Checkbox: &notionapi.DatabaseQueryCheckboxFilter{
	// 			Equals: true,
	// 		},
	// 	},
	// })

	// ここで取得してくる

	// ここでLogPropertyに型が合致しているかを確認する。

	// LogPropertyをLogに変換する なんかいい感じにやってくれる関数を作成する

	// ここでLogに型が合致しているかを確認する。

}

func GetDB() (Database, err error) {
	err = godotenv.Load()
	if err != nil {
		fmt.Printf("読み込み出来ませんでした: %v", err)

	}
	integration_token := os.Getenv("NOTION_API_KEY")

	client := notionapi.NewClient(notionapi.Token(integration_token))
	db, err = client.Database.Get(context.Background(), notionapi.DatabaseID("8af74dfac9a0482bab353741bb355971"))
	if err != nil {
		return nil, err
	}
	return db, nil
}
