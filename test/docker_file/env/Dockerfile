FROM golang:1.22-alpine

ENV APP_PORT=8080 APP_LOG=temp/log/app.log

RUN mkdir app/
COPY . app/

CMD cd app/ && go mod tidy && go run main.go
