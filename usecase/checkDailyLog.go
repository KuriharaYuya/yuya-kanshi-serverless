package usecase

import (
	"fmt"
	"sync"

	linepkg "github.com/KuriharaYuya/yuya-kanshi-serverless/repository/line"
	notionpkg "github.com/KuriharaYuya/yuya-kanshi-serverless/repository/notion"
)

func CheckDailyLog() {
	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		url, err := notionpkg.GetLog()
		fmt.Println(url)
		if err != nil {
			linepkg.ReplyToUser("日報の取得に失敗しました")
			wg.Done()
			return
		}
		linepkg.ReplyToUser("aa")
		wg.Done()
	}()
	wg.Wait()
	return

}
