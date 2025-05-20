# ----------------------------
# 1) Build stage
# ----------------------------
FROM golang:1.24.3-alpine AS builder

# ติดตั้งเครื่องมือที่ต้องใช้ (gcc, musl, git)
RUN apk add --no-cache gcc musl-dev git

# เปิดใช้งาน cgo
ENV CGO_ENABLED=1

# ตั้ง GOBIN ให้ตรงกับ GOPATH/bin และเพิ่มใน PATH
ENV GOBIN=/go/bin
ENV PATH=$PATH:$GOBIN

WORKDIR /app

# โหลด dependencies ก่อน เพื่อ cache layer ให้ build เร็วขึ้น
COPY go.mod go.sum ./
RUN go mod download

# ก็อปทั้งโปรเจกต์เข้าไป
COPY . .

# ติดตั้ง swag CLI (จะไปอยู่ที่ /go/bin/swag)
RUN go install github.com/swaggo/swag/cmd/swag@latest

# สร้าง swagger docs ออกมาในโฟลเดอร์ docs
RUN swag init -g main.go -o docs

# คอมไพล์ binary
RUN go build -o app main.go



# ----------------------------
# 2) Run stage
# ----------------------------
FROM alpine:latest

WORKDIR /app

# ก็อปเฉพาะ binary จาก builder
COPY --from=builder /app/app .

# ก็อป swagger docs มาด้วย (ถ้าคุณใช้ docs ใน runtime เช่น serve static)
COPY --from=builder /app/docs ./docs

EXPOSE 3030

CMD ["./app"]
