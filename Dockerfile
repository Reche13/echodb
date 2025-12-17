FROM golang:1.25.5-alpine AS builder
WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
    go build -ldflags="-s -w" -o echodb

FROM alpine

RUN mkdir -p /app/data
VOLUME /app/data

COPY --from=builder /app/echodb .

ENV ECHODB_PORT=6380
EXPOSE ${ECHODB_PORT}

ENTRYPOINT ["./echodb"]
