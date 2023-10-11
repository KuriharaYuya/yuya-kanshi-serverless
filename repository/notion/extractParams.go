package notionpkg

// 与えられたURLからtitleを取得して構造体に格納する

import (
	"fmt"

	"github.com/jomei/notionapi"
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

	// filledAttrebuteProperty, exists := pageURL.Properties["filledAtr"]
	// if(exists){
	// 	logProp.FilledAtr = string(filledAttrebuteProperty.GetType())
	// }else{
	// 	fmt.Println("filledAttrebuteProperty is not exists")
	// }
	morningImageProperty, exists := pageURL.Properties["morningImage"]
	if(exists){
		logProp.MorningImage = string(morningImageProperty.GetType())
	}else{
		fmt.Println("morningImageProperty is not exists")
	}
	myFitnessPalProperty, exists := pageURL.Properties["myFitnessPal"]
	if(exists){
		logProp.MyFitnessPal = string(myFitnessPalProperty.GetType())
	}else{
		fmt.Println("myFitnessPalProperty is not exists")
	}
	todayCalorieProperty, exists := pageURL.Properties["todayCalorie"]
	if(exists){
		logProp.TodayCalorie = string(todayCalorieProperty.GetType())
	}else{
		fmt.Println("todayCalorieProperty is not exists")
	}
	screenTimeProperty, exists := pageURL.Properties["screenTime"]
	if(exists){
		logProp.ScreenTime = string(screenTimeProperty.GetType())
	}else{
		fmt.Println("screenTimeProperty is not exists")
	}
 	todayScreenTimeProperty, exists := pageURL.Properties["todayScreenTime"]
	if(exists){
		logProp.TodayScreenTime = string(todayScreenTimeProperty.GetType())
	}else{
		fmt.Println("todayScreenTimeProperty is not exists")
	}
	morningActivityTimeProperty, exists := pageURL.Properties["morningActivityTime"]
	if(exists){
		logProp.MorningActivityTime = string(morningActivityTimeProperty.GetType())
	}else{
		fmt.Println("morningActivityTimeProperty is not exists")
	}
	publishedProperty, exists := pageURL.Properties["published"]
	if(exists){
		logProp.Published = string(publishedProperty.GetType())
	}else{
		fmt.Println("publishedProperty is not exists")
	}
	tweetUrlProperty, exists := pageURL.Properties["tweetUrl"]
	if(exists){
		logProp.TweetUrl = string(tweetUrlProperty.GetType())
	}else{
		fmt.Println("tweetUrlProperty is not exists")
	}
	isDiaryDoneProperty, exists := pageURL.Properties["isDiaryDone"]
	if(exists){
		logProp.IsDiaryDone = string(isDiaryDoneProperty.GetType())
	}else{
		fmt.Println("isDiaryDoneProperty is not exists")
	}
	isChatLogDoneProperty, exists := pageURL.Properties["isChatLogDone"]
	if(exists){
		logProp.IsChatLogDone = string(isChatLogDoneProperty.GetType())
	}else{
		fmt.Println("isChatLogDoneProperty is not exists")
	}
	todayHostsImageProperty, exists := pageURL.Properties["todayHostsImage"]
	if(exists){
		logProp.TodayHostsImage = string(todayHostsImageProperty.GetType())
	}else{
		fmt.Println("todayHostsImageProperty is not exists")
	}

	if len(pageURL.Title) > 0 {
			logProp.Title = pageURL.Title[0].PlainText
	}
	if !pageURL.CreatedTime.IsZero() {
		logProp.Date = pageURL.CreatedTime.String()
	}
	return logProp, err // Titleが空の場合は空文字列を返す
}

func Println(morningImageProperty notionapi.PropertyConfig) {
	panic("unimplemented")
}
