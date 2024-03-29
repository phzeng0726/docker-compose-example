# 使用 Python 官方映像作為基礎映像
FROM python:3.9

# 設置工作目錄
WORKDIR /app

COPY requirements.txt ./

RUN pip install --no-cache-dir -r requirements.txt


# 複製整個應用程式代碼到容器中
COPY . .

# 開放容器的端口
EXPOSE 5000

# 定義啟動容器時運行的命令
CMD ["python", "./main.py"]
