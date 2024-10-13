FROM golang:1.22-alpine AS builder

WORKDIR /app

RUN apk --no-cache add make curl

COPY go.mod go.sum ./

RUN go mod download

RUN curl -L https://github.com/golang-migrate/migrate/releases/latest/download/migrate.linux-amd64.tar.gz | tar xz && \
    mv migrate /usr/local/bin/ && \
    chmod +x /usr/local/bin/migrate

COPY . .

RUN go build -o myapp .

RUN ls -la /app

FROM alpine:latest

RUN apk --no-cache add postgresql-client

COPY --from=builder /usr/local/bin/migrate /usr/local/bin/migrate

WORKDIR /app

COPY --from=builder /app/myapp .
COPY --from=builder /app/entrypoint.sh .
COPY --from=builder /app/docker-compose.yml .
COPY --from=builder /app/migrations ./migrations

RUN chmod +x ./entrypoint.sh

ENTRYPOINT ["./entrypoint.sh"]
