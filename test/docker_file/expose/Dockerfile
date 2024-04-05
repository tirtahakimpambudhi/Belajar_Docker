FROM golang:1.18-alpine

RUN mkdir app/
COPY main.go app/

EXPOSE 8081
CMD [ "go" , "run" , "main.go" ]