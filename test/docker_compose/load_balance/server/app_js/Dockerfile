FROM oven/bun:1.1.7-alpine AS build
ARG APP_FILENAME=index
ENV APP_BIN_NAME=${APP_FILENAME}
WORKDIR /app
COPY package.json .
COPY bun.lockb .
RUN bun install --frozen-lockfile
COPY . .
RUN bun build index.js --compile --outfile ${APP_BIN_NAME}

FROM alpine:latest
FROM build
ENV APP_FILENAME=${APP_BIN_NAME}
WORKDIR /app_bin
RUN mkdir bin/
COPY  --from=build /app/${APP_FILENAME} ./bin/${APP_FILENAME}
CMD ./bin/${APP_FILENAME}