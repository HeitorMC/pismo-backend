FROM golang:1.23-alpine AS builder

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN go build -o migrations ./migrate/migrate.go

RUN go build -o server ./cmd/main.go

FROM alpine:latest

WORKDIR /app

COPY --from=builder /app/server .

COPY --from=builder /app/migrations .

EXPOSE ${PORT}

CMD ["sh", "-c", "./migrations && ./server"]
