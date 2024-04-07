FROM golang:1.22-alpine

RUN mkdir -p app
COPY . app/
RUN cd app/ && go mod tidy

