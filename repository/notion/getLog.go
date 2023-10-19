package notionpkg

import (
	"context"
	"fmt"

	"github.com/jomei/notionapi"
)

func GetLog(date string) (interface{}, error) {
	db, err := setLogDB()
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
	return results, nil
}
