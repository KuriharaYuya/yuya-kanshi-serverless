package notionpkg

import (
	"fmt"

	"github.com/jomei/notionapi"
)

// func DebugStructExchange
func debug(results *notionapi.DatabaseQueryResponse) {
	resultProps := results.Results[0].Properties
	nameProp, nameOk := resultProps["Name"].(*notionapi.TitleProperty)
	allowPublishProp, allowPublishOk := resultProps["allowPublish"].(*notionapi.CheckboxProperty)

	// name ok
	if !nameOk || !allowPublishOk {
		fmt.Println("Error extracting properties")
		fmt.Println(nameProp, "resultProps")
		return
	}
	debugLog := DebugLogProperty{
		Name:         nameProp,
		allowPublish: allowPublishProp,
	}

	fmt.Println(debugLog, "debugLog")
}
