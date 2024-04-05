FROM golang:1.22-alpine

RUN mkdir app/
COPY . app/

EXPOSE 8081 3500 
CMD cd app/ && go mod tidy && go run main.go