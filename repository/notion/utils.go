package notionpkg

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/jomei/notionapi"
)

func serializeToLogProp(result *notionapi.DatabaseQueryResponse) (LifeLog, error) {
	if len(result.Results) == 0 {
		return LifeLog{}, fmt.Errorf("no results found")
	}

	resultProps := result.Results[0].Properties

	uuidProp, ok := resultProps[NotionUUID].(*notionapi.FormulaProperty)
	if !ok || uuidProp == nil {
		fmt.Println(uuidProp, "と", resultProps)
		return LifeLog{}, fmt.Errorf("failed to cast UUID property")
	}

	filledAtrProp, ok := resultProps[NotionFilledAtr].(*notionapi.FormulaProperty)
	if !ok || filledAtrProp == nil {
		fmt.Println(filledAtrProp, "と", resultProps)
		return LifeLog{}, fmt.Errorf("failed to cast FilledAtr property")
	}

	titleProp, ok := resultProps[NotionTitle].(*notionapi.TitleProperty)
	if !ok || titleProp == nil {
		return LifeLog{}, fmt.Errorf("failed to cast Title property")
	}

	dateProp, ok := resultProps[NotionDate].(*notionapi.DateProperty)
	if !ok || dateProp == nil {
		return LifeLog{}, fmt.Errorf("failed to cast Date property")
	}

	morningImageProp, ok := resultProps[NotionMorningImage].(*notionapi.FilesProperty)
	if !ok || morningImageProp == nil {
		return LifeLog{}, fmt.Errorf("failed to cast MorningImage property")
	}

	myFitnessPalProp, ok := resultProps[NotionMyFitnessPal].(*notionapi.FilesProperty)
	if !ok || myFitnessPalProp == nil {
		return LifeLog{}, fmt.Errorf("failed to cast MyFitnessPal property")
	}

	todayCalorieProp, ok := resultProps[NotionTodayCalorie].(*notionapi.NumberProperty)
	if !ok || todayCalorieProp == nil {
		return LifeLog{}, fmt.Errorf("failed to cast TodayCalorie property")
	}

	screenTimeProp, ok := resultProps[NotionScreenTime].(*notionapi.FilesProperty)
	if !ok || screenTimeProp == nil {
		return LifeLog{}, fmt.Errorf("failed to cast ScreenTime property")
	}

	todayScreenTimeProp, ok := resultProps[NotionTodayScreenTime].(*notionapi.NumberProperty)
	if !ok || todayScreenTimeProp == nil {
		return LifeLog{}, fmt.Errorf("failed to cast TodayScreenTime property")
	}

	morningActivityTimeProp, ok := resultProps[NotionMorningActivityTime].(*notionapi.DateProperty)
	if !ok || morningActivityTimeProp == nil {
		return LifeLog{}, fmt.Errorf("failed to cast MorningActivityTime property")
	}

	publishedProp, ok := resultProps[NotionPublished].(*notionapi.FormulaProperty)
	if !ok || publishedProp == nil {
		return LifeLog{}, fmt.Errorf("failed to cast Published property")
	}

	tweetURLProp, ok := resultProps[NotionTweetURL].(*notionapi.URLProperty)
	if !ok || tweetURLProp == nil {
		return LifeLog{}, fmt.Errorf("failed to cast TweetURL property")
	}

	isDiaryDoneProp, ok := resultProps[NotionIsDiaryDone].(*notionapi.CheckboxProperty)
	if !ok || isDiaryDoneProp == nil {
		return LifeLog{}, fmt.Errorf("failed to cast IsDiaryDone property")
	}

	isChatLogDoneProp, ok := resultProps[NotionIsChatLogDone].(*notionapi.CheckboxProperty)
	if !ok || isChatLogDoneProp == nil {
		return LifeLog{}, fmt.Errorf("failed to cast IsChatLogDone property")
	}

	todayHostsImageProp, ok := resultProps[NotionTodayHostsImage].(*notionapi.FilesProperty)
	if !ok || todayHostsImageProp == nil {
		return LifeLog{}, fmt.Errorf("failed to cast TodayHostsImage property")
	}

	log := LifeLog{
		UUID:                uuidProp.Formula.String,
		FilledAtr:           filledAtrProp.Formula.Boolean,
		Title:               titleProp.Title[0].PlainText,
		Date:                convertDateFormat(dateProp.Date.Start.String()),
		MorningImage:        morningImageProp.Files[0].File.URL,
		MyFitnessPal:        myFitnessPalProp.Files[0].File.URL,
		TodayCalorie:        int(todayCalorieProp.Number),
		ScreenTime:          screenTimeProp.Files[0].File.URL,
		TodayScreenTime:     int(todayScreenTimeProp.Number),
		MorningActivityTime: morningActivityTimeProp.Date.Start.String(),
		Published:           publishedProp.Formula.Boolean,
		TweetURL:            tweetURLProp.URL,
		IsDiaryDone:         isDiaryDoneProp.Checkbox,
		IsChatLogDone:       isChatLogDoneProp.Checkbox,
		TodayHostsImage:     todayHostsImageProp.Files[0].File.URL,
	}
	return log, nil
}

func parseTime(date string) time.Time {
	// dateを正規表現で検証する。
	re := regexp.MustCompile(`\d{4}-\d{2}-\d{2}`)
	if !re.MatchString(date) {
		panic("日付の形式が正しくありません" + date)
	}

	// parsedTime, err := time.Parse("2006-01-02", date)
	parsedTime, err := time.Parse("2006-01-02", date)
	if err != nil {
		panic(err)
	}
	return parsedTime
}

func createDateQuery(targetDate notionapi.Date) *notionapi.DatabaseQueryRequest {
	q := &notionapi.DatabaseQueryRequest{
		Filter: notionapi.PropertyFilter{
			Property: NotionDate,
			Date: &notionapi.DateFilterCondition{
				Equals: &targetDate,
			},
		},
	}
	return q
}

func convertDateFormat(dateStr string) string {
	parts := strings.Split(dateStr, "T")
	datePart := parts[0]
	return datePart
}
