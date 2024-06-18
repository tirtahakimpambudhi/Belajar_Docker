FROM golang:1.22-alpine as build

RUN apk add --no-cache make
WORKDIR /app
COPY . .
RUN go mod tidy && make migration && make build 


FROM alpine:latest
RUN apk add --no-cache make
WORKDIR /app
COPY --from=build /app .

CMD make migrationUp && make migrationUp && make run
