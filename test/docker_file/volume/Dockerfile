FROM golang:1.22-alpine

ENV APP_PORT="8081" APP_ENV="development"

RUN mkdir -p app
COPY . app/
RUN cd app/ && go mod tidy

CMD cd app/ && go run main.go