# ビルドステージ
FROM golang:1.21-alpine AS builder

WORKDIR /app

# 依存関係をコピーしてダウンロード
COPY go.mod go.sum ./
RUN go mod download

# ソースコードをコピーしてビルド
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o link_slider_backend .

# 実行ステージ
FROM alpine:latest

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/link_slider_backend .

EXPOSE 8080

CMD ["./link_slider_backend"]
