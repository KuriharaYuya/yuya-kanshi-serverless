package notionpkg

import (
	"context"
	"fmt"
	"os"

	"github.com/jomei/notionapi"
)

// TODO:notionapiのインポート整備
// TODO:destructureForCalenderDataを使えるようにする

func GetLogListFromNow(onlyPublish bool) (LogListOutPut, error) {
	conditions := switchSearchCondition(onlyPublish)
	// filterとsortsを適用してクエリを行う
	params := &notionapi.DatabaseQueryRequest{
		Filter: conditions,
		Sorts: []notionapi.SortObject{
			{
				Property:  "date",
				Direction: "descending",
			},
		},
	}
	integration_token := os.Getenv("NOTION_API_KEY")
	client := notionapi.NewClient(notionapi.Token(integration_token))
	data, err := client.Database.Query(context.Background(), notionapi.DatabaseID("8af74dfac9a0482bab353741bb355971"), params)
	if err != nil {
		fmt.Println("Error:", err)
		return LogListOutPut{}, err
	}
	// calenderData := destructureForCalenderData(data)
	tableData := destructureForTable(data)
	return LogListOutPut{
		// CalenderData: calenderData, 
		TableData: tableData}, nil
}

type CalenderData map[string]int

type LogListOutPut struct {
	// CalenderData CalenderData
	TableData    []LogListType
}

// func destructureForCalenderData(data LogListPropertyForGitLikeCalender) CalenderData {
// 	values := make(map[string]int)
// 	for _, log := range data.Properties.Date.Date {
// 		values[log.d ] = 4
// 	}
// 	return values
// }

func destructureForTable(data []LogListPropertyForGitLikeCalender) []LogListType {
	var tgt []LogListType
	for _, log := range data {
		logItem := LogListType{
			UUID:      log.Properties.UUID.Formula.String,
			Date:      log.Properties.Date.Date,
			Title:     log.Properties.Title.Title,
			TweetUrl:  log.Properties.TweetUrl.Url,
			Published: log.Properties.Published.Fomula.Bool,
		}
		tgt = append(tgt, logItem)
	}
	return tgt
}

func switchSearchCondition(onlyPublish bool) interface{} {
	var conditions interface{}
	if onlyPublish {
		conditions = map[string]interface{}{
			"and": []interface{}{
				map[string]interface{}{
					"property": "published",
					"formula": map[string]interface{}{
						"checkbox": map[string]interface{}{
							"equals": true,
						},
					},
				},
				map[string]interface{}{
					"property": "date",
					"date": map[string]interface{}{
						"past_year": map[string]interface{}{},
					},
				},
			},
		}
	} else {
		conditions = map[string]interface{}{
			"property": "date",
			"date": map[string]interface{}{
				"past_year": map[string]interface{}{},
			},
		}
	}
	return conditions
}

func countPublishedLogs(onlyPubLish bool) (int, error) {
	conditions := switchSearchCondition(onlyPubLish)
	params := &notionapi.DatabaseQueryRequest{
		Filter:  conditions,
		Sorts: []notionapi.SortObject{
			{
				Property:  "date",
				Direction: "descending",
			},
		},
	}
	integration_token := os.Getenv("NOTION_API_KEY")
	client := notionapi.NewClient(notionapi.Token(integration_token))

	data, err := client.Database.Query(context.Background(), notionapi.DatabaseID("8af74dfac9a0482bab353741bb355971"), params)

	if err != nil {
		return 0, err
	}

	
	count := len(data)
	return count, nil
}

type Log struct {
	Properties struct {
		Date struct {
			Date struct {
				Start string `json:"start"`
			} `json:"date"`
		} `json:"date"`
	} `json:"properties"`
}

type LogListType struct {
	UUID      string
	Date      string
	Title     string
	TweetUrl  string
	Published bool
}
