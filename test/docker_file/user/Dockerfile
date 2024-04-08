FROM golang:1.22-alpine

RUN mkdir /app \
    && addgroup -S tirtahakimgroup \
    && adduser -S tirtahakim -G tirtahakimgroup \
    && chown -R tirtahakim:tirtahakimgroup /app

USER tirtahakim

WORKDIR /app
COPY . .
RUN go mod tidy

CMD go test ./test -v
