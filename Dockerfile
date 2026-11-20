FROM golang:1.22 AS builder

WORKDIR /app

COPY . .

RUN go mod tidy && go build -o main ./cmd/main.go

# 使用更小的基础镜像，减少体积
FROM debian:bookworm-slim

RUN apt-get update && apt-get install -y --no-install-recommends ca-certificates \
    && rm -rf /var/lib/apt/lists/*

WORKDIR /app

COPY --from=builder /app/main .

CMD ["./main"]
