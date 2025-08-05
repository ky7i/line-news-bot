## LINE News Bot

Send news articles to LINE every morning using AWS Lambda and EventBridge.

---

## Features
- Fetch news articles using NewsAPI
- Automatically send fetched news to LINE via the Messaging API
- Run as a container on AWS Lambda
- Scheduled execution with AWS EventBridge

## Prerequisites
- Docker
- AWS Console
- AWS CLI

*These are required to build and register the Docker container to Amazon ECR.*

## Technologies Used
- Go
- AWS Lambda
- AWS EventBridge

## APIs Used
- [LINE Messaging API](https://developers.line.biz/en/services/messaging-api/): Send messages to LINE users
- [NewsAPI](https://newsapi.org/): Fetch news articles

## Setup Steps
1. **Register the container to Amazon ECR**

2. **Create an AWS Lambda function**
   - Specify a custom runtime and select the container registered in ECR
   - Function name: `line-news-bot`
   - Set the following environment variables:

     ```env
     LINE_API_ACCESS_TOKEN=<<Access token from LINE Developer Console>>
     LINE_API_USER_ID=<<LINE user ID from LINE Developer Console>>
     LINE_API_URI="https://api.line.me/v2/bot/message/push"

     NEWS_API_URI="https://newsapi.org/v2/everything"
     NEWS_API_PARAMETER="?q=tech&sortBy=relevancy&pageSize=5&apiKey="
     NEWS_API_KEY=<<API key from NewsAPI>>
     ```

3. **Create an Amazon EventBridge schedule**
   - Set `line-news-bot` as the target
   - Create an IAM policy and role to allow Lambda execution
   - See: [EventBridge scheduled Lambda example (Qiita, Japanese)](https://qiita.com/shimabee/items/9cc7451eff44ef7f4769)

*You can set search conditions for articles as described in the [NewsAPI documentation](https://newsapi.org/docs/endpoints/everything).* 

## CI
- **Build and run the Docker container:**
  ```sh
  docker build -t line-news-bot-ci .
  docker run -p 9000:8080 -v .\aws-lambda-rie\aws-lambda-rie:/aws-lambda-rie --name line-news-bot-ci --env-file .env -e _HANDLER=function.handler -itd line-news-bot-ci
  ```
- **Check connectivity with curl:**
  ```sh
  curl -X POST http://localhost:9000/2015-03-31/functions/function/invocations ^
    -H "Content-Type: application/json" ^
    -d "{}"
  ```

## TODO
- CI/CD pipeline
- Feature expansion: fetch schedule from Google Calendar and send reminders
- Use Flex Message for improved design

## Key Points
- **Container runtime configuration:**
  Integration between Lambda execution environment and Lambda service
  - [AWS Lambda Runtime API Deep Dive (Japanese)](https://zenn.dev/hkdord/articles/lambda-handler-deep-dive)
  - [Java] Reduce boilerplate code with Lombok: https://www.casleyconsulting.co.jp/blog/engineer/107/

- **CI environment setup:**
  Since the execution environment is AWS Lambda, the test environment emulates Lambda.
  - https://github.com/aws/aws-lambda-runtime-interface-emulator
