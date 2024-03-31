FROM golang:1.18-alpine

RUN mkdir app
RUN echo "package main \nimport \"fmt\"\nfunc main(){fmt.Println(\"Hello World\")}" > main.go
CMD go run main.go
