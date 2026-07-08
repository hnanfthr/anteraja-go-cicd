# Tahap 1: Memasak Kode Go
FROM golang:1.22-alpine AS builder
WORKDIR /app
COPY go.mod ./
COPY . .
# Compile kode Go jadi file binary mandiri bernama 'main'
RUN CGO_ENABLED=0 GOOS=linux go build -o main .

# Tahap 2: Menjalankan Aplikasi (Image super ringan)
FROM alpine:latest
WORKDIR /app
# Ambil hasil compile dari Tahap 1
COPY --from=builder /app/main .
# Ambil folder UI HTML lu
COPY --from=builder /app/templates ./templates

EXPOSE 8080
CMD ["./main"]