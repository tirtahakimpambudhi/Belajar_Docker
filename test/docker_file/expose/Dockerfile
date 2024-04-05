FROM golang:1.18-alpine

RUN mkdir app/
COPY ./* app/
RUN go mod tidy

EXPOSE 8081 3500 
CMD [ "go" , "run" , "main.go" ]