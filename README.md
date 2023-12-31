# YUYA-KANSHI

自身を監視するアプリ
システム受託会社を起業し、仕事へのコミットが完全に自分次第になり、堕落してしまう日々が続いていました。
そのような背景のもと、仕事へのコミットを最大化するために開発しました。

開発期間:　 2023-11-05 ~ 11-8 おおよそ 30h

# 目次

- [概要](#概要)
- [システム構成](#システム構成)
- [システムの機能](#システムの機能)
- [前身システム課題と解決策](#前身システム課題と解決策)
- [技術選定と開発時に工夫したこと](#技術選定と開発時に工夫したこと)
- [今後](#今後)

おおよそ 3 分ほどで読めます

## 概要

この web アプリは[前身となるアプリ](https://github.com/KuriharaYuya/yuya-kanshi)が存在していましたが、今回のアプリにリプレイスしました。
パフォーマンスが低かったため、使いずらい、の 2 点の理由です。

**"朝活の開始時間"**,**"スマホのスクリーンタイム"**, **"健康(食事・筋トレ)"** というこれら 3 点において毎日アプリで記録をとり、アプリを用いて公開します。

開発背景は前身アプリの readme をご覧ください。この readme では割愛します。

## システム構成

vercel の tweet 専用サーバーについての説明は割愛します。

| 項目                     | 使用技術                     |
| ------------------------ | ---------------------------- |
| presentation             | LINE・Notion・Twitter        |
| 管理画面・データ入力機構 | Notion (Web & iOS アプリ)    |
| サーバー                 | AWS Lambda(go lang runtime)  |
| インフラプロビジョニング | serverless framework         |
| API                      | notionAPI,LINE messaging api |



<img width="1438" alt="Screenshot 2023-11-08 at 9 53 41" src="https://github.com/KuriharaYuya/yuya-kanshi-serverless/assets/109059044/8cabc42d-6e39-4c3a-a349-bde61579518c">




## システムの機能
<img width="733" alt="Screenshot 2023-11-08 at 7 13 55" src="https://github.com/KuriharaYuya/yuya-kanshi-serverless/assets/109059044/987b8e04-703b-48e7-97cc-ffd64017f06b">
<img width="730" alt="Screenshot 2023-11-08 at 7 14 07" src="https://github.com/KuriharaYuya/yuya-kanshi-serverless/assets/109059044/0ea8d1b0-ffd0-4fc5-98fc-cb0938e6292a">
<img width="730" alt="Screenshot 2023-11-08 at 7 14 16" src="https://github.com/KuriharaYuya/yuya-kanshi-serverless/assets/109059044/01bfc988-eb75-4b71-a54d-49105b976c0e">

<img width="506" alt="Screenshot 2023-11-14 at 23 55 46" src="https://github.com/KuriharaYuya/yuya-kanshi-serverless/assets/109059044/7c652ac1-0448-42a7-8308-bae2de235fa2">


## 前身システム課題と解決策

機能面

- 使いにくい。入力項目が多く、投稿作業も煩雑である。それゆえ面倒くさくなってしまい、投稿しなくなっていた。

  -> 入力項目を減らした。notion の関数・リレーション&ロールアップの使用で解決。

  -> 投稿作業はシステムで半自動化

  -> 投稿などのハンドリングは line bot から行えるように変更。より使いやすくなった。

- 監視結果を対外的に公開するページがわかりにくかった。ワイヤーデザイン程度の css しか書いていなかった。

  -> notion の公開ページで表示することで解決

  ->報告専用の bot を設けることで、友人との意識高い系グループに毎日報告できるようになった。サボったらすぐにバレるのでサボりにくくなった。

  非機能面

  - 高画質画像を扱うと、システムが落ちる。
    next.js on vercel の運用であり、vercel 上では 1 関数の最大実行時間が 10 秒であった。監視項目の画像が高画質である場合など、しばしば 10 秒を超えていたためリプレイスが必須であった。

    -> 該当機能を lambda(go lang)に載せ替えることで解決。lambda のランタイムが高パフォーマンスで、全体の処理時間も 30 秒 -> 8 秒程度に短縮された。

  - 煩雑なコードと分離できていない責務
    元々フロントエンドアプリに api という形で投稿機能を実装していたため、アプリ自体もコードも責務が分離できていなかった。開発効率低下、バグ特定に多大な時間が必要になるなどの弊害が起きていた。

    -> クリーンアーキテクチャの思想を取り入れたコーディングと、lambda に必要機能を切り出したことで解決。

## 技術選定と開発時に工夫したこと

いくつかピックアップします。

技術選定

- lambda
  本システムはイベント駆動型であるため、かつリクエストは自分しかいないため、時稼働するサーバーは必要ない。したがって lambda を選定。
- serverless framework
  開発効率・デプロイとプロビジョニングがしやすい、という理由で選定。
  lambda での開発を行う際にネックなのが 2 点。
  1.AWS 提供のエディタでコーディングする必要がある。go の場合型エラーを確認しながら開発したいのに、できない。 2.ライブラリのためのレイヤー追加が面倒。これはデプロイ前にビルドすることで解決された。aws 上でコーディングしていたらそれはできない。
  上記２点を解決するのが serverless framework であったため。

  他の選択肢として terraform があったが、そこまでの機能が必要ないため却下した。
  また、serverless framework は lambda のデプロイ用の雛形があったことも大きい。

工夫したこと

- ローカルでの開発基盤構築
  環境変数を用いてローカルで開発しやすいようにした。イベント駆動で lambda で動くためのコードを書いているため、ローカルでの動作確認がしにくかったため基盤環境を開発した。

## 今後

このアプリを用いて強固な生活習慣を構築し、以下のようなことに取り組む予定です：

- 自分の会社の 2~4 月の月間平均営業利益を 100 万以上にする
- 週に 10h 以上本を読書して、仕事で必要になるであろう領域を学習する。
  直近は web マーケを学びます。

今回作ったアプリを用いて、コツコツ努力してこれらの目標を達成します！！！

お読みいたいだきありがとうございました！！

# for dev

## デプロイ

make build && sls deploy
