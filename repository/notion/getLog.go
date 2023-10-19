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
	debug(results)
	return results, nil
}

func switchTgtDb() string {
	debugDbID := "b2c4752a33904be3a434f2c6542a4b75"
	prodDbID := "8af74dfac9a0482bab353741bb355971"
	debugMode := os.Getenv("DEBUG_MODE")
	if debugMode == "true" {
		return debugDbID
	}
	return prodDbID
}
func setDB() (db *notionapi.Database, err error) {
	dbId := switchTgtDb()
	db, err = client.Database.Get(context.Background(), notionapi.DatabaseID(dbId))
	if err != nil {
		return nil, err
	}
	return db, err
}

// debug code
func debug(results *notionapi.DatabaseQueryResponse) {
	resultProps := results.Results[0].Properties
	nameProp, nameOk := resultProps["Name"].(*notionapi.TitleProperty)
	allowPublishProp, allowPublishOk := resultProps["allowPublish"].(*notionapi.CheckboxProperty)

	// name ok
	if !nameOk || !allowPublishOk {
		fmt.Println("Error extracting properties")
		fmt.Println(nameProp, "resultProps")
		return
	}

	debugLog := DebugLogProperty{
		Name:         nameProp,
		allowPublish: allowPublishProp,
	}

	fmt.Println(debugLog, "debugLog")
}
