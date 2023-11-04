package tweet

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	notionpkg "github.com/KuriharaYuya/yuya-kanshi-serverless/repository/notion"
	"github.com/KuriharaYuya/yuya-kanshi-serverless/repository/storage"
	repoutils "github.com/KuriharaYuya/yuya-kanshi-serverless/repository/utils"
)

func CallVercelTwitterAPI(log *notionpkg.LifeLog) string {
	// タイトルを取得
	title := log.Title
	// 日付を取得
	date := log.Date

	uuid := log.UUID
	// ツイート文を生成
	tweetText := generateTweetText(title, date, uuid)

	//カレンダーの画像を取得
	s3URL := repoutils.GetImageExternalURl(date, storage.CalenderPic)
	// vercelをcallしてツイートを投稿
	tweetId := callVercelAPI(tweetText, s3URL)
	return tweetId

}

func generateTweetText(t string, d string, u string) string {
	baseURL := "https://yuyanki.notion.site/"
	url := baseURL + u
	nl := "\n"
	txt := d + "の記録" + nl + t + nl + "下記リンクで今日行った詳細を確認することができます！" + nl + url
	return txt
}

const (
	prodApiURL = "https://yuya-kanshi.vercel.app/api/tweet/postTweetFromLambda"
	devApiURL  = "https://7071-58-156-74-242.ngrok-free.app/api/tweet/postTweetFromLambda"
)

func callVercelAPI(t string, s3URL string) string {
	data := map[string]string{
		"text":     t,
		"mediaUrl": s3URL,
	}
	// media
	fmt.Println("mediaUrl:", s3URL)

	bodyData, err := json.Marshal(data)
	resForVercel, err := http.Post(prodApiURL, "application/json", bytes.NewBuffer(bodyData))
	if err != nil {
		fmt.Println("Error sending request:", err)
		return ""
	}
	defer resForVercel.Body.Close()

	body, err := ioutil.ReadAll(resForVercel.Body)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return ""
	}
	// ex body) {lastTweetId: 1426210000000000000}
	// bodyをjsonに変換
	var resp map[string]string
	json.Unmarshal(body, &resp)
	lastTweetId := resp["lastTweetId"]
	fmt.Println("Response from Vercel:", string(body))
	return lastTweetId
}
