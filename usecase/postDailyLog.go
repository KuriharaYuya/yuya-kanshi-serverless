package usecase

import (
	"encoding/json"
	"fmt"
	"sync"

	linepkg "github.com/KuriharaYuya/yuya-kanshi-serverless/repository/line"
	notionpkg "github.com/KuriharaYuya/yuya-kanshi-serverless/repository/notion"
	"github.com/KuriharaYuya/yuya-kanshi-serverless/repository/storage"
)

func PostDailyLog(date string) {
	s := storage.SetUp()

	wg := sync.WaitGroup{}
	wg.Add(1)
	var msg string
	go func() {
		defer wg.Done()
		d, _ := notionpkg.GetLog(date)
		marshaled, _ := json.Marshal(d)
		msg = string(marshaled)
		fmt.Println(msg)

		storage.UploadImages(&d, s)
	}()

	wg.Wait() // この行でgoroutineが完了するのを待ちます。
	linepkg.ReplyToUser(msg)
}
