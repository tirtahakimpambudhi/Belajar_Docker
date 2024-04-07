FROM golang:1.22-alpine

ENV APP_PORT="8081" APP_ENV="development" APP_LOG="app/temp/log"

RUN mkdir -p app
COPY . app/
RUN cd app/ && go mod tidy

VOLUME ${APP_LOG}
CMD cd app/ && go run main.go