package notionpkg

// 与えられたURLからtitleを取得して構造体に格納する

import (
	"fmt"
)

type LogProperty struct {
	Uuid string `json:"uuid"`
	FilledAtr bool `json:"filledAtr"`
	Title string `json:"title"`
	Date string `json:"date"`
	MorningImage string `json:"morningImage"`
	MyFitnessPal string `json:"myFitnessPal"`
	TodayCalorie string `json:"todayCalorie"`
	ScreenTime string `json:"screenTime"`
	TodayScreenTime string `json:"todayScreenTime"`
	MorningActivityTime string `json:"morningActivityTime"`
	Published string `json:"published"`
	TweetUrl string `json:"tweetUrl"`
	IsDiaryDone string `json:"isDiaryDone"`
	IsChatLogDone string `json:"isChatLogDone"`
	TodayHostsImage string `json:"todayHostsImage"`
}

func GetLogProperty() (LogProperty, error)  {
	pageURL, err  := GetLog()
	if err != nil {
		fmt.Println(err)
	}

	logProp := LogProperty{
		Uuid: "",
		FilledAtr: false,
		Title: "",
		Date: "",
		MorningImage: "",
		MyFitnessPal: "",
		TodayCalorie: "",
		ScreenTime: "",
		TodayScreenTime: "",
		MorningActivityTime: "",
		Published: "",
		TweetUrl: "",
		IsDiaryDone: "",
		IsChatLogDone: "",
		TodayHostsImage: "",
	}

	if len(pageURL.ID) > 0 {
		logProp.Uuid = pageURL.ID.String()
	}
	// if len(pageURL.Properties) > 0 {
	// 	logProp.FilledAtr = pageURL.Properties["FilledAtr"].GetType()
	// }
		morningImageProperty, exists := pageURL.Properties["MorningImage"]
		if exists {
			logProp.MorningImage = string(morningImageProperty.GetType())
		}
		// logProp.MorningImage = pageURL.Properties["MorningImage"]
		// logProp.MyFitnessPal = string(pageURL.Properties["MyFitnessPal"].GetType())
		// logProp.TodayCalorie = string(pageURL.Properties["TodayCalorie"].GetType())
		// logProp.ScreenTime = string(pageURL.Properties["ScreenTime"].GetType())
		// logProp.TodayScreenTime = string(pageURL.Properties["TodayScreenTime"].GetType())
		// logProp.MorningActivityTime = string(pageURL.Properties["MorningActivityTime"].GetType())
		// logProp.Published = string(pageURL.Properties["Published"].GetType())
		// logProp.TweetUrl = string(pageURL.Properties["TweetUrl"].GetType())
		// logProp.IsDiaryDone = string(pageURL.Properties["IsDiaryDone"].GetType())
		// logProp.IsChatLogDone = string(pageURL.Properties["IsChatLogDone"].GetType())
		// logProp.TodayHostsImage = string(pageURL.Properties["TodayHostsImage"].GetType())
	if len(pageURL.Title) > 0 {
			logProp.Title = pageURL.Title[0].PlainText
	}
	if !pageURL.CreatedTime.IsZero() {
		logProp.Date = pageURL.CreatedTime.String()
	}
	return logProp, err // Titleが空の場合は空文字列を返す
}
