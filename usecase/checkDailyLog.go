package usecase

import (
	"fmt"
	"sync"

	linepkg "github.com/KuriharaYuya/yuya-kanshi-serverless/repository/line"
	notionpkg "github.com/KuriharaYuya/yuya-kanshi-serverless/repository/notion"
)

func CheckDailyLog(date string) {
	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		logData, err := notionpkg.GetLog(date)
		if err != nil {
			fmt.Println(err)
			return
		}

		title := logData.Title
		tweetURL := logData.TweetURL
		textData := "タイトル: " + title + "\n" + "ツイートURL: " + tweetURL
		linepkg.ReplyToUser(textData)
		wg.Done()
	}()
	wg.Wait()
	return

}
