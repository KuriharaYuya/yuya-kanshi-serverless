package notionpkg

// import (
// 	"context"
// 	"errors"
// 	"fmt"

// 	// "fmt"
// 	"os"

// 	"github.com/jomei/notionapi"
// )

// func getLogDetail(tgtDate string) (LogOutPut, error) {
// // 	// "佑弥管理DB"から指定された日付に一致するページを取得する
// 	integration_token := os.Getenv("NOTION_API_KEY")
// 	client := notionapi.NewClient(notionapi.Token(integration_token))
// 	databaseID := "8af74dfac9a0482bab353741bb355971"

// 	filter := notionapi.AndCompoundFilter{
// 		notionapi.PropertyFilter{
// 			Property: "date",
// 			Date: &notionapi.DateFilterCondition{
// 				Equals:  &notionapi.Date{Date:  tgtDate},
// 			},
// 		},

// 	}
// 	results, err := client.Databases.Query(context.Background(), notionapi.DatabaseID(databaseID), filter)
// 	if err != nil {
// 		fmt.Println(err)
// 		return err
// 	}

// 	// 結果を加工して返す
// 	logOutput, err := createLogOutput(results)
// 	if err != nil {
// 		fmt.Println(err)
// 		return err
// 	}

// 	// まず、ここで投稿が不完全であることを検証する。trueであれば、500エラーを返す

// 	if len(results) == 0 || (len(results) > 1 && !results[0].Properties["published"].(bool)) {
// 		// 500エラーを返す
// 		return "存在しない", errors.New("500 Error")
// 	}

// // logIdとlogPropertyの取得
// var logId string
// if len(results) > 0 {
// 	logId = results[0].ID.String()
// }
// logProperty := results[0].Properties

// 	// fillAtrがfalseの場合は、処理を中断する
// 	if logProperty["filledAtr"].(bool) == false {
// 			return "準備未完了"
// 	}

// 	morningActivity, err := getMorningActivity(logId)
// 	if err != nil {
// 		return err
// 	}

// 	monthlyRecord, err := getMonthlyRecord(logId)
// 	if err != nil {
// 		return err
// 	}

// 	HostsImageRecord, err := getHostImageRecord(logId)
// 	if err != nil {
// 		return err
// 	}

// 	// フロント側で使いたい形に整形し返却
// 	return createLogOutput(
// 			logProperty,
// 			morningActivity,
// 			monthlyRecord,
// 			HostsImageRecord,
// 	)
// }

// func getMorningActivity(logID string) (MorningActivity, error) {
// 	databaseID := "472561d83e9247e6b018b271e5e95dad"
// 	filter := notionapi.DatabaseQueryParams{
// 		Filter: notionapi.PropertyFilter{
// 			Property: "佑弥管理",
// 			Relation: &notionapi.RelationCondition{
// 				Contains: logID,
// 			},
// 		},
// 	}
// 	results, err := client.Databases.Query(context.Background(), filter)
// 	if err != nil {
// 		return MorningActivity{}, err
// 	}

// 	if len(results) == 0 {
// 		return MorningActivity{}, errors.New("no results found")
// 	}

// 	var morningProps MorningActivity

// 	return morningProps
// }

// func getMonthlyRecord(logID string) (MonthlyRecord, error) {
// 	databaseID := "cf0777ea654d49d7b09b9f9a9cab747f"
// 	filter := notionapi.DatabaseQueryParams{
// 		Filter: notionapi.PropertyFilter{
// 			Property: "佑弥管理",
// 			Relation: &notionapi.RelationCondition{
// 				Contains: logID,
// 			},
// 		},
// 	}
// 	results, err := client.Databases.Query(context.Background(), filter)
// 	if err != nil {
// 		return err
// 	}

// 	if len(results) == 0 {
// 		return MonthlyRecord{}
// 	}

// 	var monthlyProps MonthlyRecord

// 	return monthlyProps
// }

// func getHostImageRecord(logID string) (HostsImageRecord, error) {
// 	databaseID := "4e0e1441d11b41b2a526dc5c0f0f27fc"
// 	filter := notionapi.DatabaseQueryParams{
// 		Filter: notionapi.PropertyFilter{
// 			Property: "佑弥管理",
// 			Relation: &notionapi.RelationCondition{
// 				Contains: logID,
// 			},
// 		},
// 	}
// 	results, err := client.Databases.Query(context.Background(), filter)
// 	if err != nil {
// 		return err
// 	}

