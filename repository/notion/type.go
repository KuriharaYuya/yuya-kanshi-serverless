package notionpkg

import "github.com/jomei/notionapi"

// こんな感じ

type LogProperty struct {
	UUID      notionapi.TitleProperty
	FilledAtr notionapi.FormulaProperty
}

type Log struct {
	UUID      string
	FilledAtr bool
}
