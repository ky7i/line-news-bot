# line-news-bot
ニュース記事を毎朝LINEに通知する

## やりたいこと
* ニュース記事取得のAPIを利用すること
* LINEのAPIを使用し、取得したニュースを自動送信
* AWS Lambda上に実行環境を作成する
* AWS EventBridgeで指定した時刻に上記アプリケーションを起動する

## 前提環境
以下が使用可能なこと
* docker
* AWS

## 使用技術一覧
* Go
* AWS Lambda
* AWS EventBridge

## 使用API一覧
* Messaging API  
https://developers.line.biz/ja/services/messaging-api/  
LINEのAPI  
LINEへのメッセージ送信に使用  

* News API  
https://newsapi.org/  
ニュース記事の取得を行うAPI  
海外のニュースがメイン  

## 手順
// TODO

* プロジェクトのルートディレクトリに .env ファイルを作成する。  
下記を .env に記述する。  
```
LINE_API_ACCESS_TOKEN=<<LINE Developer のコンソールから取得したアクセストークン>>  
LINE_API_USER_ID=<<LINE Developer のコンソールから取得したLINEユーザID>>  
LINE_API_URI="https://api.line.me/v2/bot/message/push"  

NEWS_API_URI="https://newsapi.org/v2/everything"  
NEWS_API_PARAMETER="?q=tech&sortBy=relevancy&pageSize=5&apiKey="  
NEWS_API_KEY=<<NewsAPIから取得したAPIキー>>  
```

NES_API_PARAMETERはニュース記事の検索条件  
下記ドキュメントに従い、欲しい記事の検索条件を設定できます。  
https://newsapi.org/docs/endpoints/everything  

## TODO
・本当は日本語のニュース記事が欲しい
