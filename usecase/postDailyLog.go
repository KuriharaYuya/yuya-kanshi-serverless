package usecase

import (
	"encoding/json"
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

		l, valid := notionpkg.ValidateLog(date)
		marshaled, _ := json.Marshal(l)
		msg = string(marshaled)
		if valid {
			linepkg.ReplyToUser("バリデーションを通過しました")
		} else {
			linepkg.ReplyToUser("バリデーションに失敗しました")
			linepkg.ReplyToUser(msg)
			return
		}

		storage.UploadImages(&l, s)
		linepkg.ReplyToUser("画像のアップロードが完了しました")

		s := notionpkg.DiaryHeaderTemplate(&l)

		m := notionpkg.MorningTemplate(&l)
		d := notionpkg.DeviceTemplate(&l)
		h := notionpkg.HealthTemplate(&l)
		notionpkg.AppendContentToPage(l.UUID, &s, &m, &d, &h)
		linepkg.ReplyToUser("Notionへの書き込みが完了しました")

	}()

	wg.Wait() // この行でgoroutineが完了するのを待ちます。
	linepkg.ReplyToUser(msg)
}
