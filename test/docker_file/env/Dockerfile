FROM golang:1.22-alpine

ENV APP_PORT=8081 

RUN mkdir app/
COPY . app/

CMD cd app/ && go mod tidy && go run main.go