FROM golang:1.22.2-alpine AS build
ARG APP_FILENAME=main
ENV APP_PORT=8080 APP_BIN_NAME=${APP_FILENAME}

RUN apk add --no-cache make
WORKDIR /app
COPY . .

RUN go mod tidy && go build -o bin/${APP_FILENAME} main.go

FROM alpine:latest 
FROM build 
ENV APP_PORT=${APP_PORT} APP_FILENAME=${APP_BIN_NAME}

WORKDIR /app
COPY --from=build /app/bin/${APP_FILENAME} ./bin/${APP_FILENAME}

CMD ./bin/${APP_FILENAME}