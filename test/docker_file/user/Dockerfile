FROM golang:1.18-alpine

RUN mkdir /app
RUN adduser -S tirtahakim
RUN adduser -S -D -h tirtahakim tirtahakimgroup
RUN chown -R tirtahakim:tirtahakimgroup /app
USER tirtahakim

WORKDIR /app
COPY . .
RUN go mod tidy

CMD go test ./test -v