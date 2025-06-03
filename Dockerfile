FROM golang:1.24.1-alpine3.21 as build
WORKDIR /home
RUN apk add git
RUN git clone https://github.com/ky7i/line-news-bot
WORKDIR /home/line-news-bot/src
RUN CGO_ENABLED=0 go build -tags lambda.norpc -o bootstrap

FROM public.ecr.aws/lambda/provided:al2
COPY --from=build /home/line-news-bot/src/bootstrap /var/runtime/bootstrap
RUN chmod +x /var/runtime/bootstrap