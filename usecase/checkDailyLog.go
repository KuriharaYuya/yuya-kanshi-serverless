package usecase

import (
	"sync"

	linepkg "github.com/KuriharaYuya/yuya-kanshi-serverless/repository/line"
	notionpkg "github.com/KuriharaYuya/yuya-kanshi-serverless/repository/notion"
)

func CheckDailyLog() {
	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		_, err := notionpkg.GetLog("2021-09-01")
		if err != nil {
			linepkg.ReplyToUser(err.Error())
			wg.Done()
			return
		}
		// linepkg.ReplyToUser("aa")
		wg.Done()
	}()
	wg.Wait()
	return

}
