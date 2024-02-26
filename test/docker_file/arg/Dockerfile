FROM golang:1.22.0-alpine
ARG APP_FILENAME="main"
ENV APP_PORT="8081" APP_NAME=${APP_FILENAME} APP_ENV="development" APP_LOG="/app/temp/log"

WORKDIR /app
COPY . .
RUN mv main.go ${APP_FILENAME}.go
RUN go mod tidy

VOLUME ${APP_LOG}
CMD go run ${APP_NAME}.go