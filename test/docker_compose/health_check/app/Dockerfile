FROM golang:1.22.0-alpine as build
ARG APP_FILENAME="main"
ENV APP_PORT="8081" APP_NAME=${APP_FILENAME} APP_ENV="development" APP_LOG="/app/temp/log"


WORKDIR /app
COPY . .
RUN mv main.go ${APP_FILENAME}.go
RUN go mod tidy

VOLUME ${APP_LOG}
RUN go build -o ./bin/${APP_FILENAME} ./${APP_FILENAME}.go

FROM alpine:latest
FROM build
ENV APP_FILENAME=${APP_NAME}
RUN apk --no-cache add curl
WORKDIR /bin
COPY --from=build /app/bin/${APP_FILENAME} .


CMD ./${APP_FILENAME}

