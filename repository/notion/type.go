package notionpkg

import "github.com/jomei/notionapi"

type DebugLogProperty struct {
	Name         *notionapi.TitleProperty
	AllowPublish *notionapi.CheckboxProperty
}

type LifeLogProperty struct {
	UUID                *notionapi.FormulaProperty
	FilledAtr           *notionapi.FormulaProperty
	Title               *notionapi.TitleProperty
	Date                *notionapi.DateProperty
	MorningImage        *notionapi.FilesProperty
	MyFitnessPal        *notionapi.FilesProperty
	TodayCalorie        *notionapi.NumberProperty
	ScreenTime          *notionapi.FilesProperty
	TodayScreenTime     *notionapi.NumberProperty
	MorningActivityTime *notionapi.DateProperty
	Published           *notionapi.FormulaProperty
	TweetURL            *notionapi.URLProperty
	IsDiaryDone         *notionapi.CheckboxProperty
	IsChatLogDone       *notionapi.CheckboxProperty
	TodayHostsImage     *notionapi.FilesProperty
}

type LifeLog struct {
	UUID                string
	FilledAtr           bool
	Title               string
	Date                string
	MorningImage        string
	MyFitnessPal        string
	TodayCalorie        int
	ScreenTime          string
	TodayScreenTime     int
	MorningActivityTime string
	Published           bool
	TweetURL            string
	IsDiaryDone         bool
	IsChatLogDone       bool
	TodayHostsImage     string
}

const (
	NotionUUID                = "uuid"
	NotionFilledAtr           = "filledAtr"
	NotionTitle               = "title"
	NotionDate                = "date"
	NotionMorningImage        = "morningImage"
	NotionMyFitnessPal        = "myFitnessPal"
	NotionTodayCalorie        = "todayCalorie"
	NotionScreenTime          = "screenTime"
	NotionTodayScreenTime     = "todayScreenTime"
	NotionMorningActivityTime = "morningActivityTime"
	NotionPublished           = "published"
	NotionTweetURL            = "tweetUrl"
	NotionIsDiaryDone         = "isDiaryDone"
	NotionIsChatLogDone       = "isChatLogDone"
	NotionTodayHostsImage     = "todayHostsImage"
)
