package notionpkg

import (
	"context"
	"fmt"

	"github.com/jomei/notionapi"
)

func GetLog(date string) (LifeLog, error) {
	db, err := setLogDB()
	if err != nil {
		return LifeLog{}, err
	}

	parsedTime := parseTime(date)

	// time.Time 型を notionapiのDate 型へキャスト
	targetDate := notionapi.Date(parsedTime)

	// クエリを作成
	query := createDateQuery(targetDate)

	// クエリを使用してデータを取得
	results, err := client.Database.Query(context.Background(), notionapi.DatabaseID(db.ID), query)
	if err != nil {
		fmt.Println("repository/notion/getLog.goのdatabase.Queryでエラーが発生しました", err)
		return LifeLog{}, err
	}
	return serializeToLogProp(results)
}