// 	if len(results) == 0 {
// 		return HostsImageRecord{}, errors.New("no results found")
// 	}

// 	var hostProps HostsImageRecord

// 	return hostProps
// }

// func createLogOutput(logProps LogProperty, morningActivity MorningActivity, monthlyProps MonthlyRecord, hostsImageRecord HostsImageRecord) LogOutPut {
// 	mornings := digMorningData(logProps, morningActivity)
// 	device := digDeviceData(logProps, monthlyProps)
// 	diet := digDietData(logProps, monthlyProps)
// 	diary := digDiaryData(logProps)
// 	hostsImage := digHostImageData(logProps, hostsImageRecord)

// 	logOutput := LogOutPut{
// 		UUID:      logProps.UUID.Formula.String,
// 		FilledAtr: logProps.FilledAtr.Formula.Bool,
// 		Title:     logProps.Title.Title,
// 		Date:      logProps.Date.Date,
// 		Mornings:  mornings,
// 		Diet:      diet,
// 		Device:    device,
// 		Published: logProps.Published.Formula.Bool,
// 		TweetUrl:  logProps.TweetUrl.Url,
// 		Diary:     diary,
// 		HostsImage: HostsImage{
// 			HostsLastEditedImage: hostsImage.HostsLastEditedImage,
// 			TodayHostsImage:      hostsImage.TodayHostsImage,
// 		},
// 		WorkOut: struct {
// 			PageID      string `json:"pageId"`
// 			Title       string `json:"title"`
// 			GymLoginPic string `json:"gymLoginPic"`
// 		}{
// 			PageID:      "2d326199ef154c7aa3d48f979e239b00",
// 			Title:       "a",
// 			GymLoginPic: "https://yuya-kanshi.vercel.app/_next/image?url=https%3A%2F%2Fs3.us-west-2.amazonaws.com%2Fsecure.notion-static.com%2Fc5b870f9-f5bb-4f25-8277-d8d6440ec7ed%2F4837EA60-C160-4FA0-86A5-B8E3F40B5C68.jpeg%3FX-Amz-Algorithm%3DAWS4-HMAC-SHA256%26X-Amz-Content-Sha256%3DUNSIGNED-PAYLOAD%26X-Amz-Credential%3DAKIAT73L2G45EIPT3X45%252F20230806%252Fus-west-2%252Fs3%252Faws4_request%26X-Amz-Date%3D20230806T075002Z%26X-Amz-Expires%3D3600%26X-Amz-Signature%3D7a352d3d2f09a4cb48dc4a9eb0020aaf1b332e555576bd93c9db07f0877919e2%26X-Amz-SignedHeaders%3Dhost%26x-id%3DGetObject&w=256&q=75",
// 		},
// 		CalenderPic: "https://yuya-kanshi.vercel.app/_next/image?url=https%3A%2F%2Fs3.us-west-2.amazonaws.com%2Fsecure.notion-static.com%2Fc5b870f9-f5bb-4f25-8277-d8d6440ec7ed%2F4837EA60-C160-4FA0-86A5-B8E3F40B5C68.jpeg%3FX-Amz-Algorithm%3DAWS4-HMAC-SHA256%26X-Amz-Content-Sha256%3DUNSIGNED-PAYLOAD%26X-Amz-Credential%3DAKIAT73L2G45EIPT3X45%252F20230806%252Fus-west-2%252Fs3%252Faws4_request%26X-Amz-Date%3D20230806T075002Z%26X-Amz-Expires%3D3600%26X-Amz-Signature%3D7a352d3d2f09a4cb48dc4a9eb0020aaf1b332e555576bd93c9db07f0877919e2%26X-Amz-SignedHeaders%3Dhost%26x-id%3DGetObject&w=256&q=75",
// 	}

// 	return logOutput
// }

// func getAllLogsDate() ([]string, error) {
// 	logs, err := getLogListFromNow(true)
// 	if err != nil {
// 		return  err
// 	}

// 	var dates []string
// 	for _, log := range logs.TableData {
// 		dates = append(dates, log.Date)
// 	}

// 	return dates, nil
// }
