package usecase

import (
	"sync"

	linepkg "github.com/KuriharaYuya/yuya-kanshi-serverless/repository/line"
)

func CheckDailyLog() {
	wg := sync.WaitGroup{}
	wg.Add(1)

	go func() {
		linepkg.ReplyToUser("hello")
		wg.Done()
	}()
	wg.Wait()
	return

}
