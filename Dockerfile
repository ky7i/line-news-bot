FROM golang:1.24.1-alpine3.21 as build
WORKDIR /home
RUN apk add git \
&& git clone --depth 1 -b develop https://github.com/ky7i/line-news-bot
WORKDIR /home/line-news-bot/src
RUN CGO_ENABLED=0 go build -tags lambda.norpc -o main

FROM public.ecr.aws/lambda/provided:al2
COPY --from=build /home/line-news-bot/bootstrap ${LAMBDA_RUNTIME_DIR}
# Copy function code
COPY --from=build /home/line-news-bot/function.sh ${LAMBDA_TASK_ROOT}
# Copy exec file 
COPY --from=build /home/line-news-bot/src/main /home/ 
RUN chmod +x "${LAMBDA_RUNTIME_DIR}"/bootstrap \
&& chmod +x "${LAMBDA_TASK_ROOT}"/function.sh
# Set the CMD to your handler (could also be done as a parameter override outside of the Dockerfile)
CMD [ "function.handler" ]  