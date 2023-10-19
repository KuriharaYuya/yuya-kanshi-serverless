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
	err := godotenv.Load()
	if err != nil {
		fmt.Printf("読み込み出来ませんでした: %v", err)

	}
	client = CreateClient()
}

func CreateClient() *notionapi.Client {
	integration_token := os.Getenv("NOTION_API_KEY")
	return notionapi.NewClient(notionapi.Token(integration_token))
}

func GetLog(date string) (interface{}, error) {
	db, err := setDB()
	if err != nil {
		return nil, err
	}

	// クエリを作成
	query := &notionapi.DatabaseQueryRequest{
		Filter: notionapi.PropertyFilter{
			Property: "allowPublish",
			Checkbox: &notionapi.CheckboxFilterCondition{
				Equals: true,
			},
		},
	}

	// クエリを使用してデータを取得
	results, err := client.Database.Query(context.Background(), notionapi.DatabaseID(db.ID), query)
	if err != nil {
		fmt.Println("エラー")
		return nil, err
	}
	tgt := results.Results[0].Properties
	fmt.Println(tgt, "りざると")
	return results, nil
}

// setDB関数はそのまま

func setDB() (db *notionapi.Database, err error) {
	db, err = client.Database.Get(context.Background(), notionapi.DatabaseID("8af74dfac9a0482bab353741bb355971"))
	if err != nil {
		return nil, err
	}
	return db, err
}
