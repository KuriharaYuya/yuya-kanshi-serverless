package notionpkg

import (
	"context"
	"fmt"

	"github.com/jomei/notionapi"
)

// func DebugStructExchange
func GetDebugData(publish bool) (DebugLogProperty, error) {
	db, err := setLogDB()
	if err != nil {
		fmt.Println("setLogDBでエラー", err)
		return DebugLogProperty{}, err
	}
	query := setQuery(publish)

	// クエリを使用してデータを取得

	results, err := client.Database.Query(context.Background(), notionapi.DatabaseID(db.ID), query)
	if err != nil {
		fmt.Println("client.Database.Queryでエラー", err)
	}

	resultProps := results.Results[0].Properties
	nameProp, nameOk := resultProps["Name"].(*notionapi.TitleProperty)
	allowPublishProp, allowPublishOk := resultProps["allowPublish"].(*notionapi.CheckboxProperty)

	// name ok
	if !nameOk || !allowPublishOk {
		fmt.Println("Error extracting properties")
		fmt.Println(nameProp, "resultProps")
		return DebugLogProperty{}, err
	}
	debugLog := DebugLogProperty{
		Name:         nameProp,
		AllowPublish: allowPublishProp,
	}
	return debugLog, nil
}

func setQuery(publish bool) *notionapi.DatabaseQueryRequest {
	checkboxCondition := &notionapi.CheckboxFilterCondition{}
	if publish {
		checkboxCondition.Equals = true
	} else {
		checkboxCondition.DoesNotEqual = true
	}

	query := &notionapi.DatabaseQueryRequest{
		Filter: notionapi.PropertyFilter{
			Property: "allowPublish",
			Checkbox: checkboxCondition,
		},
	}
	return query
}
