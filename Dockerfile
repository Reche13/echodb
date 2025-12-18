FROM golang:1.25.5-alpine AS builder

WORKDIR /app

RUN apk add --no-cache make

COPY . .

RUN make build-prod \
    GOOS=linux \
    GOARCH=amd64 \
    BUILD_DIR=/app/dist

FROM alpine:latest

WORKDIR /app

RUN mkdir -p /app/data
VOLUME /app/data

COPY --from=builder /app/dist/echodb /app/echodb

ENV ECHODB_PORT=6380
EXPOSE 6380

ENTRYPOINT ["./echodb"]
