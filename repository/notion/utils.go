package notionpkg

// func getMinuteDifference(dateString1 string, dateString2 string) string {
// 	layout := "2006-01-02T15:04:05.000Z"

// 	date1, err := time.Parse(layout, dateString1)
// 	if err != nil {
// 		return "日付のパースに失敗しました"
// 	}

// 	date2, err := time.Parse(layout, dateString2)
// 	if err != nil {
// 		return "日付のパースに失敗しました"
// 	}

// 	differenceInMs := date2.Sub(date1).Milliseconds()
// 	differenceInMinutes := math.Abs(float64(differenceInMs)) / (1000 * 60)
// 	absDifference := int(differenceInMinutes)
// 	if differenceInMinutes < 0 {
// 		absDifference = -absDifference
// 	}

// 	if differenceInMinutes < 0 {
// 		return fmt.Sprintf("%d分遅い", absDifference)
// 	} else {
// 		return fmt.Sprintf("%d分早い", absDifference)
// 	}
// }

// func digMorningData(logProps LogProperty, morningActivity MorningActivity) Morning {
// 	return Morning{
// 		MorningImage:                 logProps.MorningImage.Files[0],
// 		MorningActivityTime:          logProps.MorningActivityTime.Date,
// 		MorningActivityEstimatedTime: morningActivity.MorningActivityEstimatedTime.Date,
// 		// TODO: ここで関数を作成して、時間差を計算する必要があるかもしれません
// 		MorningActivityGapMinutes: getMinuteDifference(logProps.MorningActivityTime.Date, morningActivity.MorningActivityEstimatedTime.Date),
// 		MorningActivityLastEdited: morningActivity.MorningActivityLastEdited.LastEditedTime,
// 		MorningTargetPlace:        morningActivity.MorningTargetPlace.Title,
// 	}
// }

// func digDiaryData(logProps LogProperty) Diary {
// 	return Diary{
// 		IsDiaryDone:   logProps.IsDiaryDone.Checkbox,
// 		IsChatLogDone: logProps.IsChatLogDone.Checkbox,
// 	}
// }

// func digHostImageData(logProps LogProperty, hostImageProps HostsImageRecord) HostsImage {
// 	return HostsImage{
// 		HostsLastEditedImage: hostImageProps.HostsLastEditedImage.Files[0],
// 		TodayHostsImage:      logProps.TodayHostsImage.Files[0],
// 	}
// }

// func getCalorieDifference(monthlyCalorie, todayCalorie int) string {
// 	difference := monthlyCalorie - todayCalorie

// 	if difference < 0 {
// 		return fmt.Sprintf("%vkcal多い", math.Abs(float64(difference)))
// 	} else if difference > 0 {
// 		return fmt.Sprintf("%vkcal少ない", difference)
// 	} else {
// 		return "目標通り"
// 	}
// }

// func isUpperToMsg(isUpper bool) string {
// 	if isUpper {
// 		return "上限"
// 	} else {
// 		return "下限"
// 	}
// }

// func DigDietData(logProps LogProperty, monthlyProps MonthlyRecord) Diet {
// 	return Diet{
// 		MyFitnessPal:          logProps.MyFitnessPal.Files[0],
// 		TodayCalorie:          logProps.TodayCalorie.Number,
// 		MonthlyCalorie:        monthlyProps.MonthlyCalorie.Number,
// 		TodayCalorieGap:       getCalorieDifference(int(monthlyProps.MonthlyCalorie.Number), int(logProps.TodayCalorie.Number)),
// 		MonthlyCalorieIsUpper: isUpperToMsg(monthlyProps.MonthlyCalorieIsUpper.Checkbox),
// 	}
// }

// func GetScreenTimeDifference(monthlyScreenTime int16, todayScreenTime int16) string {
// 	difference := monthlyScreenTime - todayScreenTime
// 	if difference < 0 {
// 		return fmt.Sprintf("%d分多い", -difference)
// 	}
// 	if difference > 0 {
// 		return fmt.Sprintf("%d分少ない", difference)
// 	} else {
// 		return "目標通り"
// 	}
// }

// func DigDeviceData(logProps LogProperty, monthlyProps MonthlyRecord) Device {
// 	screenTimeGapMinutes := GetScreenTimeDifference(monthlyProps.MonthlyScreenTime.Number, logProps.TodayScreenTime.Number)

// 	return Device{
// 		ScreenTime:           logProps.ScreenTime.Files[0],
// 		TodayScreenTime:      logProps.TodayScreenTime.Number,
// 		ScreenTimeGapMinutes: screenTimeGapMinutes,
// 		MonthlyScreenTime:    monthlyProps.MonthlyScreenTime.Number,
// 	}
// }

// const WRAPTAS_URL string = "https://yuyanki.wraptas.site/"
