# 使用 Golang 官方映像作為基礎映像
FROM golang:1.22.1

# 設置工作目錄
WORKDIR /app

# 複製整個應用程式代碼到容器中
COPY . .

# 編譯應用程式
RUN go build -o main .

# 開放容器的端口
EXPOSE 8080

# 定義啟動容器時運行的命令
CMD ["./main"]
