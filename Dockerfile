FROM golang:1.23

WORKDIR /app

# Копируем файлы зависимостей из папки backend
COPY backend/go.mod backend/go.sum ./
RUN go mod download

# Копируем исходный код проекта (только backend, если frontend не нужен)
COPY backend/ ./backend/

# Копируем файлы конфигурации (убедись, что пути соответствуют твоей структуре)
COPY backend/config/.env ./backend/internal/config/.env
COPY backend/.env .env

# Собираем бинарник, указывая путь к main.go внутри backend
RUN go build -o beer_bot ./backend/cmd/main.go

CMD ["./beer_bot"]
