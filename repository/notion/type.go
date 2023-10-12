package notionpkg

// notion Built-in types
type FileProperty struct {
	Files []string `json:"files"`
}

type LastEditedTimeProperty struct {
	LastEditedTime string `json:"lastEditedTime"`
}

type NumberProperty struct { 
	Number int16 `json:"number"`
}

type TitleProperty struct { 
	Title string `json:"plain_text"`
}

type DateProperty struct {
	Date string `json:"date"`
}

type FormulaBooleanProperty struct {
	Fomula struct {Bool bool `json:"boolean"`} `json:"formula"`
}

type FormulaStringProperty struct {
	Formula struct {
		String string `json:"string"`
	} `json:"formula"`
}

type UrlProperty struct {
	Url string `json:"url"`
}

type BooleanProperty struct {
	Checkbox bool `json:"boolean"`
}

// notion API response type components
type Morning struct {
	MorningImage              string `json:"morningImage"`
	MorningActivityTime       string `json:"morningActivityTime"`
	MorningActivityEstimatedTime string `json:"morningActivityEstimatedTime"`
	MorningActivityGapMinutes  string `json:"morningActivityGapMinutes"`
	MorningActivityLastEdited  string `json:"morningActivityLastEdited"`
	MorningTargetPlace         string `json:"morningTargetPlace"`
}

type Diet struct {
	MyFitnessPal           string `json:"myFitnessPal"`
	TodayCalorie           int16    `json:"todayCalorie"`
	MonthlyCalorie         int16    `json:"monthlyCalorie"`
	TodayCalorieGap        string `json:"todayCalorieGap"`
	MonthlyCalorieIsUpper  string `json:"monthlyCalorieIsUpper"`
}

type Device struct {
	ScreenTime             string `json:"screenTime"`
	TodayScreenTime        int16    `json:"todayScreenTime"`
	ScreenTimeGapMinutes   string `json:"screenTimeGapMinutes"`
	MonthlyScreenTime      int16    `json:"monthlyScreenTime"`
}

type Diary struct {
	IsDiaryDone   bool `json:"isDiaryDone"`
	IsChatLogDone bool `json:"isChatLogDone"`
}

type HostsImage struct {
	HostsLastEditedImage string `json:"hostsLastEditedImage"`
	TodayHostsImage      string `json:"todayHostsImage"`
}


// notion API response type components
// ----------------------------------------

// ----------------------------------------
// notion API response types



type LogProperty struct {
	UUID            FormulaStringProperty `json:"uuid"`
	FilledAtr       FormulaBooleanProperty `json:"filledAtr"`
	Title           TitleProperty         `json:"title"`
	Date            DateProperty          `json:"date"`
	MorningImage    FileProperty          `json:"morningImage"`
	MyFitnessPal    FileProperty          `json:"myFitnessPal"`
	TodayCalorie    NumberProperty        `json:"todayCalorie"`
	ScreenTime      FileProperty          `json:"screenTime"`
	TodayScreenTime NumberProperty        `json:"todayScreenTime"`
	MorningActivityTime DateProperty     `json:"morningActivityTime"`
	Published       FormulaBooleanProperty `json:"published"`
	TweetUrl        UrlProperty            `json:"tweetUrl"`
	IsDiaryDone     BooleanProperty        `json:"isDiaryDone"`
	IsChatLogDone   BooleanProperty        `json:"isChatLogDone"`
	TodayHostsImage FileProperty          `json:"todayHostsImage"`
}

type MorningActivity struct {
	MorningActivityEstimatedTime DateProperty      `json:"morningActivityEstimatedTime"`
	MorningActivityLastEdited    LastEditedTimeProperty `json:"morningActivityLastEdited"`
	MorningTargetPlace           TitleProperty      `json:"morningTargetPlace"`
}

type MonthlyRecord struct {
	MonthlyScreenTime      NumberProperty   `json:"monthlyScreenTime"`
	MonthlyCalorie         NumberProperty   `json:"monthlyCalorie"`
	MonthlyCalorieIsUpper  BooleanProperty   `json:"monthlyCalorieIsUpper"`
}

type HostsImageRecord struct {
	HostsLastEditedImage FileProperty   `json:"hostsLastEditedImage"`
	ChangedReason       TitleProperty  `json:"changedReason"`
}

// ----------------------------------------
// notion API response types

// ----------------------------------------
// notion API response in the end

type LogOutPut struct {
	UUID      string        `json:"uuid"`
	Title     string        `json:"title"`
	Date      string        `json:"date"`
	Mornings  Morning       `json:"mornings"`
	Diet      Diet          `json:"diet"`
	Device    Device        `json:"device"`
	Published bool          `json:"published"`
	FilledAtr bool          `json:"filledAtr"`
	TweetUrl  string        `json:"tweetUrl"`
	Diary     Diary         `json:"diary"`
	HostsImage HostsImage    `json:"hostsImage"`
	WorkOut   struct {
		PageID    string `json:"pageId"`
		Title     string `json:"title"`
		GymLoginPic string `json:"gymLoginPic"`
	} `json:"workOut"`
	CalenderPic string       `json:"calenderPic"`
}

// ----------------------------------------
// notion API response at LIST
type LogListProperty struct {
	UUID      FormulaStringProperty    `json:"uuid"`
	Title     TitleProperty            `json:"title"`
	TweetUrl  UrlProperty              `json:"tweetUrl"`
	Date      DateProperty             `json:"date"`
	Published FormulaBooleanProperty   `json:"published"`
}

type LogListPropertyForGitLikeCalender struct {
	Properties LogListProperty `json:"properties"`
}

// ----------------------------------------
// notion API response at LIST
// ----------------------------------------

// ----------------------------------------
// notion API response at Table
// ----------------------------------------

type LogTableProperty []LogListProperty
// ----------------------------------------
// notion API response at Table
// ----------------------------------------