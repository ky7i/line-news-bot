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

Keiya Yamaguchi