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
		l, _ := notionpkg.GetLog(date)

		storage.UploadImages(&l, s)

		m := notionpkg.MorningTemplate(&l)
		d := notionpkg.DeviceTemplate(&l)
		h := notionpkg.HealthTemplate(&l)
		notionpkg.AppendContentToPage("fd7093e6ef1a4e1ebeab0a7878c12138", &m, &d, &h)

		marshaled, _ := json.Marshal(l)
		msg = string(marshaled)
		fmt.Println(msg)
	}()

	wg.Wait() // この行でgoroutineが完了するのを待ちます。
	linepkg.ReplyToUser(msg)
}
