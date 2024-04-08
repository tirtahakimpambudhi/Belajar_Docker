FROM golang:1.18-alpine

RUN mkdir /app
RUN adduser -S tirtahakim
RUN adduser -S -D -h tirtahakim tirtahakimgroup
RUN chown -R tirtahakim:tirtahakimgroup /app
USER tirtahakim

COPY . /app

CMD go 