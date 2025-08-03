ニュース記事を毎朝LINEに通知する

## やりたいこと
* ニュース記事取得のAPIを利用すること
* LINEのAPIを使用し、取得したニュースを自動送信
* AWS Lambda上に実行環境を作成する
* AWS EventBridgeで指定した時刻に上記アプリケーションを起動する

## 前提環境
以下が使用可能なこと  
* docker
* AWS コンソール
* AWS CLI
Amazon ECR にdockerコンテナを登録するため  

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

## 手順
1. Amazon ECR へのコンテナの登録  

2. AWS Lambda の関数作成  
* カスタムランタイムを指定し、ECRに登録したコンテナを選択する。  
* 関数名は line-news-bot  
* 以下を環境変数に設定する。  
```
LINE_API_ACCESS_TOKEN=<<LINE Developer のコンソールから取得したアクセストークン>>  
LINE_API_USER_ID=<<LINE Developer のコンソールから取得したLINEユーザID>>  
LINE_API_URI="https://api.line.me/v2/bot/message/push"  

NEWS_API_URI="https://newsapi.org/v2/everything"  
NEWS_API_PARAMETER="?q=tech&sortBy=relevancy&pageSize=5&apiKey="  
NEWS_API_KEY=<<NewsAPIから取得したAPIキー>>  
```

3. Amazon Eventbridge スケジュールの作成  
* ターゲットには line-news-bot を指定 

* Lambda関数にアクセスするための IAMポリシー、IAMロールを作成する。  
> [!Tips]  
> EventBridge スケジュールでLambdaを定期実行する  
> https://qiita.com/shimabee/items/9cc7451eff44ef7f4769  

下記ドキュメントに従い、欲しい記事の検索条件を設定できます。  
https://newsapi.org/docs/endpoints/everything  

## CI  
* Dockerコンテナのビルド、起動  
下記コマンドを実行  
```
docker build -t line-news-bot-ci .  
docker run -p 9000:8080 -v .\aws-lambda-rie\aws-lambda-rie:/aws-lambda-rie --name line-news-bot-ci --env-file .env -e _HANDLER=function.handler -itd line-news-bot-ci  
```
* curlでの疎通確認  
```
curl -X POST http://localhost:9000/2015-03-31/functions/function/invocations ^
  -H "Content-Type: application/json" ^
  -d "{}"
```
## TODO
- CI/CD   
- 機能の拡張(Googleカレンダーから予定を取得しリマインド)  
- Flex Message を使ったデザイン

## こだわりポイント
・コンテナランライムの設定  
lambdaの実行環境とlambdaサービスの連携  

> [!TIP]
> AWS lambda のRuntime API の  
> https://zenn.dev/hkdord/articles/lambda-handler-deep-dive  
> 
> 【Java】Lombokで冗長コードを削減しよう  
> https://www.casleyconsulting.co.jp/blog/engineer/107/  

・CI環境の作成  
実行環境が AWS Lambda のため、テスト環境はlambdaを再現する必要がある。  
> https://github.com/aws/aws-lambda-runtime-interface-emulator
