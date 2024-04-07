FROM golang:1.22-alpine

ENV APP_PORT="8081" APP_ENV="development" APP_LOG="/app/temp/log"

WORKDIR /app
COPY . .
RUN go mod tidy

VOLUME ${APP_LOG}
CMD go run main.go