FROM golang:1.22-alpine

RUN mkdir /app \
    && addgroup -S tirtahakimgroup \
    && adduser -S tirtahakim -G tirtahakimgroup \
    && chown -R tirtahakim:tirtahakimgroup /app

USER tirtahakim

WORKDIR /app
COPY . .
RUN chmod -R 755 /app   # Giving full permissions for simplicity, adjust as necessary
RUN go mod tidy

CMD go test ./test -v
