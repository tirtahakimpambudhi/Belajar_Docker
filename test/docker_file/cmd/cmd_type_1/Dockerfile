FROM golang:1.18-alpine

RUN mkdir /app && \
    echo "package main" > /app/main.go && \
    echo "import \"fmt\"" >> /app/main.go && \
    echo "func main() { fmt.Println(\"Hello, World!\") }" >> /app/main.go

CMD go run /app/main.go
