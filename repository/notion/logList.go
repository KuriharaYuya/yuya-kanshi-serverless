package notionpkg

// import (
// 	"context"
// 	"fmt"
// 	"os"

// 	"github.com/jomei/notionapi"
// )

// func GetLogListFromNow(onlyPublish bool) (LogListOutPut, error) {
// 	conditions := switchSearchCondition(onlyPublish)
// 	// filterとsortsを適用してクエリを行う
// 	params := &notionapi.DatabaseQueryRequest{
// 		Filter: &notionapi.DatabaseQueryFilter(conditions),
// 	}
// 	integration_token := os.Getenv("NOTION_API_KEY")
// 	client := notionapi.NewClient(notionapi.Token(integration_token))
// 	data, err := client.Database.Query(context.Background(), notionapi.DatabaseID("8af74dfac9a0482bab353741bb355971"), params)
// 	if err != nil {
// 		fmt.Println("Error:", err)
// 		return LogListOutPut{}, err
// 	}

// 	// TODO dataの型にあっていないので、型を変換する
// 	// TODO:destructureForCalenderDataを使えるようにする
// 	// calenderData := destructureForCalenderData(data)
// 	// TODO dataの型にあっていないので、型を変換する
// 	tableData := destructureForTable(data)
// 	return LogListOutPut{
// 		// CalenderData: calenderData,
// 		TableData: tableData}, nil
// }

// type CalenderData map[string]int

// type LogListOutPut struct {
// 	// CalenderData CalenderData
// 	TableData []LogListType
// }

// // TODO:destructureForCalenderDataを使えるようにする
// //
// //	func destructureForCalenderData(data LogListPropertyForGitLikeCalender) CalenderData {
// //		values := make(map[string]int)
// //		for _, log := range data.Properties.Date.Date {
// //			values[log.d ] = 4
// //		}
// //		return values
// //	}
// func destructureForTable(data []LogListPropertyForGitLikeCalender) []LogListType {
// 	var tgt []LogListType
// 	for _, log := range data {
// 		logItem := LogListType{
// 			UUID:      log.Properties.UUID.Formula.String,
// 			Date:      log.Properties.Date.Date, // 日付にアクセスする際はStartを使用
// 			Title:     log.Properties.Title.Title,
// 			TweetUrl:  log.Properties.TweetUrl.Url,           // URLにアクセスする際は大文字のURLを使用
// 			Published: log.Properties.Published.Formula.Bool, // Formulaのスペルが間違っていました
// 		}
// 		tgt = append(tgt, logItem)
// 	}
// 	return tgt
// }

// func switchSearchCondition(onlyPublish bool) interface{} {
// 	var conditions interface{}
// 	if onlyPublish {
// 		conditions = map[string]interface{}{
// 			"and": []interface{}{
// 				map[string]interface{}{
// 					"property": "published",
// 					"formula": map[string]interface{}{
// 						"checkbox": map[string]interface{}{
// 							"equals": true,
// 						},
// 					},
// 				},
// 				map[string]interface{}{
// 					"property": "date",
// 					"date": map[string]interface{}{
// 						"past_year": map[string]interface{}{},
// 					},
// 				},
// 			},
// 		}
// 	} else {
// 		conditions = map[string]interface{}{
// 			"property": "date",
// 			"date": map[string]interface{}{
// 				"past_year": map[string]interface{}{},
// 			},
// 		}
// 	}
// 	return conditions
// }

// func countPublishedLogs(onlyPubLish bool) (int, error) {
// 	conditions := switchSearchCondition(onlyPubLish)
// 	params := &notionapi.DatabaseQueryRequest{
// 		Filter: notionpkg.DatabaseQueryFilter(conditions),
// 		Sorts: []notionapi.SortObject{
// 			{
// 				Property:  "date",
// 				Direction: "descending",
// 			},
// 		},
// 	}
// 	integration_token := os.Getenv("NOTION_API_KEY")
// 	client := notionapi.NewClient(notionapi.Token(integration_token))

// 	data, err := client.Database.Query(context.Background(), notionapi.DatabaseID("8af74dfac9a0482bab353741bb355971"), params)

// 	if err != nil {
// 		return 0, err
// 	}

// 	// TODO dataからPublishedLogsのcountを算出

// 	return count, nil
// }

// type Log struct {
// 	Properties struct {
// 		Date struct {
// 			Date struct {
// 				Start string `json:"start"`
// 			} `json:"date"`
// 		} `json:"date"`
// 	} `json:"properties"`
// }

// type LogListType struct {
// 	UUID      string
// 	Date      string
// 	Title     string
// 	TweetUrl  string
// 	Published bool
// }
