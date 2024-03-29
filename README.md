# Docker Compose Example

讓 Python 與 Golang 的 Docker Image 可以一起 Run 起來。

## Running

```SHELL=
docker compose up --build
```

- Python API: http://localhost:5000/tasks (獲取全部 tasks)
- Golang API: http://localhost:8080/tasks/1 (獲取單一 task)

> 目前兩個 APP 的 port 不同，要如何變成同一個 Port?
> 疑似可以用 nginx
