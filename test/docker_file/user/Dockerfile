FROM golang:1.18-alpine

RUN mkdir /app \
    && addgroup -S tirtahakimgroup \
    && adduser -S tirtahakim -G tirtahakimgroup \
    && chown -R tirtahakim:tirtahakimgroup /app

USER tirtahakim

WORKDIR /app
COPY . .

# Change ownership to tirtahakim before modifying permissions
USER root
RUN chown -R tirtahakim:tirtahakimgroup /app
RUN go mod init go_user

# Switch back to tirtahakim to run subsequent commands
USER tirtahakim

# Adjust permissions
RUN chmod -R 755 /app

RUN go mod tidy

CMD go test ./test -v
